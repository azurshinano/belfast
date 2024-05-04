package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func ShopStreetAnswer(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_22102{
		Street: &protobuf.SHOPPINGSTREET{
			// 目前只有1, 写死
			Lv:            proto.Uint32(1),
			NextFlashTime: proto.Uint32(0),
			// 目前没有升级, 写死
			LvUpTime:   proto.Uint32(0),
			GoodsList:  make([]*protobuf.STREETGOODS, 0),
			FlashCount: proto.Uint32(0),
		},
	}

	return client.SendMessage(22102, &response)
}
