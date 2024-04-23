package orm

type OwnedStory struct {
	CommanderID uint32 `gorm:"not_null;primaryKey"`
	StoryID     uint32 `gorm:"not_null;primaryKey"`

	Commander Commander `gorm:"foreignKey:CommanderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
