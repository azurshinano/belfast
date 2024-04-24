package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/logger"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
	"strconv"
)

// 使用快速完成工具
func BuildQuickFinish(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_12008
	err := proto.Unmarshal(*buffer, &data)
	if err != nil {
		return 0, 12008, err
	}

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

	logger.LogEvent("test", "test0", strconv.Itoa(int(minPos))+","+strconv.Itoa(int(maxPos)), logger.LOG_LEVEL_INFO)
	builds, err := client.Commander.GetBuildRange(minPos, maxPos)
	if err != nil {
		return 0, 12008, err
	}
	logger.LogEvent("test", "test1", strconv.Itoa(len(builds)), logger.LOG_LEVEL_INFO)

	for _, build := range builds {
		if err := build.QuickFinish(client.Commander); err != nil {
			return 0, 12008, err
		}
	}

	if err != nil {
		return 0, 12008, err
	}

	response := protobuf.SC_12009{
		Result:  proto.Uint32(0),
		PosList: data.GetPosList(),
	}

	return client.SendMessage(12009, &response)
}
