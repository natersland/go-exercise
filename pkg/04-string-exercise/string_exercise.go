package stringexercise

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type StringExercise interface {
	WindowPath()
	PrintJSON()
	RawConcat()
	CountTheChars()
	ImprovedBanger()
	ToLowercase()
	TrimIt()
	RightTrimIt()
}

type stringExercise struct {
}

func NewStringExercise() StringExercise {
	return &stringExercise{}
}

func (s *stringExercise) WindowPath() {
	// ---------------------------------------------------------
	// EXERCISE: Windows Path
	//
	//  1. Change the following program
	//  2. It should use a raw string literal instead
	//
	// HINT
	//  Run this program first to see its output.
	//  Then you can easily understand what it does.
	//
	// EXPECTED OUTPUT
	//  Your solution should output the same as this program.
	//  Only that it should use a raw string literal instead.
	// ---------------------------------------------------------
	// HINTS:
	// \\ equals to backslash character
	// \n equals to newline character

	path1 := "c:\\program files\\duper super\\fun.txt\n" +
		"c:\\program files\\really\\funny.png"
	fmt.Println("path1:", path1)

	path2 := `c:\program files\duper super\fun.txt
	c:\program files\really\funny.png`
	fmt.Println("path2:", path2)

}

func (s *stringExercise) PrintJSON() {

	// ---------------------------------------------------------
	// EXERCISE: Print JSON
	//
	//  1. Change the following program
	//  2. It should use a raw string literal instead
	//
	// HINT
	//  Run this program first to see its output.
	//  Then you can easily understand what it does.
	//
	// EXPECTED OUTPUT
	//  Your solution should output the same as this program.
	//  Only that it should use a raw string literal instead.
	// ---------------------------------------------------------
	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	json1 := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"

	fmt.Println("json1:", json1)

	json2 := `
	{
		"Items": [{
			"Item": {
				"name": "Teddy Bear"
			}
		}]
	}
	`

	fmt.Println("json2:", json2)

}

func (s *stringExercise) RawConcat() {
	// ---------------------------------------------------------
	// EXERCISE: Raw Concat
	//
	//  1. Initialize the name variable
	//     by getting input from the command line
	//
	//  2. Use concatenation operator to concatenate it
	//     with the raw string literal below
	//
	// NOTE
	//  You should concatenate the name variable in the correct
	//  place.
	//
	// EXPECTED OUTPUT
	//  Let's say that you run the program like this:
	//   go run main.go inanç
	//
	//  Then it should output this:
	//   hi inanç!
	//   how are you?
	// ---------------------------------------------------------
	// uncomment the code below
	name := "Unknow Name"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// replace and concatenate the `name` variable
	// after `hi ` below

	msg := `hi ` + name + `!
how are you?`

	fmt.Println(msg)
	fmt.Println("original string:", name)
}

func (s *stringExercise) CountTheChars() {
	// ---------------------------------------------------------
	// EXERCISE: Count the Chars
	//
	//  1. Change the following program to work with unicode
	//     characters.
	//
	// INPUT
	//  "İNANÇ"
	//
	// EXPECTED OUTPUT
	//  5
	// ---------------------------------------------------------
	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "İNANÇ", it should return 5 not 7.

	var length int
	if len(os.Args) > 1 {
		length = utf8.RuneCountInString(os.Args[1])

	}

	fmt.Println("string count:", length)
}

func (s *stringExercise) ImprovedBanger() {
	// ---------------------------------------------------------
	// EXERCISE: Improved Banger
	//
	//  Change the Banger program the work with Unicode
	//  characters.
	//
	// INPUT
	//  "İNANÇ"
	//
	// EXPECTED OUTPUT
	//  İNANÇ!!!!!
	// ---------------------------------------------------------

	var msg string
	sr := "You don't pass msg args, Pls pass it for ImprovedBanger()"

	if len(os.Args) > 1 {
		msg = os.Args[1]
		sr = msg + strings.Repeat("!", utf8.RuneCountInString(msg))
	}

	fmt.Println("Improved Banger:", sr)
}

func (s *stringExercise) ToLowercase() {
	// ---------------------------------------------------------
	// EXERCISE: ToLowercase
	//
	//  1. Look at the documentation of strings package
	//  2. Find a function that changes the letters into lowercase
	//  3. Get a value from the command-line
	//  4. Print the given value in lowercase letters
	//
	// HINT
	//  Check out the strings package from Go online documentation.
	//  You will find the lowercase function there.
	//
	// INPUT
	//  "SHEPARD"
	//
	// EXPECTED OUTPUT
	//  shepard
	// ---------------------------------------------------------

	msg := "No input pass from Args for ToLowercase()"

	if len(os.Args) > 1 {
		msg = strings.ToLower(os.Args[1])
	}

	fmt.Println("ToLowercase:", msg)
}

func (s *stringExercise) TrimIt() {
	// ---------------------------------------------------------
	// EXERCISE: Trim It
	//
	//  1. Look at the documentation of strings package
	//  2. Find a function that trims the spaces from
	//     the given string
	//  3. Trim the text variable and print it
	//
	// EXPECTED OUTPUT
	//  The weather looks good.
	//  I should go and play.
	// ---------------------------------------------------------
	msg := `
	
	The weather looks good.
I should go and play.



	`

	trimMsg := strings.TrimSpace(msg)

	fmt.Println("TrimIt:", trimMsg)
}

func (s *stringExercise) RightTrimIt() {
	// ---------------------------------------------------------
	// EXERCISE: Right Trim It
	//
	//  1. Look at the documentation of strings package
	//  2. Find a function that trims the spaces from
	//     only the right-most part of the given string
	//  3. Trim it from the right part only
	//  4. Print the number of characters it contains.
	//
	// RESTRICTION
	//  Your program should work with unicode string values.
	//
	// EXPECTED OUTPUT
	//  5
	// ---------------------------------------------------------

	// currently it prints 17
	// it should print 5

	name := "inanç           "
	trimNameRight := strings.TrimRight(name, " ")
	fmt.Println("TrimRight:", utf8.RuneCountInString(trimNameRight))
}
