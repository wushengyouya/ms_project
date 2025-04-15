package common

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// 其他模块用到启停抽象出公共启停
func Run(r *gin.Engine, serverName string, addr string, stop func()) {
	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 便于后续优雅的启动
	go func() {
		log.Printf(" %s running in %s \n", serverName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	// SIGINT 用户发送INTR字符(CTRL+C)触发
	// SIGTEM 结束程序(可以被捕获，阻塞，忽略)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("shutdown %s web server... \n", serverName)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Microsecond)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("%s web server shutdown,cause by: %v\n", serverName, err)
	}
	if stop != nil {
		stop()
		log.Printf("grpc service stop")
	}
	select {
	case <-ctx.Done():
		log.Println("关闭超时")
	}
	log.Printf("%s web server stop success", serverName)
}
