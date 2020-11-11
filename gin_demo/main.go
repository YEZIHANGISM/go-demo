package main

import (
	"context"
	"gin_demo/app/home"
	"gin_demo/app/jwt"
	"gin_demo/app/test"
	"gin_demo/app/validate"
	"gin_demo/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	routers.Include(
		home.Routers,
		jwt.Routers,
		validate.Routers,
		test.Routers,
	)
	r := routers.Init()
	// 正常执行
	// if err := r.Run(); err != nil {
	// 	fmt.Println("starting server failed, err: ", err)
	// }

	// 安全关机
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown: ", err)
	}
	log.Println("server exiting")

	// 热重启
	// windows环境没有syscall.SIGUSR1和syscall.SIGUSR2两个信号，会导致报错
	// if err := endless.ListenAndServe(":8080", r); err != nil {
	// 	log.Fatalf("listen: %s\n", err)
	// }
	// log.Println("server exiting")
}
