package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
	"time"
)

// TODO: MOCK
func GuildShopAnswer(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_60034{
		Info: &protobuf.SHOP_INFO{
			RefreshCount:    proto.Uint32(0),
			NextRefreshTime: proto.Uint32(uint32(time.Now().Unix() + 1231111)),
			GoodList: []*protobuf.GOODS_INFO{
				{
					Id:           proto.Uint32(113),
					Count:        proto.Uint32(5),
					BoughtRecord: []uint32{1},
				},
			},
		},
		Result: proto.Uint32(0),
	}
	return client.SendMessage(60034, &response)
}
