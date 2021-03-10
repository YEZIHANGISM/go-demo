package main

import (
	"fmt"
	"goto_v1/apis"
	"net/http"
)

// 项目：长URL转换为短URL
// 访问短URL时，自动重定向至对应的长URL
func main() {
	http.HandleFunc("/", apis.Redirect)
	http.HandleFunc("/add", apis.Add)
	fmt.Println("Server start with http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
