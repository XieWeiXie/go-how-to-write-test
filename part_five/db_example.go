package part_five

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Information struct {
	gorm.Model
	Name string `gorm:"type:varchar" json:"name"`
}

type OtherInformation struct {
	gorm.Model
	Name string `gorm:"type:varchar" json:"name"`
}

var (
	POSTGRES *gorm.DB
)

func DBInit() {
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "root"
	dbName := "gorm_example"
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
	POSTGRES.AutoMigrate(&Information{}, &OtherInformation{})
}

func InsertByFirstOrInit(id uint) {
	var infor Information
	infor.Name = strconv.Itoa(int(id)) + " hello"
	if dbError := POSTGRES.FirstOrInit(&infor, map[string]interface{}{"name": strconv.Itoa(int(id)) + " hello"}).Error; dbError != nil {
		fmt.Println(dbError, id)
	}
	POSTGRES.Save(&infor)
}

func InsertByFirstOrCreate(id uint) {
	var infor OtherInformation
	POSTGRES.FirstOrCreate(&infor, OtherInformation{Name: strconv.Itoa(int(id)) + " world"})
}
