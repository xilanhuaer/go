package global

import (
	"fmt"
	"interface/utils"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB         *gorm.DB
	UserSecret = []byte("254%^FuCo610N!3N")
)

const (
	TokenExpireDuration = time.Hour * 24 * 30
)

func Connection(path string) {
	config := utils.ReadYaml(path)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DatabaseName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("error:%v", err)
	}
}
