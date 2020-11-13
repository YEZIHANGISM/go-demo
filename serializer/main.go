package main

import (
	"encoding/json"
	"fmt"
)

func flatStructSerializer() {
	// 扁平结构体序列化
	type Person struct {
		Name   string  `json:"name"`
		Age    int64   `json:"age"`
		Weight float64 `json:"-"`               // 序列化时忽略该字段
		Email  string  `json:"email,omitempty"` // 为空时忽略该字段
	}
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

func nestedStructSerializer() {
	// 嵌套结构体序列化
	// 匿名嵌套序列化后所有字段以一维平铺展示
	type Profile struct {
		Web string `json:"web"`
		App string `json:"app"`
	}

	type anonyS struct {
		Name string `json:"name"`
		Profile
	}

	a1 := anonyS{
		Name: "jhin",
	}
	data, err := json.Marshal(a1)
	if err != nil {
		fmt.Println("nestedStruct2FlatSerializer failed, err: ", err)
		return
	}
	fmt.Printf("nestedStruct2FlatSerializer data: %s\n", data)

	// 具名嵌套或定义嵌套字段的tag，序列化将按照嵌套格式展示
	type namedS struct {
		Name    string `json:"name"`
		Profile `json:"profile"`
		// NestedProfile Profile `json:"profile"`
	}

	n1 := namedS{
		Name: "jinx",
	}
	n1d, err := json.Marshal(n1)
	if err != nil {
		fmt.Println("nestedStruct2NestedSerializer failed, err: ", err)
		return
	}
	fmt.Printf("nestedStruct2NestedSerializer data: %s\n", n1d)

	// 忽略内嵌结构体为空值的字段
	// 需要内嵌结构体指针
	type NestedEmptyS struct {
		Name     string `json:"name"`
		*Profile `json:"profile,omitempty"`
	}
	e1 := NestedEmptyS{
		Name: "ashy",
	}
	e1d, err := json.Marshal(e1)
	if err != nil {
		fmt.Println("nestedStructOmitEmptySerializer failed, err: ", err)
		return
	}
	fmt.Printf("nestedStructOmitEmptySerializer data: %s\n", e1d)
}

func OACStructSerializer() {
	// 不修改原结构体的情况下忽略空值字段
	// 使用匿名嵌套原结构体
	// 指定需要忽略空值的字段为匿名结构体指针类型
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	type ExtraUser struct {
		*User
		Email *struct{} `json:"email,omitempty"`
	}

	u1 := User{
		Name: "mauz",
	}
	data, err := json.Marshal(u1)
	if err != nil {
		fmt.Println("OACStructSerializer failed, err: ", err)
	}
	fmt.Printf("OACStructSerializer data: %s\n", data)
}

func numberOfStringFieldDeserializer() {
	// 字符串类型的数字反序列化
	// 指定string tag
	type Card struct {
		ID    int64   `json:"id,string"`
		Score float64 `json:"score,string"`
	}

	jsonStr := `{"id": "45", "score": "68.5"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr), &c1); err != nil {
		fmt.Println("numberOfStringFieldDeserializer failed, err: ", err)
		return
	}
	fmt.Printf("numberOfStringFieldDeserializer data: %v\n", c1)
}

func main() {
	flatStructSerializer()
	nestedStructSerializer()
	OACStructSerializer()
	numberOfStringFieldDeserializer()
}
