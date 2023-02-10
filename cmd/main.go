package main

import (
	"gondola/lexer"

	"bufio"
	"fmt"
	"os"
)

func getInput() string {

	_, err := fmt.Print("$~ ")
	if err != nil {
		panic(err)
	}

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	return input.Text()
}

func main() {

	err := os.Chdir("../db/storage")
	if err != nil {
		panic(err)
	}

	fmt.Print("Welcome in gondola, a project for understanding \nin deep how a database works by implementing it \n\n")
	
	for {
		//Get the input by the user
		input := getInput()

		//Call the lexer
		lexer.Lexer(input)

	}


}