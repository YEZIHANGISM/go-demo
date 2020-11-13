package fmtdatetime

import "time"

// Order struct
type Order struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func FmtDatetimeSerializer() {
	// 格式化时间格式
	// 2006-01-02 12:23:11

	const layout = "2006-01-02 15:04:56"
	return
}
