package writepractice

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Page 结构体
type Page struct {
	Title string
	Body  []byte
}

// Save 保存文件
func (p *Page) Save() {
	title := p.Title
	body := p.Body

	err := ioutil.WriteFile(title+".txt", body, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}

// Load 根据标题读取对应文件内容
func (p *Page) Load(title string) (string, error) {
	buf, err := ioutil.ReadFile(title)
	if err != nil {
		return "read failed", err
	}

	return string(buf), nil
}

// WriteFromStruct 从结构体中读取数据写入文件
func WriteFromStruct() {
	s := "hello, this is a page"
	body := []byte(s)
	p := &Page{"page", body}
	fmt.Println(p)
	p.Save()
	fmt.Println("save done")

	content, err := p.Load("page.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
}
