package numberofstring

import (
	"encoding/json"
	"fmt"
)

// Card struct
type Card struct {
	ID    int64   `json:"id,string"`
	Score float64 `json:"score,string"`
}

func NumberOfStringFieldDeserializer() {
	// 字符串类型的数字反序列化
	// 指定string tag

	jsonStr := `{"id": "45", "score": "68.5"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr), &c1); err != nil {
		fmt.Println("numberOfStringFieldDeserializer failed, err: ", err)
		return
	}
	fmt.Printf("numberOfStringFieldDeserializer data: %v\n", c1)
}
