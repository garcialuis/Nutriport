package models

type CarbLevel struct {
	ID uint8 `json:"id" gorm:"primary_key:auto_increment"`

	Description string `json:"description" gorm:"size:255;not null"`
}
