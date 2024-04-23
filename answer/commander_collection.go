package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/orm"

	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

type MaxShipStat struct {
	GroupID     uint32 `gorm:"column:group_id"`
	MaxStar     uint32 `gorm:"column:max_star"`
	MaxIntimacy uint32 `gorm:"column:max_intimacy"`
	MaxLevel    uint32 `gorm:"column:max_level"`
}

func CommanderCollection(buffer *[]byte, client *connection.Client) (int, int, error) {
	// Out of all commander's OwnedShips, return the max star, max intimacy, and max level
	// of each ship group (= TemplateID divided by 10)
	stats := []*protobuf.SHIP_STATISTICS_INFO{}
	orm.GormDB.Raw(`
	SELECT
		ship_id / 10 AS group_id,
		MAX(ships.star) AS max_star,
		MAX(intimacy) AS max_intimacy,
		MAX(level) AS max_level,
		MAX(propose::int) AS marry_flag,
		(SELECT COUNT(*) FROM likes WHERE group_id = owned_ships.ship_id / 10 AND liker_id = ?) AS heart_flag,
		(SELECT COUNT(*) FROM likes WHERE group_id = owned_ships.ship_id / 10) AS heart_count
	FROM owned_ships
	INNER JOIN ships ON owned_ships.ship_id = ships.template_id
	WHERE owner_id = ?
	GROUP BY group_id, owned_ships.ship_id
	`, client.Commander.CommanderID, client.Commander.CommanderID).Scan(&stats)

	awardsMap := map[uint32][]uint32{}
	for _, v := range client.Commander.OwnedAwards {
		if m, ok := awardsMap[v.TypeID]; ok {
			awardsMap[v.TypeID] = append(m, v.AwardID)
		} else {
			awardsMap[v.TypeID] = []uint32{v.AwardID}
		}
	}
	awards := []*protobuf.SHIP_STATISTICS_AWARD{}
	for k, v := range awardsMap {
		awards = append(awards, &protobuf.SHIP_STATISTICS_AWARD{
			Id:         proto.Uint32(k),
			AwardIndex: v,
		})
	}

	response := protobuf.SC_17001{
		DailyDiscuss:  proto.Uint32(0),
		ShipInfoList:  stats,
		ShipAwardList: awards,
	}
	return client.SendMessage(17001, &response)
}
