package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf/out/p10_pb"
	"google.golang.org/protobuf/proto"
)

func Forge_SC10996(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := p10_pb.Sc_10997{
		Version1:    proto.Uint32(1),
		Version2:    proto.Uint32(3),
		Version3:    proto.Uint32(0),
		Version4:    proto.Uint32(0),
		GatewayIp:   proto.String("line1-login-bili-blhx.bilibiligame.net"),
		GatewayPort: proto.Uint32(80),
		Url:         proto.String("https://blhx.biligame.com"),
	}

	return client.SendMessage(10997, &response)
}
