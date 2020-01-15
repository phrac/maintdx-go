package common

import (
	"github.com/go-pg/pg/v9"
	"github.com/spf13/viper"
	"log"
)

var DB *pg.DB

func InitDB() *pg.DB {

	viper.SetConfigFile("common/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: ", err)
	}

	db := pg.Connect(&pg.Options{
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.dbname"),
		Addr:     viper.GetString("database.host") + ":" + viper.GetString("database.port"),
	})

	DB = db
	return DB
}

func GetDB() *pg.DB {
	return DB

}
