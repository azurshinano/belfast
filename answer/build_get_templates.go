package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

// Get the corresponding template_id in the ship build list
func BuildGetTemplates(buffer *[]byte, client *connection.Client) (int, int, error) {
	buildInfos := make([]*protobuf.BUILD_INFO, len(client.Commander.Builds))
	builds, err := client.Commander.GetBuildRange(1, uint32(len(client.Commander.Builds)))
	if err != nil {
		return 0, 12044, err
	}

	for i, work := range builds {
		buildInfos[i] = &protobuf.BUILD_INFO{
			Pos: proto.Uint32(uint32(i + 1)),
			Tid: proto.Uint32(work.ShipID),
		}
	}

	response := protobuf.SC_12044{
		InfoList: buildInfos,
	}

	return client.SendMessage(12044, &response)
}
