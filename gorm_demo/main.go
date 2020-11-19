package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     int
	Name   string
	Gender int16
	Age    int
	Status int16
	Remark string
}

func main() {
	db, err := gorm.Open("mysql", "root:5655@(localhost)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&UserInfo{})

	// u1 := UserInfo{1, "ism", 1, 13, 0, "chatting"}
	// u2 := UserInfo{2, "cyan", 1, 13, 0, "waiting"}

	// INSERT
	// db.Create(&u1)
	// db.Create(&u3)

	// READ
	var u = new(UserInfo)
	// db.First(&u) // SELECT * FROM `user_infos` ORDER BY `id` LIMIT 1
	// db.Find(&u, 2) // SELECT * FROM `user_infos` WHERE ID = 2
	db.Find(&u, "remark = ?", "chatting") // SELECT * FROM `user_infos` WHERE `remark = chatting`
	fmt.Printf("%#v\n", u)

	// UPDATE
	db.Model(&u).Update("remark", "singing")
	fmt.Printf("%#v\n", u)
}
