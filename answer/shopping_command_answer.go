package answer

import (
	"fmt"

	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/logger"
	"github.com/ggmolly/belfast/orm"

	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

// Drop List type :
// 0 = upgrades
// 1 = resource
// 2 = items
// 3 = does not exist
// 4 = ship
// 5 = does not exist
// 6 = skins
// 7 = not an ownable
// 8 = not an ownable
// 9 = equip skin
// 10 = not an ownable
// 12 = operation siren stuff
// 20 = ?

func ShoppingCommandAnswer(buffer *[]byte, client *connection.Client) (int, int, error) {
	var boughtOffer protobuf.CS_16001
	err := proto.Unmarshal(*buffer, &boughtOffer)
	if err != nil {
		return 0, 16002, err
	}

	var shopOffer orm.ShopOffer
	orm.GormDB.Where("id = ?", boughtOffer.GetId()).First(&shopOffer)
	logger.LogEvent("Shop", "Purchase", fmt.Sprintf("uid=%d is buying #%d", client.Commander.CommanderID, shopOffer.ID), logger.LOG_LEVEL_INFO)

	response := protobuf.SC_16002{
		Result: proto.Uint32(0),
	}

	if !client.Commander.HasEnoughResource(shopOffer.ResourceID, shopOffer.ResourceNumber*boughtOffer.GetNumber()) {
		logger.LogEvent("Shop", "Purchase", fmt.Sprintf("uid=%d does not have enough resources", client.Commander.CommanderID), logger.LOG_LEVEL_INFO)
		response.Result = proto.Uint32(1)
		return 0, 16002, nil
	}

	response.DropList = make([]*protobuf.DROPINFO, len(shopOffer.Effects))

	switch shopOffer.Type {
	case 1: // bought resources
		for i, resourceId := range shopOffer.Effects {
			client.Commander.AddResource(uint32(resourceId), shopOffer.Number*boughtOffer.GetNumber())
			response.DropList[i] = &protobuf.DROPINFO{
				Type:   proto.Uint32(shopOffer.Type), // ressource
				Id:     proto.Uint32(uint32(resourceId)),
				Number: proto.Uint32(shopOffer.Number * boughtOffer.GetNumber()),
			}
		}
	case 2: // packs
		for i, packId := range shopOffer.Effects {
			client.Commander.AddItem(uint32(packId), shopOffer.Number*boughtOffer.GetNumber())
			response.DropList[i] = &protobuf.DROPINFO{
				Type:   proto.Uint32(shopOffer.Type), // item
				Id:     proto.Uint32(uint32(packId)),
				Number: proto.Uint32(shopOffer.Number * boughtOffer.GetNumber()),
			}
		}
	case 4: // merit shop, to implement
		response.Result = proto.Uint32(3)
	case 6: // skins
		for i, skinId := range shopOffer.Effects {
			client.Commander.GiveSkin(uint32(skinId))
			response.DropList[i] = &protobuf.DROPINFO{
				Type:   proto.Uint32(shopOffer.Type), // skin
				Id:     proto.Uint32(uint32(skinId)),
				Number: proto.Uint32(shopOffer.Number),
			}
		}
	case 12: // operation siren, to implement
	case 20:
		response.Result = proto.Uint32(3)
	default:
		response.Result = proto.Uint32(2)
	}

	if response.GetResult() == 0 {
		logger.LogEvent("Shop", "Purchase", fmt.Sprintf("uid=%d bought #%d successfully!", client.Commander.CommanderID, shopOffer.ID), logger.LOG_LEVEL_INFO)
	}

	client.Commander.ConsumeResource(shopOffer.ResourceID, shopOffer.ResourceNumber)
	return client.SendMessage(16002, &response)
}
