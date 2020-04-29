package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB DataBase
var DB *gorm.DB

// DBConfig DataBase configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig Create DataBase configuration
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		DBName:   "startrek",
		Password: "admin",
	}
	return &dbConfig
}

// URL Return the DataBase configuration as a URL
func URL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
