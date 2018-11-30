package store

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)


func Init(){
	var err error
	db, err = gorm.Open("mysql", "root:multifort@/multifort?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}


}
