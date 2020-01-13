package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dbConnectionString := BuildConnectionString()
	db, err := gorm.Open("postgres", dbConnectionString)

	if err != nil {
		log.Fatal("DB error: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

func BuildConnectionString() string {
	viper.SetConfigFile("common/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: ", err)
	}
	sslMode := viper.GetString("database.sslmode")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	dbName := viper.GetString("database.dbname")
	password := viper.GetString("database.password")

	c := "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbName + " password=" + password + " sslmode=" + sslMode

	return c
}
