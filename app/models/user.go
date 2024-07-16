package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FistName string `gorm:"type:varchar(30)" json:"fist_name"`
	LastName string `gorm:"type:varchar(30)" json:"last_name"`
	Email    string `gorm:"type:varchar(100);unique;not_null" json:"email"`
	Password string `gorm:"type:varchar(255);not_null" json:"-"`
}

func NewUser() User {
	return User{}
}

type RegisterBody struct {
	FistName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (b *RegisterBody) MapTo(p any) {
	u := p.(*User)
	u.FistName = b.FistName
	u.LastName = b.LastName
	u.Email = b.Email
	u.Password = b.Password
}

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (b *LoginBody) MapTo(p any) {
	u := p.(*User)
	u.Email = b.Email
	u.Password = b.Password
}

func (u *User) FindByID(id uint) (*User, error) {
	var user User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (u *User) Create(user User) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	err = db.Create(&user).Error

	return &user, err
}

func (u *User) Authenticate(user User) (*User, bool) {
	var foundUser User
	if err := db.First(&foundUser, "email = ?", user.Email).Error; err != nil {
		return nil, false
	}

	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return nil, false
	}

	return &foundUser, true
}
