package orm

type OwnedAppreciate struct {
	CommanderID    uint32 `gorm:"not_null;primaryKey"`
	AppreciateID   uint32 `gorm:"not_null;primaryKey"`
	AppreciateType uint32 `gorm:"not_null;primaryKey"`

	Commander Commander `gorm:"foreignKey:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
