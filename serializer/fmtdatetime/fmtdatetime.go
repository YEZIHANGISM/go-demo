package fmtdatetime

import (
	"encoding/json"
	"fmt"
	"time"
)

// Order struct
type Order struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

const layout = "2006-01-02 15:04:05"

// MarshalJSON 当某类型实现了该方法，那么这个类型在序列化时会使用该方法
func (o *Order) MarshalJSON() ([]byte, error) {
	type TempOrder Order
	return json.Marshal(struct {
		CreatedAt string `json:"created_at"`
		*TempOrder
	}{
		CreatedAt: o.CreatedAt.Format(layout),
		TempOrder: (*TempOrder)(o),
	})
}

// UnmarshalJSON 当某类型实现了该方法，那么这个类型在反序列化时会使用该方法
func (o *Order) UnmarshalJSON(data []byte) error {
	type TempOrder Order
	ot := struct {
		CreatedAt string `json:"created_at"`
		*TempOrder
	}{
		TempOrder: (*TempOrder)(o),
	}
	if err := json.Unmarshal(data, &ot); err != nil {
		return err
	}
	var err error
	o.CreatedAt, err = time.Parse(layout, ot.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// FmtDatetimeSerializer 自定义格式化时间
func FmtDatetimeSerializer() {
	// 格式化时间格式
	// 2006-01-02 12:23:11
	o1 := Order{
		ID:        1,
		Title:     "ism",
		CreatedAt: time.Now(),
	}

	b, err := json.Marshal(&o1)
	if err != nil {
		fmt.Printf("FmtDateTimeSerializer failed, err: %v\n", err)
		return
	}
	fmt.Printf("FmtDateTimeSerializer data: %s\n", b)

	jsonStr := `{"created_at": "2020-04-05 10:18:20", "id": 13, "title": "ismble"}`
	var o2 Order
	if err := json.Unmarshal([]byte(jsonStr), &o2); err != nil {
		fmt.Printf("FmtDateTimeDeSerializer failed, err: %v\n", err)
		return
	}
	fmt.Printf("FmtDateTimeDeSerializer data: %v\n", o2)
}
