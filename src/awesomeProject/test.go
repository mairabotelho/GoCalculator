package main

import (
	New "awesomeProject/lib"
	"fmt"
)

func main() {

	fmt.Println("Welcome to the Calculator!!")
	fmt.Print("Enter an action: ")
	print(New.GetString())

}

func print(input string){

	switch input {
	case "sum":
		New.Sum(New.GetNumber(), New.GetNumber())

	case "sub":
		New.Sub(New.GetNumber(), New.GetNumber())

	case "div":
		New.Div(New.GetNumber(), New.GetNumber())

	case "mul":
		New.Mul(New.GetNumber(), New.GetNumber())

	default:
		fmt.Println("Not a valid action")
	}
}




