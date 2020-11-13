package nestedstruct

import (
	"encoding/json"
	"fmt"
)

// Profile nested struct
type Profile struct {
	Web string `json:"web"`
	App string `json:"app"`
}

// annonyS extra struct
type anonyS struct {
	Name string `json:"name"`
	Profile
}

func NestedStructSerializer() {
	// 嵌套结构体序列化
	// 匿名嵌套序列化后所有字段以一维平铺展示

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
