package utils

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"os"
	"strconv"
	"time"
)

type DbConfig struct {
	User     string
	Password string
	Addr     string
	DBName   string
}

var (
	maxIdle  = 100
	maxConn  = 200
	lifeTime = 300
)

// To make this helper create tables, set environment variable "DB_MIGRATE" to true
func InitDatabase(my DbConfig) {
	var connectString = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&collation=utf8mb4_general_ci&loc=%s",
		my.User, my.Password, my.Addr, my.DBName, "Asia%2fBangkok")
	initDBWithConnectString(connectString)
}

func initDBWithConnectString(connectStr string) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	log.Printf("Connect database with connect string: %s", connectStr)

	if err := orm.RegisterDataBase("default", "mysql",
		connectStr, maxIdle, maxConn);
		err != nil {
		log.Panic("Error when RegisterDataBase: ", err)
	}
	//
	db, err := orm.GetDB("default")
	if err != nil {
		log.Fatal("Get default DB error:" + err.Error())
	}
	db.SetConnMaxLifetime(time.Duration(lifeTime) * time.Second)
	orm.Debug, orm.DebugLog = false, orm.NewLog(os.Stdout)
	if dbMigrate := os.Getenv("DB_MIGRATE"); len(dbMigrate) == 0 {
		log.Printf("Init database without migrate, config: %#v", connectStr)
	} else {
		if isMigrate, _ := strconv.ParseBool(dbMigrate); isMigrate {
			if err := orm.RunSyncdb("default", false, false); err != nil {
				log.Print("Error when migrate db", err)
			}
			log.Printf("Init database migrate mode, config: %#v", connectStr)
		}
	}
}

