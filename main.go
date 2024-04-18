package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"suyuti.com/famtrees/common"
	"suyuti.com/famtrees/model"
	"suyuti.com/famtrees/router"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.Person{})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	db := common.Init()
	MigrateDB(db)

	v1 := r.Group("/api/v1")
	router.Tree(v1.Group("/tree"))

	r.Run()
}
