package models

import "gorm.io/gorm"

type Cars struct {
	gorm.Model
	UUID    int    `json:"uuid"`
	License string `json:"License" gorm:"unique"`
	Year    int    `json:"year"`
}
