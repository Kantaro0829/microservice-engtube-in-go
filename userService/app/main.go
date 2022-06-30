package main

import (
	"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
	"github.com/Kantaro0829/microservice-engtube-in-go/userService/infrastructure"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
	dsn = "root:ecc@tcp(user-db:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	dbinit()
	infrastructure.Init()

	router := gin.Default()
	router.Run(":3000")
}

func dbinit() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	db.Migrator().CreateTable(domain.User{})
}
