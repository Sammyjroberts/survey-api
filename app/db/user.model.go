package db

import (
	"crypto/rand"
	"crypto/sha512"

	"bytes"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/pbkdf2"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `json:"email"`
	Hash      []byte `gorm:"type:BLOB"json:"-"`
	Salt      []byte `json:"-"`
	Password  string `gorm:"-"json:"-"`
}

func (user *User) RegisterAccount() {
	//do work
	user.setHash()
	user.setSalt()

	//create user and defer closing of r
	DB.Create(&user)
}
func (user *User) Login() bool {
	var trueUser User
	DB.Where("email = ?", user.Email).First(&trueUser)
	testHash := pbkdf2.Key([]byte(user.Password), []byte(user.Salt), 128, 256, sha512.New)
	isEqual := bytes.Equal(trueUser.Hash, testHash)
	return isEqual
}

//generate a hash
func (u *User) setHash() {
	u.Hash = pbkdf2.Key([]byte(u.Password), []byte(u.Salt), 128, 256, sha512.New)
}

//generate a salt
func (u *User) setSalt() {
	salt := make([]byte, 32)
	rand.Read(salt)
	u.Salt = salt
}
