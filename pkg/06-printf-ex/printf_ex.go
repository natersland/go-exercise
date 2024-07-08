package printfex

import (
	"fmt"
	"os"
)

type PrintfEx interface {
	PrintYourAge()
	PrintYourNameAndLastName()
	FalseClaims()
	PrintTheTemperature()
	DoubleQuotes()
	PrintTheType()
	PrintTheType2()
	PrintTheType3()
	PrintTheType4()
	PrintYourFullName()
}

type printfEx struct {
}

func (p *printfEx) DoubleQuotes() {
	// ---------------------------------------------------------
	// EXERCISE: Double Quotes
	//
	//  Print "hello world" with double-quotes using Printf
	//  (As you see here)
	//
	// NOTE
	//  Output "shouldn't" be just: hello world
	//
	// EXPECTED OUTPUT
	//  "hello world"
	// ---------------------------------------------------------

	fmt.Printf("\"hello world\"\n")
}

func (p *printfEx) FalseClaims() {
	// ---------------------------------------------------------
	// EXERCISE: False Claims
	//
	//  Use printf to print the expected output using a variable.
	//
	// EXPECTED OUTPUT
	//  These are false claims.
	// ---------------------------------------------------------
	// UNCOMMENT THE FOLLOWING CODE
	// AND DO NOT CHANGE IT AFTERWARDS
	tf := false

	// TYPE YOUR CODE HERE

	fmt.Printf("These are %t claims.\n", tf)
}

func (p *printfEx) PrintTheTemperature() {

	// ---------------------------------------------------------
	// EXERCISE: Print the Temperature
	//
	//	Print the current temperature in your area using Printf
	//
	// NOTE
	//
	//	Do not use %v verb
	//	Output "shouldn't" be like 29.500000
	//
	// EXPECTED OUTPUT
	//
	//	Temperature is 29.5 degrees.
	//
	// ---------------------------------------------------------
	temp := 29.5

	fmt.Printf("Temperature is %.1f degrees.\n", temp)

}

func (p *printfEx) PrintTheType() {
	// ---------------------------------------------------------
	// EXERCISE: Print the Type
	//
	//  Print the type and value of 3 using fmt.Printf
	//
	// EXPECTED OUTPUT
	//  Type of 3 is int
	// ---------------------------------------------------------

	v := 3

	fmt.Printf("Type of 3 is %T\n", v)
}

func (p *printfEx) PrintTheType2() {
	// ---------------------------------------------------------
	// EXERCISE: Print the Type #2
	//
	//  Print the type and value of 3.14 using fmt.Printf
	//
	// EXPECTED OUTPUT
	//  Type of 3.14 is float64
	// ---------------------------------------------------------

	v := 3.14

	fmt.Printf("Type of 3.14 is %T\n", v)
}

func (p *printfEx) PrintTheType3() {

	// ---------------------------------------------------------
	// EXERCISE: Print the Type #3
	//
	//  Print the type and value of "hello" using fmt.Printf
	//
	// EXPECTED OUTPUT
	// 	Type of hello is string
	// ---------------------------------------------------------

	v := "hello"
	fmt.Printf("Type of hello is %T\n", v)
}

func (p *printfEx) PrintTheType4() {
	// ---------------------------------------------------------
	// EXERCISE: Print the Type #4
	//  Print the type and value of true using fmt.Printf
	//
	// EXPECTED OUTPUT
	//  Type of true is bool
	// ---------------------------------------------------------

	v := true
	fmt.Printf("Type of true is %T\n", v)

}

func (p *printfEx) PrintYourAge() {

	// ---------------------------------------------------------
	// EXERCISE: Print Your Age
	//
	//  Print your age using Printf
	//
	// EXPECTED OUTPUT
	//  I'm 30 years old.
	//
	// NOTE
	//  You should change 30 to your age, of course.
	// ---------------------------------------------------------

	age := 29
	yearsStr := "years"

	if age < 1 {
		yearsStr = "year"
	}

	fmt.Printf("I'm %d %s old\n", age, yearsStr)
}

func (p *printfEx) PrintYourFullName() {
	// ---------------------------------------------------------
	// EXERCISE: Print Your Fullname
	//
	//  1. Get your name and lastname from the command-line
	//  2. Print them using Printf
	//
	// EXAMPLE INPUT
	//  Inanc Gumus
	//
	// EXPECTED OUTPUT
	//  Your name is Inanc and your lastname is Gumus.
	// ---------------------------------------------------------
	// BONUS: Use a variable for the format specifier

	var (
		name     string
		lastName string
	)

	args := len(os.Args) - 1
	isValidName := args >= 1
	isValidLastName := args >= 2
	fxName := "PrintYourFullName()"

	if !isValidName {
		fmt.Printf("%s: You don't pass name to args.\n", fxName)
	}
	if !isValidLastName {
		fmt.Printf("%s: You don't pass last name to args.\n", fxName)
	}

	if isValidName && isValidLastName {
		name = os.Args[1]
		lastName = os.Args[2]
		fmt.Printf("Your name is %s and your lastname is %s.\n", name, lastName)
	}

}

func (p *printfEx) PrintYourNameAndLastName() {

	// ---------------------------------------------------------
	// EXERCISE: Print Your Name and LastName
	//
	//	Print your name and lastname using Printf
	//
	// EXPECTED OUTPUT
	//
	//	My name is Inanc and my lastname is Gumus.
	//
	// BONUS
	//
	//	Store the formatting specifier (first argument of Printf)
	//	  in a variable.
	//	Then pass it to printf
	//
	// ---------------------------------------------------------

	name := "Nobi"
	lastName := "Nobita"

	fmt.Printf("My name is %s and my lastname is %s.\n", name, lastName)
}

func NewPrintfEx() PrintfEx {
	return &printfEx{}
}
