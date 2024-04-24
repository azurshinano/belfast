package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/logger"

	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func RetireShip(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_12004
	err := proto.Unmarshal(*buffer, &data)
	if err != nil {
		return 0, 12005, err
	}
	response := protobuf.SC_12005{
		Result: proto.Uint32(0),
	}
	if err := client.Commander.RetireShips(&data.ShipIdList); err != nil {
		response.Result = proto.Uint32(1)
		logger.LogEvent("RetireShip", "Fail", err.Error(), logger.LOG_LEVEL_ERROR)
	} else {
		response.ShipIdList = data.ShipIdList
	}
	return client.SendMessage(12005, &response)
}
