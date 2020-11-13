package int2float

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Int2FloatDeserializer() {
	// 反序列化保持整数不变为浮点型
	// json的number不区分int或float

	// map[string]interface{} -> json string
	var m = make(map[string]interface{}, 1)
	m["count"] = 1
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("int2FloatDeserializer marshal failed, err: %v\n", err)
	}
	fmt.Printf("int2FloatDeserializer marshal data: %#v\n", string(b))

	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&m2)
	if err != nil {
		fmt.Printf("int2FloatSerializer marshal failed, err: %v\n", err)
		return
	}
	fmt.Printf("int2FloatSerializer marshal data: %v\n", m2["count"])

	// 将得到的m2["count"]转为json.Number后调用Int64()得到int64的值
	count, err := m2["count"].(json.Number).Int64()
	if err != nil {
		fmt.Printf("int2FloatSerializer Int64 failed, err: %v\n", err)
		return
	}
	fmt.Printf("int2FloatSerializer Int64 data type: %T\n", int(count))
}
