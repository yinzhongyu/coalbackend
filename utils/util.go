package utils

import (
	"ginVue/model"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("lkqwjPQWIDLSAKJkwa;dk;jklqjdlk1wzxc,./;KJC;KAS/1231-09")
	res := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}
func TelExists(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
