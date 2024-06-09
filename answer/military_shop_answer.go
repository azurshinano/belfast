package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
	"time"
)

// TODO: MOCK
func MilitaryShopAnswer(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_18101{
		ArenaShopList: []*protobuf.ARENASHOP{
			{
				ShopId: proto.Uint32(33021),
				Count:  proto.Uint32(0),
			},
		},
		NextFlashTime: proto.Uint32(uint32(time.Now().Unix() + 1231111)),
		FlashCount:    proto.Uint32(0),
	}
	return client.SendMessage(18101, &response)
}
