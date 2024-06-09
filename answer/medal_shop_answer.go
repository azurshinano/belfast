package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
	"time"
)

// TODO: MOCK
func MedalShopAnswer(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_16107{
		Result:        proto.Uint32(0),
		ItemFlashTime: proto.Uint32(uint32(time.Now().Unix() + 1231111)),
		GoodList: []*protobuf.GOODS_INFO{
			{
				Id:    proto.Uint32(38),
				Count: proto.Uint32(1),
			},
		},
	}
	return client.SendMessage(16107, &response)
}
