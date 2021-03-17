package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"goto_v5/key"
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
			s.Set(&r.Key, &r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	log.Println("Error decoding URLStore: ", err)
	return err
}

// Get 获取key对应的url
func (s *URLStore) Get(k, url *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if u, ok := s.urls[*k]; ok {
		*url = u
		return nil
	}
	return errors.New("key not found")
}

// Set 保存一个key-url
func (s *URLStore) Set(key, url *string) error {
	fmt.Printf("SET s")
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[*key]; present {
		return errors.New("key already exists")
	}
	s.urls[*key] = *url
	vKey := *key
	vURL := *url
	fmt.Printf("SETed %s %s\n", vKey, vURL)
	return nil
}

// Count 计算urls数量
func (s *URLStore) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.urls)
}

// Put 根据url生成一个短url，并保存

// 将record写入带缓冲的save通道

// 修改函数签名以符合RPC调用的形式

// RPC调用方法签名：func (t T) Name(args *ArgType, reply *ReplyType) error
func (s *URLStore) Put(url, k *string) error {
	for {
		*k = key.GenKey(s.Count())
		if err := s.Set(k, url); err == nil {
			break
		}
	}
	if s.save != nil {
		s.save <- record{*k, *url}
	}
	return nil
}

// NewURLStore 构造URLStore结构体，返回结构体指针

// 接收一个filename，打开本地文件，将返回的文件作为file字段存储在URLStore中

// 调用load，将记录加载到URLStore中

// 如果filename为空，则不写入任何文件
func NewURLStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string)}
	if filename != "" {
		s.save = make(chan record, saveQueueLength)
		if err := s.load(filename); err != nil {
			log.Println("Error loading data in URLStore: ", err)
		}
		go s.saveloop(filename)
	}
	return s
}
