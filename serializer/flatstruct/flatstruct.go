package flatstruct

import (
	"encoding/json"
	"fmt"
)

// Person struct
type Person struct {
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
	Weight float64 `json:"-"`               // 序列化时忽略该字段
	Email  string  `json:"email,omitempty"` // 为空时忽略该字段
}

func FlatStructSerializer() {
	// 扁平结构体序列化
	p := Person{
		Name: "ism",
		Age:  13,
	}
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("flatStructSerializer failed, err: ", err)
		return
	}
	fmt.Printf("flatStructSerializer data: %s\n", data)
}
