package main

import (
	"flag"
	"fmt"
	"goto_v4/apis"
	"net/http"
)

var (
	listenAddr = flag.String("http", ":8080", "http listen address")
	dataFile   = flag.String("file", "store.gob", "data store file name")
	hostname   = flag.String("host", "localhost:8080", "host name and port")
)

// v2: 将数据持久化到磁盘

// v3: 利用并发，将保存和写入磁盘解耦，优化性能

// v4: 用json持久化存储

func main() {
	flag.Parse()
	http.HandleFunc("/", apis.Redirect)
	http.HandleFunc("/add", apis.Add)
	fmt.Printf("Server start with http://%s\n", *hostname)
	http.ListenAndServe(*listenAddr, nil)
}
