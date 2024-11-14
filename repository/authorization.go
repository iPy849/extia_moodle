package repository

import "gorm.io/gorm"

type Authorization struct {
	gorm.Model
	ApiKey string `validate:"required,gte=20"`
}
