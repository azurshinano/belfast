package misc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ggmolly/belfast/logger"
	"github.com/ggmolly/belfast/orm"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io/ioutil"
)

var (
	dataFn = map[string]func(*gorm.DB) error{
		"Items":      importItems,
		"Buffs":      importBuffs,
		"Ships":      importShips,
		"Skins":      importSkins,
		"Resources":  importResources,
		"Pools":      importPools,
		"BuildTimes": importBuildTimes,
		"ShopOffers": importShopOffers,
	}
	// Golang maps are unordered, so we need to keep track of the order of the keys ourselves
	order = []string{"Items", "Buffs", "Ships", "Skins", "Resources", "Pools", "BuildTimes", "ShopOffers"}
)

func getBelfastData(file string) (*json.Decoder, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("./belfast-data/%s", file))
	if err != nil {
		return nil, err
	}
	return json.NewDecoder(bytes.NewReader(data)), nil
}

// TODO: A lot of code duplication here, could be refactored

func importItems(tx *gorm.DB) error {
	decoder, err := getBelfastData("item_data_statistics.json")
	if err != nil {
		return err
	}
	decoder.Token() // Consume the start of the array '['

	// Decode each elements
	for decoder.More() {
		var item orm.Item
		if err := decoder.Decode(&item); err != nil {
			return err
		} else if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&item).Error; err != nil {
			return err
		}
	}
	return nil
}

func importBuffs(tx *gorm.DB) error {
	decoder, err := getBelfastData("benefit_buff_template.json")
	if err != nil {
		return err
	}
	decoder.Token() // Consume the start of the array '['

	// Decode each elements
	for decoder.More() {
		var buff orm.Buff
		if err := decoder.Decode(&buff); err != nil {
			return err
		} else if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&buff).Error; err != nil {
			return err
		}
	}
	return nil
}

func importShips(tx *gorm.DB) error {
	decoder, err := getBelfastData("ship_data_statistics.json")
	if err != nil {
		return err
	}
	decoder.Token() // Consume the start of the array '['

	// Decode each elements
	for decoder.More() {
		var ship orm.Ship
		if err := decoder.Decode(&ship); err != nil {
			return err
		} else if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&ship).Error; err != nil {
			return err
		}
	}
	return nil
}

func importSkins(tx *gorm.DB) error {
	decoder, err := getBelfastData("ship_skin_template.json")
	if err != nil {
		return err
	}
	decoder.Token() // Consume the start of the array '['

	// Decode each elements
	for decoder.More() {
		var skin orm.Skin
		if err := decoder.Decode(&skin); err != nil {
			return err
		} else if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&skin).Error; err != nil {
			return err
		}
	}
	return nil
}

func importResources(tx *gorm.DB) error {
	decoder, err := getBelfastData("player_resource.json")
	if err != nil {
		return err
	}
	decoder.Token() // Consume the start of the array '['

	// Decode each elements
	for decoder.More() {
		var resource orm.Resource
		if err := decoder.Decode(&resource); err != nil {
			return err
		} else if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resource).Error; err != nil {
			return err
		}
	}
	return nil
}

func importPools(tx *gorm.DB) error {
	decoder, err := getBelfastData("build_pools.json")
	if err != nil {
		return err
	}
	// [{"id": 101451, "pool": 2}, {"id": 702011, "pool": 2}, {"id": 101491, "pool": 2}, {"id": 702031, "pool": 2}...]
	decoder.Token() // Consume the start of the array '['

	// Decode each elements
	for decoder.More() {
		var pool struct {
			ID   uint32 `json:"id"`
			Pool uint32 `json:"pool"`
		}
		if err := decoder.Decode(&pool); err != nil {
			return err
		}

		// Update each ship with the pool
		var ship orm.Ship
		if err := tx.Where("template_id = ?", pool.ID).First(&ship).Error; err != nil {
			return err
		}
		ship.PoolID = proto.Uint32(pool.Pool)
		if err := tx.Save(&ship).Error; err != nil {
			return err
		}
	}
	return nil
}

func importBuildTimes(tx *gorm.DB) error {
	decoder, err := getBelfastData("build_times.json")
	if err != nil {
		return err
	}
	// {"101031": 1380, "101041": 1380, "101061": 1500, "101071": 1500, ...}
	var buildTimes map[string]uint32
	if err := decoder.Decode(&buildTimes); err != nil {
		return err
	}

	// Update each ship with the build time
	for id, time := range buildTimes {
		var ship orm.Ship
		if err := tx.Where("template_id = ?", id).First(&ship).Error; err != nil {
			return err
		}
		ship.BuildTime = time
		if err := tx.Save(&ship).Error; err != nil {
			return err
		}
	}
	return nil
}

func importShopOffers(tx *gorm.DB) error {
	decoder, err := getBelfastData("shop_template.json")
	if err != nil {
		return err
	}
	decoder.Token() // Consume the start of the array '['

	// Decode each elements
	for decoder.More() {
		var offer struct { // decoding to json via a pq.Int64Array is not supported, so we need to decode the effects manually
			orm.ShopOffer
			Effects_ []uint32 `json:"effect_args" gorm:"-"`
		}
		if err := decoder.Decode(&offer); err != nil {
			return err
		}
		// Manually convert the effects to pq.Int64Array
		offer.ShopOffer.Effects = make([]int64, len(offer.Effects_))
		for i, effect := range offer.Effects_ {
			offer.ShopOffer.Effects[i] = int64(effect)
		}
		shopOffer := orm.ShopOffer{
			ID:             offer.ID,
			Effects:        offer.Effects,
			Number:         offer.Number,
			ResourceNumber: offer.ResourceNumber,
			ResourceID:     offer.ResourceID,
			Type:           offer.Type,
		}
		if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&shopOffer).Error; err != nil {
			return err
		}
	}
	return nil
}

// XXX: The database can end in a limbo state if an error occurs while updating the data (e.g. network error, invalid JSON, etc.)
// upon restarting Belfast, the database won't be re-populated because some tables were already populated
// this could be fixed by passing a single transaction to all the data import functions, but requires some refactoring
// due to the way data is being initialized (mix of 'misc' and 'orm' packages)
func UpdateAllData() {
	logger.LogEvent("GameData", "Updating", "Updating all game data.. this may take a while.", logger.LOG_LEVEL_INFO)
	tx := orm.GormDB.Begin()
	for _, key := range order {
		fn := dataFn[key]
		logger.LogEvent("GameData", "Updating", fmt.Sprintf("Updating %s", key), logger.LOG_LEVEL_INFO)
		// defer func() {
		// 	if r := recover(); r != nil {
		// 		logger.LogEvent("GameData", "Updating", fmt.Sprintf("panic occurred while updating %s: %v", key, r), logger.LOG_LEVEL_ERROR)
		// 		tx.Rollback()
		// 	}
		// }()
		if err := fn(tx); err != nil {
			logger.LogEvent("GameData", "Updating", fmt.Sprintf("failed to update %s: %s", key, err.Error()), logger.LOG_LEVEL_ERROR)
			tx.Rollback()
		}
	}
	if err := tx.Commit().Error; err != nil {
		logger.LogEvent("GameData", "Updating", fmt.Sprintf("failed to commit transaction: %s", err.Error()), logger.LOG_LEVEL_ERROR)
	}
}
