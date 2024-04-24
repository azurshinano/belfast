package answer

import (
	"fmt"
	"time"

	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/logger"

	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

// Reimplementation of SC_11000
func LastLogin(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_11000{
		Timestamp:               proto.Uint32(uint32(time.Now().Unix())),
		Monday_0OclockTimestamp: proto.Uint32(1606114800), // 23/11/2020 08:00:00
	}
	client.Commander.BumpLastLogin()
	logger.LogEvent("Server", "SC_11000", "Updated last login of uid="+fmt.Sprint(client.Commander.CommanderID), logger.LOG_LEVEL_INFO)
	return client.SendMessage(11000, &response)
}
