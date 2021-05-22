package db

import "C"
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	Cons   map[string]*gorm.DB
	Sync sync.Once
)

type DbConfig struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewDb(dbName string) *gorm.DB {
	return Cons[dbName]
}

func InitConfig(dsn string)  {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // data source name
		DefaultStringSize: 256, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Cons["d"] = db
}

func Init(){
	Sync.Do(func() {

	})
}