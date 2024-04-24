package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func GetShipDiscuss(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_17101
	if err := proto.Unmarshal(*buffer, &data); err != nil {
		return 0, 17102, err
	}
	var response protobuf.SC_17102

	response.ShipDiscuss = &protobuf.SHIP_DISCUSS_INFO{
		ShipGroupId:       data.ShipGroupId,
		DiscussCount:      proto.Uint32(0),
		HeartCount:        proto.Uint32(0),
		DailyDiscussCount: proto.Uint32(0),
	}

	return client.SendMessage(17102, &response)
}
