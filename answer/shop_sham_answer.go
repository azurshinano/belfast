package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
	"time"
)

// 核心月度商店和原型商店数据列表, 目前似乎只传了月份?
func ShopShamAndFragmentAnswer(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_16200{
		Month: proto.Uint32(uint32(time.Now().Month())),
	}

	return client.SendMessage(16200, &response)
}
