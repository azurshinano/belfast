package answer

import (
	"fmt"
	"strconv"

	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/logger"
	"github.com/ggmolly/belfast/orm"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func Forge_SC10021(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_10020
	err := proto.Unmarshal(*buffer, &data)
	if err != nil {
		return 0, 10021, err
	}

	response := protobuf.SC_10021{
		ServerTicket: proto.String("=*=*=*=BELFAST=*=*=*="),
		Result:       proto.Uint32(0),
		Device:       proto.Uint32(0),
	}

	var yostarusAuth orm.YostarusMap
	intArg2, err := strconv.Atoi(data.GetArg2())
	if err != nil {
		return 0, 10021, fmt.Errorf("failed to convert arg2 to int: %s", err.Error())
	}

	err = orm.GormDB.
		Where("arg2 = ?", intArg2).
		First(&yostarusAuth).
		Error

	// Check if the account exists
	if err != nil {
		// If not, create it
		accountId, err := client.CreateCommander(uint32(intArg2))
		if err != nil {
			logger.LogEvent("Server", "SC_10021", fmt.Sprintf("failed to create commander: %s", err.Error()), logger.LOG_LEVEL_ERROR)
			return 0, 10021, err
		}
		response.AccountId = proto.Uint32(accountId)
	} else {
		response.AccountId = proto.Uint32(yostarusAuth.AccountID)
	}

	// Update server list
	var belfastServers []orm.Server
	// Decrease by 1 the state of all servers
	err = orm.GormDB.Order("id asc").Find(&belfastServers).Error
	if err != nil {
		logger.LogEvent("Server", "SC_10021", fmt.Sprintf("failed to fetch servers: %s", err.Error()), logger.LOG_LEVEL_ERROR)
		return 0, 10021, err
	}
	Servers = make([]*protobuf.SERVERINFO, len(belfastServers))
	for i, server := range belfastServers {
		Servers[i] = &protobuf.SERVERINFO{
			Ids:   []uint32{server.ID},
			Ip:    proto.String(server.IP),
			Port:  proto.Uint32(server.Port),
			State: proto.Uint32(*server.StateID - 1), // StateID is 0-based in Azur Lane, but 1-based in the database
			Name:  proto.String(server.Name),
			Sort:  proto.Uint32(uint32(i + 1)),
		}
	}
	response.Serverlist = Servers
	logger.LogEvent("Server", "SC_10021", fmt.Sprintf("sending %d servers", len(response.Serverlist)), logger.LOG_LEVEL_WARN)
	return client.SendMessage(10021, &response)
}
