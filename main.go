package main

import (
	"fmt"
	"nfthook/app/router"
	"nfthook/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.Load(r.Group("/api/v1"))

	r.Run(fmt.Sprintf(":%s", config.Get().Http.Port))
}
