package lib

import "fmt"

func Sum(num1 int, num2 int) {
	fmt.Println("Result is: ", num1 + num2)
}

func Sub(num1 int, num2 int)  {
	fmt.Println("Result is: ", num1 - num2)
}

func Div(num1 int, num2 int) {
	fmt.Println("Result is: ", num1 / num2)
}

func Mul(num1 int, num2 int) {
	fmt.Println("Result is: ", num1 * num2)
}

func GetString() string{

	var input string
	fmt.Scanln(&input)

	return input
}

func GetNumber() int{

	fmt.Print("Enter a number: ")
	var num int
	fmt.Scanln(&num)

	return num
}

