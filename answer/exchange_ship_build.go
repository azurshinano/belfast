package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/orm"

	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func ExchangeShipBuild(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_16100
	err := proto.Unmarshal(*buffer, &data)
	if err != nil {
		return 0, 16101, err
	}
	var response protobuf.SC_16101

	if !client.Commander.HasEnoughItem(15001, 6*data.GetCnt()) {
		response.Result = proto.Uint32(1)
		return client.SendMessage(16101, &response)
	}

	response.ShipList = make([]*protobuf.SHIPINFO, data.GetCnt())
	for i := 0; i < int(data.GetCnt()); i++ {
		ship, err := orm.GetRandomPoolShip(0)
		if err != nil {
			return 0, 16101, err
		}
		info, err := client.Commander.AddShip(ship.TemplateID)
		if err != nil {
			return 0, 16101, err
		}
		response.ShipList[i] = &info
	}
	response.Result = proto.Uint32(0)
	client.Commander.ConsumeItem(15001, 6*data.GetCnt())
	return client.SendMessage(16101, &response)
}
