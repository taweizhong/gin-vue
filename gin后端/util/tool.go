package util

import (
	"awesomeProject1/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func RandomString(n int) string {
	str := "asdfghjklqwertyuiopzxcvbnm"
	rand.Seed(time.Now().Unix())
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		ch := str[rand.Intn(len(str))]
		s[i] = ch
	}
	return string(s)
}

func IsIphone(db *gorm.DB, num string) bool {
	var user model.User
	db.Where("iphone = ?", num).First(&user)
	fmt.Println(user)
	if user.ID != 0 {
		return true
	}
	return false
}

func IfPassword(db *gorm.DB, num string, password string) model.User {
	var user model.User
	db.Where("iphone = ?", num).First(&user)
	fmt.Println(user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
		return user
	}
	return user
}
