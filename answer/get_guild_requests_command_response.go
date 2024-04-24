package answer

import (
	"github.com/ggmolly/belfast/connection"

	"github.com/ggmolly/belfast/protobuf"
)

func GetGuildRequestsCommandResponse(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_60004
	return client.SendMessage(60004, &response)
}
