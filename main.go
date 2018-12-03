package main

import (
	"fmt"
	"github.com/spf13/viper"
	"multifort/log"
)

func main() {
	fmt.Println(viper.Get("db.default"))
	log.InitLogger()
}
