package apis

import (
	"fmt"
	"goto_v1/store"
	"net/http"
)

var s = store.NewURLStore("store.gob")

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
	key := s.Put(url)
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

// Redirect 重定向
func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := s.Get(key)
	fmt.Println(url)
	fmt.Println(s.Count())
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
