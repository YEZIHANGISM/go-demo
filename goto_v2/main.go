package main

import (
	"fmt"
	"goto_v1/apis"
	"net/http"
)

// 将数据持久化到磁盘
func main() {
	http.HandleFunc("/", apis.Redirect)
	http.HandleFunc("/add", apis.Add)
	fmt.Println("Server start with http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
