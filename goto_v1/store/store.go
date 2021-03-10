package store

import (
	"fmt"
	"goto_v1/key"
	"sync"
)

// URLStore URLStore结构体，包括一个锁对象
type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

// Get 获取key对应的url
func (s *URLStore) Get(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.urls[key]
}

// Set 保存一个key-url
func (s *URLStore) Set(key, url string) bool {
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
func (s *URLStore) Put(url string) string {
	for {
		key := key.GenKey(s.Count())
		if s.Set(key, url) {
			return key
		}
	}
}

// NewURLStore 构造URLStore结构体，返回结构体指针
func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]string)}
}

// Store store
func Store() {
	s := NewURLStore()
	ok := s.Set("gog", "https://google.com")
	if ok {
		url := s.Get("gog")
		fmt.Println(url)
	}

	if url := s.Get("test"); url != "" {
		fmt.Println(url)
	} else {
		fmt.Println("url not found")
	}

	count := s.Count()
	fmt.Printf("length of urls: %d\n", count)

}
