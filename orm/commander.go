package orm

import (
	"fmt"
	"sort"
	"time"

	"github.com/ggmolly/belfast/logger"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Commander struct {
	CommanderID   uint32    `gorm:"primary_key"`
	AccountID     uint32    `gorm:"not_null"`
	Level         int       `gorm:"default:1;not_null"`
	Exp           int       `gorm:"default:0;not_null"`
	Name          string    `gorm:"size:30;not_null"`
	Adv           string    `gorm:"default:'Belfast';not_null"`
	LastLogin     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not_null"`
	RoomID        uint32    `gorm:"default:0;not_null"`
	ExchangeCount uint32    `gorm:"default:0;not_null"` // Number of times the commander has built ships, can be exchanged for UR ships
	CreatedAt     time.Time

	Punishments    []Punishment        `gorm:"foreignKey:PunishedID;references:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Ships          []OwnedShip         `gorm:"foreignKey:OwnerID;references:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items          []CommanderItem     `gorm:"foreignKey:CommanderID;references:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MiscItems      []CommanderMiscItem `gorm:"foreignKey:CommanderID;references:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OwnedResources []OwnedResource     `gorm:"foreignKey:CommanderID;references:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Builds         []Build             `gorm:"foreignKey:BuilderID;references:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Mails          []Mail              `gorm:"foreignKey:ReceiverID;references:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OwnedSkins     []OwnedSkin         `gorm:"foreignKey:CommanderID;references:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Secretaries    []*OwnedShip        `gorm:"-"`
	OwnedStories   []OwnedStory        `gorm:"foreignKey:CommanderID;references:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OwnedAwards    []OwnedAward        `gorm:"foreignKey:CommanderID;references:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// These maps will be populated by the Load() method
	OwnedShipsMap     map[uint32]*OwnedShip         `gorm:"-"`
	OwnedResourcesMap map[uint32]*OwnedResource     `gorm:"-"`
	CommanderItemsMap map[uint32]*CommanderItem     `gorm:"-"`
	MiscItemsMap      map[uint32]*CommanderMiscItem `gorm:"-"`
	BuildsMap         map[uint32]*Build             `gorm:"-"`
	OwnedSkinsMap     map[uint32]*OwnedSkin         `gorm:"-"`
	MailsMap          map[uint32]*Mail              `gorm:"-"`
}

func (c *Commander) HasEnoughGold(n uint32) bool {
	return c.HasEnoughResource(1, n)
}

func (c *Commander) HasEnoughCube(n uint32) bool {
	return c.HasEnoughItem(20001, n)
}

func (c *Commander) HasEnoughItem(itemId uint32, n uint32) bool {
	if item, ok := c.CommanderItemsMap[itemId]; ok {
		return item.Count >= n
	} else if miscItem, ok := c.MiscItemsMap[itemId]; ok {
		return miscItem.Data >= n
	} else {
		return false
	}
}

func (c *Commander) HasEnoughResource(resourceId uint32, n uint32) bool {
	DealiasResource(&resourceId)
	if resource, ok := c.OwnedResourcesMap[resourceId]; ok {
		return resource.Amount >= n
	} else {
		return false
	}
}

func (c *Commander) CreateBuild(poolId uint32, runningBuilds *int) (protobuf.BUILDINFO, error) {
	var response protobuf.BUILDINFO
	ship, err := GetRandomPoolShip(poolId)
	if err != nil {
		return response, err
	}
	newBuild := Build{
		BuilderID:  c.CommanderID,
		ShipID:     ship.TemplateID,
		FinishesAt: time.Now().Add(time.Second * time.Duration(ship.BuildTime)),
	}
	err = GormDB.Create(&newBuild).Error
	if err != nil {
		return response, err
	}
	response.Time = proto.Uint32(ship.BuildTime)
	response.FinishTime = proto.Uint32(uint32(newBuild.FinishesAt.Unix()))
	response.BuildId = proto.Uint32(uint32(*runningBuilds))
	*runningBuilds++ // the game requires us to send a sequential build id

	// Add the build to the commander's list of BuildsMap
	c.Builds = append(c.Builds, newBuild)
	c.BuildsMap[newBuild.ID] = &newBuild

	return response, nil
}

func (c *Commander) AddShip(shipId uint32) (protobuf.SHIPINFO, error) {
	var response protobuf.SHIPINFO
	var ship Ship
	err := GormDB.Where("template_id = ?", shipId).First(&ship).Error
	if err != nil {
		return response, err
	}
	newShip := OwnedShip{
		ShipID:  ship.TemplateID,
		OwnerID: c.CommanderID,
	}
	tx := GormDB.Create(&newShip)
	if tx.Error != nil {
		return response, tx.Error
	}
	response.Id = proto.Uint32(newShip.ID)
	response.TemplateId = proto.Uint32(newShip.ShipID)
	response.Level = proto.Uint32(1)
	response.Exp = proto.Uint32(0)
	response.Energy = proto.Uint32(100)
	response.State = &protobuf.SHIPSTATE{
		State: proto.Uint32(0),
	}
	response.IsLocked = proto.Uint32(0)
	response.Intimacy = proto.Uint32(5000)
	response.CreateTime = proto.Uint32(uint32(time.Now().Unix()))
	response.SkinId = proto.Uint32(0)
	response.Propose = proto.Uint32(0)
	response.Commanderid = proto.Uint32(0)
	response.MaxLevel = proto.Uint32(0)
	response.BluePrintFlag = proto.Uint32(0)
	response.ActivityNpc = proto.Uint32(0)
	response.Proficiency = proto.Uint32(0)

	// Add the ship to the commander's list of owned ships
	c.Ships = append(c.Ships, newShip)
	c.OwnedShipsMap[newShip.ID] = &newShip
	return response, nil
}

func (c *Commander) ConsumeItem(itemId uint32, count uint32) error {
	// check if the commander has enough of the item
	if item, ok := c.CommanderItemsMap[itemId]; ok {
		if item.Count >= count {
			item.Count -= count
			return GormDB.Save(&item).Error
		}
	} else if miscItem, ok := c.MiscItemsMap[itemId]; ok {
		if miscItem.Data >= count {
			miscItem.Data -= count
			return GormDB.Save(&miscItem).Error
		}
	}
	return fmt.Errorf("not enough items")
}

func (c *Commander) ConsumeResource(resourceId uint32, count uint32) error {
	DealiasResource(&resourceId)
	// check if the commander has enough of the resource
	if resource, ok := c.OwnedResourcesMap[resourceId]; ok {
		if resource.Amount >= count {
			resource.Amount -= count
			return GormDB.Save(&resource).Error
		}
	}
	return fmt.Errorf("not enough resources")
}

func (c *Commander) SetResource(resourceId uint32, amount uint32) error {
	// check if the commander already has the resource, if so set the amount and save
	if resource, ok := c.OwnedResourcesMap[resourceId]; ok {
		resource.Amount = amount
		return GormDB.Save(resource).Error
	}
	// otherwise create a new resource
	newResource := OwnedResource{
		CommanderID: c.CommanderID,
		ResourceID:  resourceId,
		Amount:      amount,
	}
	err := GormDB.Create(&newResource).Error
	if err != nil {
		// append the new resource to the commander's list of resources
		c.OwnedResources = append(c.OwnedResources, newResource)
		c.OwnedResourcesMap[resourceId] = &newResource
	}
	return err
}

func (c *Commander) SetItem(itemId uint32, amount uint32) error {
	// check if the commander already has the item, if so set the amount and save
	if item, ok := c.CommanderItemsMap[itemId]; ok {
		item.Count = amount
		return GormDB.Save(item).Error
	}
	// otherwise create a new item
	newItem := CommanderItem{
		CommanderID: c.CommanderID,
		ItemID:      itemId,
		Count:       amount,
	}
	err := GormDB.Create(&newItem).Error
	if err != nil {
		// append the new item to the commander's list of items
		c.Items = append(c.Items, newItem)
		c.CommanderItemsMap[itemId] = &newItem
	}
	return err
}

func (c *Commander) AddResource(resourceId uint32, amount uint32) error {
	// check if the commander already has the resource, if so increment the amount and save
	if resource, ok := c.OwnedResourcesMap[resourceId]; ok {
		resource.Amount += amount
		return GormDB.Save(resource).Error
	}
	// otherwise create a new resource
	newResource := OwnedResource{
		CommanderID: c.CommanderID,
		ResourceID:  resourceId,
		Amount:      amount,
	}
	err := GormDB.Create(&newResource).Error
	if err != nil {
		// append the new resource to the commander's list of resources
		c.OwnedResources = append(c.OwnedResources, newResource)
		c.OwnedResourcesMap[resourceId] = &newResource
	}
	return err
}

func (c *Commander) AddItem(itemId uint32, amount uint32) error {
	// check if the commander already has the item, if so increment the amount and save
	if item, ok := c.CommanderItemsMap[itemId]; ok {
		item.Count += amount
		return GormDB.Save(item).Error
	}
	// otherwise create a new item
	newItem := CommanderItem{
		CommanderID: c.CommanderID,
		ItemID:      itemId,
		Count:       amount,
	}
	err := GormDB.Create(&newItem).Error
	if err != nil {
		// append the new item to the commander's list of items
		c.Items = append(c.Items, newItem)
		c.CommanderItemsMap[itemId] = &newItem
	}
	return err
}

func (c *Commander) GetItem(itemId uint32) (CommanderItem, error) {
	if item, ok := c.CommanderItemsMap[itemId]; ok {
		return *item, nil
	}
	return CommanderItem{}, fmt.Errorf("item not found")
}

func (c *Commander) GetResource(resourceId uint32) (OwnedResource, error) {
	if resource, ok := c.OwnedResourcesMap[resourceId]; ok {
		return *resource, nil
	}
	return OwnedResource{}, fmt.Errorf("resource not found")
}

// GetItemCount returns the amount of items the commander has, returns 0 if the item is not found
func (c *Commander) GetItemCount(itemId uint32) uint32 {
	if item, ok := c.CommanderItemsMap[itemId]; ok {
		return item.Count
	}
	return 0
}

// GetResourceCount returns the amount of resources the commander has, returns 0 if the resource is not found
func (c *Commander) GetResourceCount(resourceId uint32) uint32 {
	DealiasResource(&resourceId)
	if resource, ok := c.OwnedResourcesMap[resourceId]; ok {
		return resource.Amount
	}
	return 0
}

func (c *Commander) Punish(liftTimestamp *time.Time, permanent bool) error {
	punishment := Punishment{
		PunishedID:    c.CommanderID,
		IsPermanent:   permanent,
		LiftTimestamp: liftTimestamp,
	}
	return GormDB.Create(&punishment).Error
}

func (c *Commander) RevokeActivePunishment() error {
	return GormDB.Where("punished_id = ? AND lift_timestamp IS NULL", c.CommanderID).Delete(&Punishment{}).Error
}

// Load loads the commander's data from the database (ships, items, resources, etc)
func (c *Commander) Load() error {
	err := GormDB.
		Preload(clause.Associations).
		Preload("Ships.Ship").        // force preload the ship's data (might be rolled back later for a lazy load instead and replacement of retire switches to map)
		Preload("Mails.Attachments"). // force preload attachments
		First(c, c.CommanderID).
		Error

	// load ships
	c.OwnedShipsMap = make(map[uint32]*OwnedShip)
	for i, ship := range c.Ships {
		c.OwnedShipsMap[ship.ID] = &c.Ships[i]
	}

	// load resources
	c.OwnedResourcesMap = make(map[uint32]*OwnedResource)
	for i, resource := range c.OwnedResources {
		c.OwnedResourcesMap[resource.ResourceID] = &c.OwnedResources[i]
	}

	// load items
	c.CommanderItemsMap = make(map[uint32]*CommanderItem)
	for i, item := range c.Items {
		c.CommanderItemsMap[item.ItemID] = &c.Items[i]
	}

	// load misc items
	c.MiscItemsMap = make(map[uint32]*CommanderMiscItem)
	for i, item := range c.MiscItems {
		c.MiscItemsMap[item.ItemID] = &c.MiscItems[i]
	}

	// load BuildsMap
	c.BuildsMap = make(map[uint32]*Build)
	for i, build := range c.Builds {
		c.BuildsMap[build.ID] = &c.Builds[i]
	}

	// load skins
	c.OwnedSkinsMap = make(map[uint32]*OwnedSkin)
	for i, skin := range c.OwnedSkins {
		c.OwnedSkinsMap[skin.SkinID] = &c.OwnedSkins[i]
	}

	// load MailsMap
	c.MailsMap = make(map[uint32]*Mail)
	for i, mail := range c.Mails {
		c.MailsMap[mail.ID] = &c.Mails[i]
	}
	return err
}

// Commit saves the commander's data to the database (ships, items, resources, etc)
func (c *Commander) Commit() error {
	return GormDB.Session(&gorm.Session{FullSaveAssociations: true}).Save(c).Error
}

// Get a range of builds (special weird query, probably to save battery on phones)
func (c *Commander) GetBuildRange(minPos, maxPos uint32) ([]Build, error) {
	var builds []Build
	err := GormDB.
		Where("builder_id = ?", c.CommanderID).
		Offset(int(minPos) - 1).
		Limit(int(maxPos) - int(minPos) + 1). // stupid hack to select a range of rows
		Order("id ASC").
		Find(&builds).
		Error
	return builds, err
}

// Bump last login
func (c *Commander) BumpLastLogin() error {
	c.LastLogin = time.Now()
	return GormDB.Save(c).Error
}

func (c *Commander) GetSecretaries() []*OwnedShip {
	if len(c.Secretaries) > 0 {
		return c.Secretaries
	}
	// filter out the ships that are not secretaries
	for i, ship := range c.Ships {
		if ship.IsSecretary {
			c.Secretaries = append(c.Secretaries, &c.Ships[i])
		}
	}
	// Sort for PlayerInfo packet (SC_11003)
	sort.Slice(c.Secretaries, func(i, j int) bool {
		if c.Secretaries[i].SecretaryPosition == nil {
			return false
		}
		return *c.Secretaries[i].SecretaryPosition < *c.Secretaries[j].SecretaryPosition
	})
	return c.Secretaries
}

func (c *Commander) GiveSkin(skinId uint32) error {
	newSkin := OwnedSkin{
		CommanderID: c.CommanderID,
		SkinID:      skinId,
	}
	if err := GormDB.Create(&newSkin).Error; err != nil {
		return err
	}
	c.OwnedSkins = append(c.OwnedSkins, newSkin)
	c.OwnedSkinsMap[skinId] = &newSkin
	return nil
}

func (c *Commander) CleanMailbox() error {
	return GormDB.Where("receiver_id = ?", c.CommanderID).Delete(&Mail{}).Error
}

func (c *Commander) DestroyShips(shipIds []uint32) error {
	return GormDB.Where("owner_id = ? AND id IN ?", c.CommanderID, shipIds).Delete(&OwnedShip{}).Error
}

// Retire a list of ships, return the amount of medals gained, and an error if any
// Data from : https://azurlane.koumakan.jp/wiki/Building#Retire
func (c *Commander) RetireShips(shipIds *[]uint32) error {
	var coins uint32            // given when a ship is retired
	var totalMedals uint32      // given when a Rare (or higher) ship is retired
	var specializedCores uint32 // given when a UR ship is retired
	for _, shipId := range *shipIds {
		ship, ok := c.OwnedShipsMap[shipId]
		if !ok {
			return fmt.Errorf("ship #%d not found", shipId)
		}
		// Give coins
		switch ship.Ship.Type {
		case 1: // destroyer
			coins += 12
		case 2: // light cruiser
			coins += 14
		case 3: // heavy cruiser
			coins += 18
		case 18: // large cruiser
			coins += 19
		case 13: // monitor
			coins += 13
		case 4: // battlecruiser
			coins += 22
		case 5: // battleship
			coins += 26
		case 10: // aviation battleship
			coins += 25
		case 6: // light carrier
			coins += 16
		case 7: // aircraft carrier
			coins += 16
		case 8: // submarine
		case 17: // submarine carrier
			coins += 10
		case 12: // repair ship
			coins += 13
		case 19: // munition ship
			coins += 11
		default:
			return fmt.Errorf("unknown ship type: %d", ship.Ship.Type)
		}

		// give medals / specialized cores
		switch ship.Ship.RarityID {
		case 2: // normal
			totalMedals += 0
		case 3: // rare
			totalMedals += 1
		case 4: // elite
			totalMedals += 4
		case 5: // super rare
			totalMedals += 10
		case 6: // ultra rare
			totalMedals += 30
			specializedCores += 500
		default:
			return fmt.Errorf("unknown ship rarity: %d", ship.Ship.RarityID)
		}
	}
	if err := c.AddResource(1, coins); err != nil {
		return err
	}
	if err := c.AddItem(15001, totalMedals); err != nil {
		return err
	}
	if err := c.AddItem(59010, specializedCores); err != nil {
		return err
	}
	logger.LogEvent("RetireShip", "Success", fmt.Sprintf("uid=%d, coins: %d, medals: %d, cores: %d", c.CommanderID, coins, totalMedals, specializedCores), logger.LOG_LEVEL_INFO)
	return c.DestroyShips(*shipIds)
}

// UpdateRoom changes the commander's room id
func (c *Commander) UpdateRoom(roomID uint32) error {
	return GormDB.Model(c).Update("room_id", roomID).Error
}

// RemoveSecretaries removes all secretaries from the commander
func (c *Commander) RemoveSecretaries() error {
	tx := GormDB.Begin()
	for _, ship := range c.GetSecretaries() {
		ship.IsSecretary = false
		ship.SecretaryPosition = nil
		if err := tx.Save(ship).Error; err != nil {
			return err
		}
	}
	return tx.Commit().Error
}

// UpdateSecretaries changes the commander's secretaries (dirty implementation, but it works)
func (c *Commander) UpdateSecretaries(shipIds []uint32) error {
	tx := GormDB.Begin() // start a transaction to update all at once
	// remove all secretaries
	for _, ship := range c.GetSecretaries() {
		ship.IsSecretary = false
		ship.SecretaryPosition = nil
		if err := tx.Save(ship).Error; err != nil {
			return err
		}
	}
	// add the new secretaries
	for i, shipId := range shipIds {
		ship, ok := c.OwnedShipsMap[shipId]
		if !ok {
			return fmt.Errorf("ship #%d not found", shipId)
		}
		ship.IsSecretary = true
		ship.SecretaryPosition = new(uint32)
		*ship.SecretaryPosition = uint32(i)
		if err := tx.Save(ship).Error; err != nil {
			return err
		}
	}
	return tx.Commit().Error
}

// Add n exchange count to the commander, n represents the number of built ships, caps at 400
func (c *Commander) IncrementExchangeCount(n uint32) error {
	c.ExchangeCount += n
	if c.ExchangeCount > 400 {
		c.ExchangeCount = 400
	}
	return GormDB.Save(c).Error
}

// Likes a ship, inserts a row into the likes table with the ship's group_id
func (c *Commander) Like(groupId uint32) error {
	like := Like{
		GroupID: groupId,
		LikerID: c.CommanderID,
	}
	return GormDB.Create(&like).Error
}

func (c *Commander) AddStory(storyId uint32) error {
	story := OwnedStory{
		CommanderID: c.CommanderID,
		StoryID:     storyId,
	}
	return GormDB.Create(&story).Error
}

func (c *Commander) AddAward(typeId uint32, awardId uint32) error {
	award := OwnedAward{
		CommanderID: c.CommanderID,
		TypeID:      typeId,
		AwardID:     awardId,
	}
	return GormDB.Create(&award).Error
}
