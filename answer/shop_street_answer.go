package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/orm"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
	"time"
)

func ShopStreetAnswer(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_22102{
		Street: &protobuf.SHOPPINGSTREET{
			// 目前只有1, 写死
			Lv:            proto.Uint32(1),
			NextFlashTime: proto.Uint32(uint32(time.Now().Unix() + 112233)),
			// 目前没有升级, 写死
			LvUpTime:   proto.Uint32(0),
			FlashCount: proto.Uint32(0),
		},
	}

	var ownedStressShop []orm.OwnedStressShop
	err := orm.GormDB.Where("commander_id = ?", client.Commander.CommanderID).Find(&ownedStressShop).Error
	if err != nil {
		return 0, 22102, err
	}
	if len(ownedStressShop) == 0 {
		if err := client.Commander.RefreshShopStreet(); err != nil {
			return 0, 22102, err
		}
		err := orm.GormDB.Where("commander_id = ?", client.Commander.CommanderID).Find(&ownedStressShop).Error
		if err != nil {
			return 0, 22102, err
		}
	}
	goods := make([]*protobuf.STREETGOODS, len(ownedStressShop))
	for i, v := range ownedStressShop {
		goods[i] = &protobuf.STREETGOODS{
			GoodsId:  proto.Uint32(v.ShopOfferID),
			Discount: proto.Uint32(v.Discount),
			BuyCount: proto.Uint32(v.BuyCount),
		}
	}
	response.Street.GoodsList = goods

	return client.SendMessage(22102, &response)
}
