package test

import (
	. "dev-go-base/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestServer(t *testing.T){
	v := ConfigBuilder("config", "toml", "../conf/")
	if err := v.Unmarshal(&C); err != nil {
		panic(C)
	}

	Dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		C.Db.User,
		C.Db.Password,
		C.Db.Host,
		C.Db.Port,
		C.Db.Database,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: Dsn, // data source name
		DefaultStringSize: 256, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	fmt.Println(product)

	// Delete - delete product
	db.Delete(&product, 1)
}