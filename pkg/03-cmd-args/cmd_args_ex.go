package cmdargs

import (
	"fmt"
	"os"
)

type CommandLineArgs interface {
	CountTheArgs()
	PrintThePath()
	PrintYourName()
	GreetMorePeople()
	GreetFivePeople()
}

type commandLineArgs struct {
}

func NewCommandLineArgs() CommandLineArgs {
	return &commandLineArgs{}
}

func (c *commandLineArgs) CountTheArgs() {
	// ---------------------------------------------------------
	// EXERCISE: Count the Arguments
	//
	//  Print the count of the command-line arguments
	//
	// INPUT
	//  bilbo balbo bungo
	//
	// EXPECTED OUTPUT
	//  There are 3 names.
	// ---------------------------------------------------------
	// UNCOMMENT & FIX THIS CODE
	count := len(os.Args) - 1

	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	fmt.Printf("CountArgs: There are %d names.\n", count)
}

func (c *commandLineArgs) PrintYourName() {
	// ---------------------------------------------------------
	// EXERCISE: Print Your Name
	//
	//  Get it from the first command-line argument
	//
	// INPUT
	//  Call the program using your name
	//
	// EXPECTED OUTPUT
	//  It should print your name
	//
	// EXAMPLE
	//  go run main.go inanc
	//
	//    inanc
	//
	// BONUS: Make the output like this:
	//
	//  go run main.go inanc
	//    Hi inanc
	//    How are you?
	// ---------------------------------------------------------
	// get your name from the command-line
	// and print it
	var name string

	isNotHaveArgs := len(os.Args) == 1
	if isNotHaveArgs {
		name = "Unknow Name"
	} else {
		name = os.Args[1]
	}

	fmt.Println("Hi", name)
	fmt.Println("How are you?")
}

func (c *commandLineArgs) PrintThePath() {
	// ---------------------------------------------------------
	// EXERCISE: Print the Path
	//
	//  Print the path of the running program by getting it
	//  from `os.Args` variable.
	//
	// HINT
	//  Use `go build` to build your program.
	//  Then run it using the compiled executable program file.
	//
	// EXPECTED OUTPUT SHOULD INCLUDE THIS
	//  myprogram
	// ---------------------------------------------------------
	path := os.Args[0]
	fmt.Println("The path of the running program is:", path)
}

func (c *commandLineArgs) GreetMorePeople() {
	// ---------------------------------------------------------
	// EXERCISE: Greet More People
	//
	// RESTRICTIONS
	//  1. Be sure to match the expected output below
	//  2. Store the length of the arguments in a variable
	//  3. Store all the arguments in variables as well
	//
	// INPUT
	//  bilbo balbo bungo
	//
	// EXPECTED OUTPUT
	//  There are 3 people!
	//  Hello great bilbo !
	//  Hello great balbo !
	//  Hello great bungo !
	//  Nice to meet you all.
	// ---------------------------------------------------------
	// TYPE YOUR CODE HERE
	// args validation
	people1, people2, people3 := "-", "-", "-"

	args := os.Args
	argsCount := len(args)
	peopleCount := argsCount - 1
	isReachPeopleLimit := peopleCount > 3

	if peopleCount > 0 {
		people1 = args[1]
	}
	if peopleCount > 1 {
		people2 = args[2]
	}
	if peopleCount > 2 {
		people3 = args[3]
	}

	if isReachPeopleLimit {
		fmt.Println("You reached the argument limit. Please pass no more than 3 arguments.")
	}

	fmt.Println("There are", peopleCount, "people!")
	fmt.Println("Hello great", people1)
	fmt.Println("Hello great", people2)
	fmt.Println("Hello great", people3)
	fmt.Println("Nice to meet you all.")
}

func (c *commandLineArgs) GreetFivePeople() {
	// ---------------------------------------------------------
	// EXERCISE: Greet 5 People
	//
	//  Greet 5 people this time.
	//
	//  Please do not copy paste from the previous exercise!
	//
	// RESTRICTION
	//  This time do not use variables.
	//
	// INPUT
	//  bilbo balbo bungo gandalf legolas
	//
	// EXPECTED OUTPUT
	//  There are 5 people!
	//  Hello great bilbo !
	//  Hello great balbo !
	//  Hello great bungo !
	//  Hello great gandalf !
	//  Hello great legolas !
	//  Nice to meet you all.
	// ---------------------------------------------------------
	// TYPE YOUR CODE HERE
	fmt.Println("There are", len(os.Args)-1, "people!")

	if len(os.Args)-1 > 0 {
		fmt.Println("Hello great", os.Args[1])
	}

	if len(os.Args)-1 > 1 {
		fmt.Println("Hello great", os.Args[2])
	}

	if len(os.Args)-1 > 2 {
		fmt.Println("Hello great", os.Args[3])
	}

	if len(os.Args)-1 > 3 {
		fmt.Println("Hello great", os.Args[4])
	}

	if len(os.Args)-1 > 4 {
		fmt.Println("Hello great", os.Args[5])
	}

	if len(os.Args)-1 > 5 {
		fmt.Println("You reached the argument limit. Please pass no more than 5 arguments.")
	}

}
