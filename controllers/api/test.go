package api

import (
	"web/models"
	"web/packs/gin"
)

type Test struct {
	Base
}

func (_ *Test)Index(c *gin.Context)  {
	db := new(models.Mysql)

	db.GetRow("select * from user limit 1")

	// var records []map[string]string
	// db.GetAll("select * from game_roles where accname = 'test'", &records)
	// log.Println(records)

	c.JSON(200,gin.H{
		"code"		:	200,
		"message"	:	"success",
	})
}
