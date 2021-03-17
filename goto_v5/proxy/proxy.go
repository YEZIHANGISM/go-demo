package proxy

import (
	"fmt"
	"goto_v5/store"
	"log"
	"net/rpc"
)

// 代表RPC客户端的类型，转发请求到RPC服务器
// 包含一个urls副本，用于缓存
type ProxyStore struct {
	urls   *store.URLStore
	client *rpc.Client
}

// NewProxyStore 创建ProxyStore对象
func NewProxyStore(addr string) *ProxyStore {
	// 连接服务器
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Println("Error constructing ProxyStore: ", err)
	}
	return &ProxyStore{urls: store.NewURLStore(""), client: client}
}

// Get 客户端的Get方法，通过rpc调用服务器的Get方法

// 首先检查缓存中是否有对应的key，如果有则返回

// 如果没有则发起RPC调用，将结果更新至ProxyStore
func (s *ProxyStore) Get(key, url *string) error {
	if err := s.urls.Get(key, url); err == nil {
		return nil
	}
	if err := s.client.Call("Store.Get", key, url); err != nil {
		return err
	}
	fmt.Println("PROXY ADD", &url, &key)
	s.urls.Set(key, url)
	return nil
}

// Put 客户端的Put方法，通过rpc调用服务器的Put方法
func (s *ProxyStore) Put(url, key *string) error {
	if err := s.client.Call("Store.Put", url, key); err != nil {
		return err
	}
	fmt.Println("PROXY ADD", &url, &key)
	s.urls.Set(key, url)
	return nil
}
