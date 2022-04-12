package models

type Zone struct {
	ID        uint `gorm:"primaryKey"`
	ShortName string
	LongName string
}
