package mysql

import (
	"RedNews_Server/internal/adapters/db"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var cfg *mysql.Config

func StartMySQL() {
	var conf db.Database
	viper.UnmarshalKey("mysql", &conf)

	cfg = &mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Password,
		Net:                  "tcp",
		Addr:                 conf.Host,
		DBName:               conf.Database,
		AllowNativePasswords: true,
	}
}

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil
	}
	return db
}
