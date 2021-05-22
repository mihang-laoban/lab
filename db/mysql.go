package db

import (
	"database/sql"
	"sync"
)

var (
	Db map[string]*sql.DB
	Sync sync.Once
)

func Init(){
	Sync.Do(func() {

	})
}