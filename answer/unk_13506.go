package answer

import (
	"github.com/ggmolly/belfast/connection"

	"github.com/ggmolly/belfast/protobuf"
)

func UNK_13506(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_13506
	return client.SendMessage(13506, &response)
}
