package answer

import (
	"fmt"
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/logger"
	"github.com/ggmolly/belfast/orm"
	"github.com/ggmolly/belfast/protobuf/out/p10_pb"
	"google.golang.org/protobuf/proto"
)

func Forge10020(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data p10_pb.Cs_10020
	err := proto.Unmarshal(*buffer, &data)
	if err != nil {
		return 0, 10021, err
	}

	response := p10_pb.Sc_10021{
		Result: proto.Uint32(0),
		// 跨服鉴权要用, 目前我们不需要
		ServerTicket: proto.String("azurlane"),
		// 网关这边的公告
		NoticeList: []*p10_pb.Noticeinfo{},
	}

	var guser orm.GatewayUser
	err = orm.GormDB.Where("username = ?", data.GetArg2()).First(&guser).Error
	// 目前逻辑是不存在就新建用户, 后续待优化
	if err != nil {
		accountId, err := client.CreateCommander(data.GetArg2())
		if err != nil {
			logger.LogEvent("Server", "SC_10021", fmt.Sprintf("failed to create commander: %s", err.Error()), logger.LOG_LEVEL_ERROR)
			return 0, 10021, err
		}
		response.AccountId = proto.Uint32(accountId)
	} else {
		response.AccountId = proto.Uint32(guser.AccountID)
	}

	// 服务器列表
	var belfastServers []orm.Server
	err = orm.GormDB.Order("id asc").Find(&belfastServers).Error
	if err != nil {
		logger.LogEvent("Server", "SC_10021", fmt.Sprintf("failed to fetch servers: %s", err.Error()), logger.LOG_LEVEL_ERROR)
		return 0, 10021, err
	}
	var servers []*p10_pb.Serverinfo
	for _, server := range belfastServers {
		servers = append(servers, &p10_pb.Serverinfo{
			Ids:  []uint32{server.ID},
			Ip:   proto.String(server.IP),
			Port: proto.Uint32(server.Port),
			// StateID is 0-based in Azur Lane, but 1-based in the database
			State: proto.Uint32(*server.StateID - 1),
			Name:  proto.String(server.Name),
		})
	}
	response.Serverlist = servers
	logger.LogEvent("Server", "SC_10021", fmt.Sprintf("sending %d servers", len(response.Serverlist)), logger.LOG_LEVEL_WARN)

	return client.SendMessage(10021, &response)
}
