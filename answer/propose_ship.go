package answer

import (
	"fmt"
	"time"

	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/logger"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func ProposeShip(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_12032
	if err := proto.Unmarshal(*buffer, &data); err != nil {
		return 0, 12033, err
	}
	logger.LogEvent("Dock", "Propose", fmt.Sprintf("uid=%d has proposed ship id=%d", client.Commander.CommanderID, data.GetShipId()), logger.LOG_LEVEL_DEBUG)

	shipId := data.GetShipId()
	response := protobuf.SC_12033{}
	response.Result = proto.Uint32(0)
	response.Time = proto.Uint32(uint32(time.Now().Unix()))
	// Check if the ship exists
	ship, ok := client.Commander.OwnedShipsMap[shipId]
	if !ok {
		logger.LogEvent("Dock", "Propose", fmt.Sprintf("uid=%d has proposed ship id=%d, but it doesn't exist", client.Commander.CommanderID, shipId), logger.LOG_LEVEL_ERROR)
		response.Result = proto.Uint32(1)
	}
	// Check if the ship is already proposed
	if ship.Propose {
		logger.LogEvent("Dock", "Propose", fmt.Sprintf("uid=%d has proposed ship id=%d, but it's already proposed", client.Commander.CommanderID, shipId), logger.LOG_LEVEL_ERROR)
		response.Result = proto.Uint32(1)
	}
	// Check if the commander has a promise ring (id=15006)
	if !client.Commander.HasEnoughItem(15006, 1) {
		logger.LogEvent("Dock", "Propose", fmt.Sprintf("uid=%d has proposed ship id=%d, but doesn't have a promise ring", client.Commander.CommanderID, shipId), logger.LOG_LEVEL_ERROR)
		response.Result = proto.Uint32(1)
	}
	// Consume the promise ring
	err := client.Commander.ConsumeItem(15006, 1)
	if err != nil {
		logger.LogEvent("Dock", "Propose", fmt.Sprintf("uid=%d has proposed ship id=%d, but failed to consume the promise ring: %s", client.Commander.CommanderID, shipId, err.Error()), logger.LOG_LEVEL_ERROR)
		response.Result = proto.Uint32(1)
	}
	// Propose the ship
	err = ship.ProposeShip()
	if err != nil {
		logger.LogEvent("Dock", "Propose", fmt.Sprintf("uid=%d has proposed ship id=%d, but it failed: %s", client.Commander.CommanderID, shipId, err.Error()), logger.LOG_LEVEL_ERROR)
		response.Result = proto.Uint32(1)
	}
	logger.LogEvent("Dock", "Propose", fmt.Sprintf("uid=%d has proposed ship id=%d successfully", client.Commander.CommanderID, shipId), logger.LOG_LEVEL_INFO)

	return client.SendMessage(12033, &response)
}
