package main

import (
	"github.com/gin-contrib/static"
	"github.com/ryannguyen89/wager-service/routers"
	"github.com/ryannguyen89/wager-service/utils"
)

const (
	httpPort = 8080
)

func main() {
	var (
		router = routers.NewRouter()
	)

	router.Use(static.Serve("/docs", static.LocalFile("./swagger", false)))

	utils.GinRunWithGracefulShutdown(router, httpPort)
}
