package answer

import (
	"github.com/ggmolly/belfast/connection"

	"github.com/ggmolly/belfast/protobuf"
)

func UNK_26102(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_26102
	return client.SendMessage(26102, &response)
}
