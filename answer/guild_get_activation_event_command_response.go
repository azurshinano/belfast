package answer

import (
	"github.com/ggmolly/belfast/connection"

	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func GuildGetActivationEventCommandResponse(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_61006
	proto.Unmarshal([]byte{0x08, 0x14}, &response)

	return client.SendMessage(61006, &response)
}
