package entities

type Talent struct {
	BaseModel
	Email      string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Name       string `gorm:"type:varchar(50);not null"`
	Phone      string `gorm:"type:varchar(20);not null"`
	IsVerified bool   `gorm:"default:false"`
}
