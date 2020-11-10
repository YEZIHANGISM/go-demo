package main

import (
	"fmt"
	"gin_demo/home"
	"gin_demo/jwt"
	"gin_demo/routers"
)

func main() {
	routers.Include(home.Routers, jwt.Routers)
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("starting server failed, err: ", err)
	}
}
