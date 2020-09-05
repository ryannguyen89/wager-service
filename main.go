package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/gin-contrib/static"
	"github.com/go-sql-driver/mysql"
	"github.com/ryannguyen89/wager-service/models"
	"github.com/ryannguyen89/wager-service/routers"
	"github.com/ryannguyen89/wager-service/utils"
)

// Against unused import
var (
	_  mysql.Config
)

const (
	httpPort = 8080
)

func init() {
	registerOrm()
	initDB()
}

func registerOrm() {
	for _, m := range [...]interface{}{
		new(models.Wager),
		new(models.WagerPurchase),
	} {
		orm.RegisterModel(m)
	}
}

func initDB() {
	utils.InitDatabase(utils.DbConfig{
		User:     "root",
		Password: "root",
		Addr:     "172.17.0.1:3306",
		DBName:   "wager",
	})
}

func main() {
	var (
		router = routers.NewRouter()
	)

	router.Use(static.Serve("/docs", static.LocalFile("./swagger", false)))

	utils.GinRunWithGracefulShutdown(router, httpPort)
}
