package store

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	db *gorm.DB
)


func Init(){
	var err error

	db, err = gorm.Open("mysql", viper.GetString("db.default"))
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}


}
