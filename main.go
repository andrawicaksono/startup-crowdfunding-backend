package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "bwa:bwa123@tcp(127.0.0.1:3306)/startup_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
  _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database is good")
}