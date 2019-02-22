package partfour

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var POSTGRES *gorm.DB

func DBInit() *gorm.DB {
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := ""
	dbName := "gin_example"
	sslMode := "disable"

	connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode)
	conn, err := gorm.Open("postgres", connectString)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	conn.LogMode(true)

	conn.DB().SetMaxIdleConns(3)

	POSTGRES = conn
	return POSTGRES

}
