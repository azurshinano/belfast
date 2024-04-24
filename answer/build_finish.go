package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func BuildFinish(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_12025
	err := proto.Unmarshal(*buffer, &data)
	if err != nil {
		return 0, 12026, err
	}

	response := protobuf.SC_12026{
		Result: proto.Uint32(0),
	}

	response.ShipList = make([]*protobuf.SHIPINFO, len(data.GetPosList()))
	var minPos uint32 = 999999
	var maxPos uint32
	for _, pos := range data.GetPosList() {
		if pos < minPos {
			minPos = pos
		}
		if pos > maxPos {
			maxPos = pos
		}
	}
	builds, err := client.Commander.GetBuildRange(minPos, maxPos)
	if err != nil {
		return 0, 12026, err
	}

	for i := range data.GetPosList() {
		ship, err := builds[i].Consume(builds[i].ShipID, client.Commander)
		if err != nil {
			return 0, 12026, err
		}
		response.ShipList[i] = ship
	}

	return client.SendMessage(12026, &response)
}
