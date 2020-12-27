package oacstruct

import (
	"encoding/json"
	"fmt"
)

// User struct
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ExtraUser struct
type ExtraUser struct {
	*User
	Email *struct{} `json:"email,omitempty"`
}

// OACStructSerializer 不修改原结构体的情况下忽略空值字段
//
// 使用匿名嵌套原结构体
//
// 指定需要忽略空值的字段为匿名结构体指针类型
func OACStructSerializer() {

	u1 := User{
		Name: "mauz",
	}
	data, err := json.Marshal(u1)
	if err != nil {
		fmt.Println("OACStructSerializer failed, err: ", err)
	}
	fmt.Printf("OACStructSerializer data: %s\n", data)
}
