package main

import (
	multifort "multifort/cmd"
	"os"
)

func main() {
	test := multifort.NewRootCmd()
	if err := test.Execute(); err != nil {
		os.Exit(-1)
	}
}
