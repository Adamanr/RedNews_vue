package config

import (
	"RedNews_Server/internal/adapters/db/mysql"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
		return
	}
	mysql.StartMySQL()
	fmt.Printf("----MySQL----\n"+
		"User: %v\n"+
		"Password: %v\n"+
		"Host: %v\n"+
		"Database: %v\n", viper.Get("mysql.user"), viper.Get("mysql.password"), viper.GetString("mysql.host"), viper.GetString("mysql.database"))
}
