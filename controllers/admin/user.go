package admin

import (
	"log"
	"web/packs/gin"
)

type User struct {
	Base
}

func (this *User) Ulist(c *gin.Context) {
	var users []map[string]string

	this.dbInstance().GetAll("select * from user", &users)

	log.Println(users)

	// var records []map[string]string
	// db.GetAll("select * from game_roles where accname = 'test'", &records)
	// log.Println(records)
}

func (_ *User) Uadd(c *gin.Context) {

}