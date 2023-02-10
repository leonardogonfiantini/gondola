package lexer

import (
	"fmt"
	"gondola/db"
	"strings"
)

/************************************************

Examle of command: 

select dimension from table where dimension = 1
create table (dimension1_type dimension2_type)
insert table (val1,val2)
exit

*************************************************/

func Lexer(input string) {
	

	command := strings.Split(input, " ")

	switch command[0] {
	
	case "exit":
		db.Exit()
	
	case "select":
		db.Select()
	
	case "create":
		name := command[1]
		dimensions := formatDimensions(command[2:])
		
		db.Create(name, dimensions)

	default:
		fmt.Println("err: unrecognized key at the beginning of the command")
	}

}


//format the dimensions in a map style
func formatDimensions(input []string) [][]string {

	size := len(input)-1

	input[0] = input[0][1:]
	input[size] = input[size][:len(input[size])-1]
	
	DimsType := make([][]string, size+1)
	for i := range DimsType {
		DimsType[i] = make([]string, 2)
	}

	for i, e := range input {
		dim := strings.Split(e, "_")
		
		DimsType[i][0] = dim[0]
		DimsType[i][1] = dim[1]
	}

	return DimsType

}