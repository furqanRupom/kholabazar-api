package db

import (
	"fmt"
    "kholabazar/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(conf *config.DBConfig) string {
	conStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",conf.User,conf.Password,conf.Host,conf.Port,conf.Name)
	if !conf.EnableSSLMode {
		conStr += " sslmode=disable"
	}
	return conStr
}

func NewConnection(conf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(conf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
