package answer

import (
	"fmt"
	"os"
	"time"

	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/consts"
	"github.com/ggmolly/belfast/misc"

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
	var data protobuf.CS_10800
	err := proto.Unmarshal(*buffer, &data)
	if err != nil {
		return 0, 10801, err
	}

	if len(versions) == 0 {
		hashes := misc.GetGameHashes()
		for _, hash := range hashes {
			versions = append(versions, hash.Hash)
		}
		versions = append(versions, "dTag-1")
	}
	belfastRegion := os.Getenv("AL_REGION")

	// It seems like the game kind of ignore anything but the versions, timestamp & Monday_0OclockTimestamp
	response := protobuf.SC_10801{
		GatewayIp:               proto.String(consts.RegionGateways[belfastRegion]),
		GatewayPort:             proto.Uint32(80),
		Url:                     proto.String(""),
		Version:                 versions,
		ProxyIp:                 proto.String(consts.RegionProxies[belfastRegion]),
		ProxyPort:               proto.Uint32(20000),
		IsTs:                    proto.Uint32(0),
		Timestamp:               proto.Uint32(uint32(time.Now().Unix())),
		Monday_0OclockTimestamp: proto.Uint32(consts.Monday_0OclockTimestamps[belfastRegion]),

		// wtf is this i don't even understand what monday_0oclock_timestamp is
		// who would even do such a thing
	}

	resolvedPlatform, ok := platformMap[data.GetPlatform()]
	if !ok {
		resolvedPlatform = "Unknown"
	}

	url, ok := consts.GamePlatformUrl[belfastRegion][data.GetPlatform()]
	if !ok {
		return 0, 10801, fmt.Errorf("unknown platform '%s' (id='%s')", resolvedPlatform, data.GetPlatform())
	} else {
		response.Url = proto.String(url)
	}

	return client.SendMessage(10801, &response)
}
