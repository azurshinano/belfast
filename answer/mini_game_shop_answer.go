package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
	"time"
)

// TODO: MOCK
func MiniGameShopAnswer(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_26151{
		NextFlashTime: proto.Uint32(uint32(time.Now().Unix() + 1231111)),
		Goods: []*protobuf.GOODS_INFO{
			{
				Id:    proto.Uint32(12),
				Count: proto.Uint32(1),
			},
		},
	}
	return client.SendMessage(26151, &response)
}
