package orm

type OwnedAward struct {
	CommanderID uint32 `gorm:"not_null;primaryKey"`
	TypeID      uint32 `gorm:"not_null;primaryKey"`
	AwardID     uint32 `gorm:"not_null;primaryKey"`

	Commander Commander `gorm:"foreignKey:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
