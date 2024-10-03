package models

type Song struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Group       string `gorm:"type:varchar(255);not null"`
	Song        string `gorm:"type:varchar(255);not null"`
	ReleaseDate string `gorm:"type:date;not null"`
	Text        string `gorm:"type:text;not null"`
	Link        string `gorm:"type:varchar(255);not null"`
}
