package answer

import (
	"time"

	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

var (
	platformMap = map[string]string{
		"0": "Android", // Possibly means 'Play Store'
		"1": "iOS",     // Possibly means 'App Store'
		// maybe more?
	}
	versions []string
)

// Answer to a CS_10800 packet with a SC_10801 packet + hashes
func Forge_SC10801(buffer *[]byte, client *connection.Client) (int, int, error) {
	const packetId = 10801
	var updateCheck protobuf.CS_10800
	err := proto.Unmarshal(*buffer, &updateCheck)
	if err != nil {
		return 0, packetId, err
	}

	if len(versions) == 0 {
		hashes := []string{"1"}
		for _, hash := range hashes {
			versions = append(versions, hash)
		}
		versions = append(versions, "dTag-1")
	}

	// It seems like the game kind of ignore anything but the versions, timestamp & Monday_0OclockTimestamp
	response := protobuf.SC_10801{
		GatewayIp:               proto.String("line1-login-bili-blhx.bilibiligame.net"),
		GatewayPort:             proto.Uint32(80),
		Url:                     proto.String(""),
		Version:                 versions,
		ProxyIp:                 proto.String("line1-bak-login-bili-blhx.bilibiligame.net"),
		ProxyPort:               proto.Uint32(20000),
		IsTs:                    proto.Uint32(0),
		Timestamp:               proto.Uint32(uint32(time.Now().Unix())),
		Monday_0OclockTimestamp: proto.Uint32(1606060800),

		// wtf is this i don't even understand what monday_0oclock_timestamp is
		// who would even do such a thing
	}

	response.Url = proto.String("https://blhx.biligame.com")

	return client.SendMessage(packetId, &response)
}
