package repository

import (
	"extia/database"
	"extia/utils"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Name   string `gorm:"unique" validate:"required,uri"`
	Key    string `validate:"required,gt=0"`
	UserID uint   `gorm:"not null" validate:"required"`
}

func (d *Domain) Validate() error {
	return utils.Validator.Struct(*d)
}

func (d *Domain) GetById(id uint) error {
	result := database.Db.Take(d, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Domain) DoesApikeyExists(key string) bool {
	result := database.Db.Where(&Domain{Key: key}).First(d)
	return result.Error == nil && result.RowsAffected > 0
}

func (d *Domain) GetAllKeysByUser(userId uint) ([]Domain, error) {
	domains := make([]Domain, 0)
	query := database.Db.Find(&domains, "user_id = ?", userId)
	if query.Error != nil {
		return nil, query.Error
	}

	return domains, nil
}

func (d *Domain) Create() error {
	if err := d.Validate(); err != nil {
		return err
	}

	result := database.Db.Create(d)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Domain) Delete() error {
	result := database.Db.Unscoped().Delete(d, d.ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
