package repository

import (
	"errors"
	"extia/database"
	"extia/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	errIncorrectPassword          = errors.New("incorrect password")
	errInvalidPassword            = errors.New("invalid password")
	errNotValidUser               = errors.New("not valid user")
	errCannotCreateIncompleteUser = errors.New("cannot create user without email or password")
)

type User struct {
	gorm.Model
	Email  string `gorm:"unique" validate:"required,email"`
	Hash   string `validate:"required,gte=8"`
	Active bool
}

func (u *User) Validate() error {
	return utils.Validator.Struct(*u)
}

func (u *User) CompareHash(hash string) error {
	hashedPasswordBytes := []byte(u.Hash)
	passwordBytes := []byte(hash)
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
	if err != nil {
		return errIncorrectPassword
	}
	return nil
}

func (u *User) GetById(id uint) error {
	result := database.Db.Take(u, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) GetUserByEmail(email string) error {
	if err := utils.Validator.Var(email, "email"); err != nil {
		return err
	}

	result := database.Db.Take(u, "email = ?", email)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) Create() error {
	if err := u.Validate(); err != nil {
		return errCannotCreateIncompleteUser
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Hash), bcrypt.MinCost)
	if err != nil {
		return errInvalidPassword
	}
	u.Hash = string(hash)

	result := database.Db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) Delete() error {
	result := database.Db.Delete(u, u.ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
