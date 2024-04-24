package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func UpdateShipLike(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_17107
	if err := proto.Unmarshal(*buffer, &data); err != nil {
		return 0, 17108, err
	}
	response := protobuf.SC_17108{
		Result: proto.Uint32(boolToUint32(client.Commander.Like(*data.ShipGroupId) != nil)),
	}
	return client.SendMessage(17108, &response)
}
