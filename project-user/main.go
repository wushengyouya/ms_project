package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/wushengyouya/project-common"
)

func main() {
	r := gin.Default()
	srv.Run(r, "my-project", ":8080")

}
