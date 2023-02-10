package db

import (
	"fmt"
	"os"
)

//Create DB
func Exit() {
	os.Exit(0)
}

func Select() {
	fmt.Println("Select...")
}

func Create(name string, dimensions [][]string) {

	err := os.Mkdir(name, os.ModePerm)
	
	if err != nil {

		if (os.IsExist(err)) {
			fmt.Println("Table already exist with this name")
		} else {
			panic(err)
		}
	}
	
}