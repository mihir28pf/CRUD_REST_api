package database

import (
	"golang-crud-rest-api/entities"
	"log"
         "fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open("root:root@tcp(host.docker.local:3306)/crud_demo"), &gorm.Config{})
	if err != nil {
		fmt.Println("\"line no 17 - client go\"")
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	log.Println("Database Migration Completed...")
}
