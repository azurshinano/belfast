package orm

// 军火商店
type OwnedStressShop struct {
	CommanderID uint32 `gorm:"not_null;primaryKey"`
	ShopOfferID uint32 `gorm:"not_null;primaryKey"`
	Discount    uint32 `gorm:"not_null;default:100"`
	// 可购买次数
	BuyCount uint32 `gorm:"not_null"`

	Commander Commander `gorm:"foreignKey:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ShopOffer ShopOffer `gorm:"foreignKey:ShopOfferID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// 军需商店
type OwnedMilitaryShop struct {
	CommanderID uint32 `gorm:"not_null;primaryKey"`
	ShopOfferID uint32 `gorm:"not_null;primaryKey"`
	// 已购买次数
	BuyCount uint32 `gorm:"not_null"`

	Commander Commander `gorm:"foreignKey:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ShopOffer ShopOffer `gorm:"foreignKey:ShopOfferID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// 舰队商店
type OwnedGuildShop struct {
	CommanderID uint32 `gorm:"not_null;primaryKey"`
	// TODO...
	ShopOfferID uint32 `gorm:"not_null;primaryKey"`
	// 可购买次数
	BuyCount uint32 `gorm:"not_null"`

	Commander Commander `gorm:"foreignKey:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// TODO...
	ShopOffer ShopOffer `gorm:"foreignKey:ShopOfferID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// 舰队商店
type OwnedMedalShop struct {
	CommanderID uint32 `gorm:"not_null;primaryKey"`
	// TODO...
	ShopOfferID uint32 `gorm:"not_null;primaryKey"`
	// 可购买次数
	BuyCount uint32 `gorm:"not_null"`

	Commander Commander `gorm:"foreignKey:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// TODO...
	ShopOffer ShopOffer `gorm:"foreignKey:ShopOfferID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// 奖券商店
type OwnedMiniGameShop struct {
	CommanderID uint32 `gorm:"not_null;primaryKey"`
	// TODO...
	ShopOfferID uint32 `gorm:"not_null;primaryKey"`
	// 可兑换次数
	BuyCount uint32 `gorm:"not_null"`

	Commander Commander `gorm:"foreignKey:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// TODO...
	ShopOffer ShopOffer `gorm:"foreignKey:ShopOfferID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
