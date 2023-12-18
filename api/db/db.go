package db

import (
	"github.com/alfatahh54/create-transaction/config"
	"github.com/alfatahh54/create-transaction/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbType struct {
	DB *gorm.DB
}

var Database DbType

func init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// for {
	// 	if config.DBConfig != nil {
	// 		break
	// 	}
	// }
	dsn := config.DBConfig.User + ":" + config.DBConfig.Pass + "@tcp(" + config.DBConfig.Host + ":" + config.DBConfig.Port + ")/" + config.DBConfig.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connecting to Database: " + err.Error())
	}
	Database.DB = Db
}
func (Db *DbType) Migrate() {
	Db.DB.AutoMigrate(&model.Product{}, &model.Transaction{})
}
