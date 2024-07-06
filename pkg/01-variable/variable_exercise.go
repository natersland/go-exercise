package variable

import (
	"fmt"
	"path"
)

type VariableExercise interface {
	MakeItBlue()
	VarToVar()
	AssignWithExpression()
	FindTheRetangle()
	MultiAssign()
	MultiAssign2()
	MultiShortTime()
	Swapper()
	Swapper2()
	DiscardTheFile()
}

type variableExercise struct {
}

func NewVariableExercise() VariableExercise {
	return &variableExercise{}
}

func (v *variableExercise) MakeItBlue() {

	// ---------------------------------------------------------
	// EXERCISE: Make It Blue
	//
	//  1. Change `color` variable's value to "blue"
	//
	//  2. Print it
	//
	// EXPECTED OUTPUT
	//
	//	blue
	//
	// ---------------------------------------------------------

	color := "green"

	// ADD YOUR CODE BELOW
	color = "blue"

	fmt.Println("color:", color)
}

func (v *variableExercise) VarToVar() {
	// ---------------------------------------------------------
	// EXERCISE: Variables To Variables
	//
	//  1. Change the value of `color` variable to "dark green"
	//
	//  2. Do not assign "dark green" to `color` directly
	//
	//     Instead, use the `color` variable again
	//     while assigning to `color`
	//
	//  3. Print it
	//
	// RESTRICTIONS
	//  WRONG ANSWER, DO NOT DO THIS:
	//  `color = "dark green"`
	//
	// HINT
	//  + operator can concatenate string values
	//
	//  Example:
	//
	//  "h" + "e" + "y" returns "hey"
	//
	// EXPECTED OUTPUT
	//  dark green
	// ---------------------------------------------------------

	color := "green"

	// ADD YOUR CODE BELOW
	color = "dark " + color

	fmt.Println(color)
}

func (v *variableExercise) AssignWithExpression() {
	// ---------------------------------------------------------
	// EXERCISE: Assign With Expressions
	//
	//  1. Multiply 3.14 with 2 and assign it to `n` variable
	//
	//  2. Print the `n` variable
	//
	// HINT
	//  Example: 3 * 2 = 6
	//  * is the product operator (it multiplies numbers)
	//
	// EXPECTED OUTPUT
	//  6.28
	// ---------------------------------------------------------

	// DON'T TOUCH THIS

	// Declares a new float64 variable
	// 0. means 0.0
	n := 0.

	// ADD YOUR CODE BELOW
	n = 3.14 * 2

	// ?

	fmt.Println("Expression: ", n)
}

func (v *variableExercise) FindTheRetangle() {
	// ---------------------------------------------------------
	// EXERCISE: Find the Rectangle's Perimeter
	//
	//  1. Find the perimeter of a rectangle
	//     Its width is  5
	//     Its height is 6
	//
	//  2. Assign the result to the `perimeter` variable
	//
	//  3. Print the `perimeter` variable
	//
	// HINT
	//  Formula = 2 times the width and height
	//
	// EXPECTED OUTPUT
	//  22
	//
	// BONUS
	//  Find more formulas here and calculate them in new programs
	//  https://www.mathsisfun.com/area.html
	// ---------------------------------------------------------
	// UNCOMMENT THE CODE BELOW:

	var (
		perimeter     int
		width, height = 5, 6
	)

	// USE THE VARIABLES ABOVE WHEN CALCULATING YOUR RESULT
	// ADD YOUR CODE BELOW
	perimeter = 2 * (width + height)
	fmt.Println("Retangle: ", perimeter)
}

func (v *variableExercise) MultiAssign() {
	// EXERCISE: Multi Assign
	//
	//  1. Assign "go" to `lang` variable
	//     and assign 2 to `version` variable
	//     using a multiple assignment statement
	//
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  go version 2
	// ---------------------------------------------------------
	// DO NOT TOUCH THIS
	var (
		lang    string
		version int
	)

	// ADD YOUR CODE BELOW
	lang = "go"
	version = 2

	// DO NOT TOUCH THIS
	fmt.Println("version:", lang, "version", version)

}

func (v *variableExercise) MultiAssign2() {
	// ---------------------------------------------------------
	// EXERCISE: Multi Assign #2
	//
	//  1. Assign the correct values to the variables
	//     to match to the EXPECTED OUTPUT below
	//
	//  2. Print the variables
	//
	// HINT
	//  Use multiple Println calls to print each sentence.
	//
	// EXPECTED OUTPUT
	//  Air is good on Mars
	//  It's true
	//  It is 19.5 degrees
	// ---------------------------------------------------------
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	// ADD YOUR CODE BELOW
	planet = "Mars"
	isTrue = true
	temp = 19.5

	fmt.Println("---Weather Report---")
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")
	fmt.Println("--------------------")

}

func (v *variableExercise) MultiShortTime() {

	// ---------------------------------------------------------
	// EXERCISE: Multi Short Func
	//
	// 	1. Declare two variables using multiple short declaration syntax
	//
	//  2. Initialize the variables using `multi` function below
	//
	//  3. Discard the 1st variable's value in the declaration
	//
	//  4. Print only the 2nd variable
	//
	// NOTE
	//  You should use `multi` function
	//  while initializing the variables
	//
	// EXPECTED OUTPUT
	//  4
	// ---------------------------------------------------------
	// ADD YOUR DECLARATIONS HERE

	// THEN UNCOMMENT THE CODE BELOW
	_, b := multi()

	fmt.Println("Multi Short Time:", b)
}

// multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}

func (v *variableExercise) Swapper() {
	// ---------------------------------------------------------
	// EXERCISE: Swapper
	//
	//  1. Change `color` to "orange"
	//     and `color2` to "green" at the same time
	//
	//     (use multiple-assignment)
	//
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  orange green
	// ---------------------------------------------------------
	// UNCOMMENT THE CODE BELOW:

	color, color2 := "red", "blue"
	color, color2 = "orange", "green"

	// ?
	fmt.Println("Swapper:", color, color2)
}
func (v *variableExercise) Swapper2() {
	// ---------------------------------------------------------
	// EXERCISE: Swapper #2
	//
	//  1. Swap the values of `red` and `blue` variables
	//
	//  2. Print them
	//
	// EXPECTED OUTPUT
	//  blue red
	// ---------------------------------------------------------
	// UNCOMMENT THE CODE BELOW:

	red, blue := "red", "blue"
	red, blue = blue, red
	// ?
	fmt.Println("Swapper2:", red, blue)
}

func (v *variableExercise) DiscardTheFile() {
	// ---------------------------------------------------------
	// EXERCISE: Discard The File
	//
	//  1. Print only the directory using `path.Split`
	//
	//  2. Discard the file part
	//
	// RESTRICTION
	//  Use short declaration
	//
	// EXPECTED OUTPUT
	//  secret/
	// ---------------------------------------------------------
	// UNCOMMENT THE CODE BELOW:

	dir, _ := path.Split("secret/file.txt")

	fmt.Println("Dir:", dir)

}
