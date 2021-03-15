package store

import (
	"encoding/json"
	"fmt"
	"goto_v4/key"
	"io"
	"log"
	"os"
	"sync"
)

const saveQueueLength = 1000

// URLStore URLStore结构体

// 包括一个锁对象，和一个record类型的通道
type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	save chan record
}

// 这里的字段必须设置为大写，否则重启后会导致找不到
type record struct {
	Key, URL string
}

// save 将给定的key和URL组成record，以gob编码的形式写入磁盘
// func (s *URLStore) save(key, url string) error {
// 	e := gob.NewEncoder(s.file)
// 	return e.Encode(record{key, url})
// }

// saveloop 保存至文件
func (s *URLStore) saveloop(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore: ", err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	for {
		r := <-s.save
		if err := e.Encode(&r); err != nil {
			log.Println("URLStore:", err)
		}
	}
}

// load 程序启动时，将磁盘上的数据读取到URLStore中
func (s *URLStore) load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening URLStore: ", err)
	}
	defer f.Close()
	d := json.NewDecoder(f)
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			s.Set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	log.Println("Error decoding URLStore: ", err)
	return err
}

// Get 获取key对应的url
func (s *URLStore) Get(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.urls[key]
}

// Set 保存一个key-url
func (s *URLStore) Set(key, url string) bool {
	fmt.Println(key, url)
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[key]; present {
		return true
	}
	s.urls[key] = url
	return true
}

// Count 计算urls数量
func (s *URLStore) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.urls)
}

// Put 根据url生成一个短url，并保存

// 将record写入带缓冲的save通道
func (s *URLStore) Put(url string) string {
	for {
		key := key.GenKey(s.Count())
		if s.Set(key, url) {
			s.save <- record{key, url}
			return key
		}
	}
}

// NewURLStore 构造URLStore结构体，返回结构体指针

// 接收一个filename，打开本地文件，将返回的文件作为file字段存储在URLStore中

// 调用load，将记录加载到URLStore中
func NewURLStore(filename string) *URLStore {
	s := &URLStore{
		urls: make(map[string]string),
		save: make(chan record, saveQueueLength),
	}
	if err := s.load(filename); err != nil {
		log.Println("Error loading data in URLStore: ", err)
	}
	go s.saveloop(filename)
	return s
}
