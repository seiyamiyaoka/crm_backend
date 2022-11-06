package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}
