package main

import (
	"fmt"
	multifort "multifort/cmd"
	"os"
)

func main() {
	test := multifort.NewRootCmd()
	if err := test.Execute(); err != nil {
		fmt.Print("-------------")
		fmt.Print(err)
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
