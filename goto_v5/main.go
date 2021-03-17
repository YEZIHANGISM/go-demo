package main

import (
	"flag"
	"fmt"
	client "goto_v5/proxy"
	"goto_v5/store"
	"net/http"
	"net/rpc"
)

var (
	listenAddr = flag.String("http", ":8080", "http listen address")
	dataFile   = flag.String("file", "store.json", "data store file name")
	hostname   = flag.String("host", "localhost:8080", "host name and port")
	masterAddr = flag.String("master", "", "RPC master address")
	rpcEnabled = flag.Bool("rpc", false, "enable RPC server")
)

// v2: 将数据持久化到磁盘

// v3: 利用并发，将保存和写入磁盘解耦，优化性能

// v4: 用json持久化存储

// v5: 多服务器架构，读写分离，主从服务器之间使用rpc通信

type Store interface {
	Get(url, key *string) error
	Put(key, url *string) error
}

var s Store

func main() {
	flag.Parse()
	if *masterAddr != "" { // slave
		s = client.NewProxyStore(*masterAddr)
	} else {
		s = store.NewURLStore(*dataFile)
	}
	if *rpcEnabled {
		rpc.RegisterName("Store", s)
		rpc.HandleHTTP()
	}
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	fmt.Printf("Server start with http://%s\n", *hostname)
	http.ListenAndServe(*listenAddr, nil)
}

// AddForm 当未指定url时，显示HTML表单
const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

// Add 将url添加到存储结构体中
func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, AddForm)
		return
	}
	var key string
	if err := s.Put(&url, &key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("ADD STORE: %#v\n", s)
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

// Redirect 重定向
func Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ADD STORE: %#v\n", s)
	key := r.URL.Path[1:]
	var url string
	if err := s.Get(&key, &url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
