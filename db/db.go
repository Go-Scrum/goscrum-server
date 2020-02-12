package db

import (
	"fmt"
	"os"

	"goscrum/server/constants"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbClient(logEnabled bool) *gorm.DB {
	connectionString := fmt.Sprintf(`%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True`,
		os.Getenv(constants.DBUsername),
		os.Getenv(constants.DBPassword),
		os.Getenv(constants.DatabaseHostName),
		os.Getenv(constants.DatabaseName))

	fmt.Println(connectionString)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	db.LogMode(logEnabled)
	return db
}
