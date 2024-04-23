package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

// TODO: Add some specific items that should be obtained after completing the story?
func UpdateOwnedStory(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_11017
	if err := proto.Unmarshal(*buffer, &data); err != nil {
		return 0, 11018, err
	}
	response := protobuf.SC_11018{
		Result: proto.Uint32(0),
	}

	if err := client.Commander.AddStory(data.GetStoryId()); err != nil {
		response.Result = proto.Uint32(1)
	}

	return client.SendMessage(11018, &response)
}
