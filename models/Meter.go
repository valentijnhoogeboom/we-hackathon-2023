package models

import "gorm.io/gorm"

type Meter struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Token      string `json:"token"`
	Active     bool   `json:"active"`
	UserID     uint   `json:"user_id"`
	Data       []Data `json:"-"`
	User       User   `json:"-"`
}
