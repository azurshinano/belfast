package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

// In the Azurlane code, it is called ServerInterconnection, used to switch Android/iOS servers
func Forge_SC10803_CN_JP_KR_TW(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_10803_CN_JP_KR_TW{
		GatewayIp:   proto.String("line1-login-bili-blhx.bilibiligame.net"),
		GatewayPort: proto.Uint32(80),
		ProxyIp:     proto.String("line1-bak-login-bili-blhx.bilibiligame.net"),
		ProxyPort:   proto.Uint32(20000),
	}

	return client.SendMessage(10803, &response)
}
