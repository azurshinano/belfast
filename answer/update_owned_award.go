package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

// TODO: Add some specific items that should be obtained after completing the award?
func UpdateOwnedAward(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_17005
	if err := proto.Unmarshal(*buffer, &data); err != nil {
		return 0, 17006, err
	}
	response := protobuf.SC_17006{
		Result: proto.Uint32(0),
	}

	if err := client.Commander.AddAward(data.GetId(), data.GetAwardIndex()); err != nil {
		response.Result = proto.Uint32(1)
	}

	return client.SendMessage(17006, &response)
}
