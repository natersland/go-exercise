package ifstatementex

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type IfStatementEx interface {
	AgeSeasons(age int)
	SimplifyIt()
	ArgCount()
	VowelOrConsonant()
	Challenge1()
	Challenge2()
}

type ifStatementEx struct {
}

func NewIfStatementEx() IfStatementEx {
	return &ifStatementEx{}
}

func (i *ifStatementEx) AgeSeasons(age int) {
	// ---------------------------------------------------------
	// EXERCISE: Age Seasons
	//
	//  Let's start simple. Print the expected outputs,
	//  depending on the age variable.
	//
	// EXPECTED OUTPUT
	//  If age is greater than 60, print:
	//    Getting older
	//  If age is greater than 30, print:
	//    Getting wiser
	//  If age is greater than 20, print:
	//    Adulthood
	//  If age is greater than 10, print:
	//    Young blood
	//  Otherwise, print:
	//    Booting up
	// ---------------------------------------------------------

	isGettingOlder := age > 60
	isGettingWiser := age > 30
	isAdulthood := age > 20
	isYoungBlood := age > 10

	// Type your if statement here.
	fmt.Println("AgeSeasons():")
	if isGettingOlder {
		fmt.Println("Getting older")
	} else if isGettingWiser {
		fmt.Println("Getting wiser")
	} else if isAdulthood {
		fmt.Println("Adulthood")
	} else if isYoungBlood {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}

}

func (i *ifStatementEx) SimplifyIt() {
	// ---------------------------------------------------------
	// EXERCISE: Simplify It
	//
	//  Can you simplify the if statement inside the code below?
	//
	//  When:
	//    isSphere == true and
	//    radius is equal or greater than 200
	//
	//    It will print "It's a big sphere."
	//
	//    Otherwise, it will print "I don't know."
	//
	// EXPECTED OUTPUT
	//  It's a big sphere.
	// ---------------------------------------------------------

	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	var big bool

	if radius >= 200 {
		big = true
	}

	fmt.Println("SimplifyIt():")
	if !big {
		fmt.Println("I don't know.")
	} else if isSphere {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}

func (i *ifStatementEx) ArgCount() {
	// ---------------------------------------------------------
	// EXERCISE: Arg Count
	//
	//  1. Get arguments from command-line.
	//  2. Print the expected outputs below depending on the number
	//     of arguments.
	//
	// EXPECTED OUTPUT
	//  go run main.go
	//    Give me args
	//
	//  go run main.go hello
	//    There is one: "hello"
	//
	//  go run main.go hi there
	//    There are two: "hi there"
	//
	//  go run main.go I wanna be a gopher
	//    There are 5 arguments
	// ---------------------------------------------------------

	agrsLength := len(os.Args) - 1
	isNotHaveArgs := agrsLength == 0
	isHaveOneArg := agrsLength == 1
	isHaveTwoArgs := agrsLength == 2

	const (
		fxName   = "ArgCount():"
		noArgs   = "Give me args"
		oneArgs  = "There is one: \"%s\""
		twoArgs  = "There are two: \"%s %s\""
		manyArgs = "There are %d arguments"
	)

	fmt.Println(fxName)
	if isNotHaveArgs {
		fmt.Println(noArgs)
	} else if isHaveOneArg {
		fmt.Printf(oneArgs, os.Args[1])
	} else if isHaveTwoArgs {
		fmt.Printf(twoArgs, os.Args[1], os.Args[2])
	} else {
		fmt.Printf(manyArgs, agrsLength)
	}

}

func (i *ifStatementEx) VowelOrConsonant() {

	// ---------------------------------------------------------
	// EXERCISE: Vowel or Consonant
	//
	//	Detect whether a letter is vowel or consonant.
	//
	// NOTE
	//
	//	y or w is called a semi-vowel.
	//	Check out: https://www.merriam-webster.com/words-at-play/why-y-is-sometimes-a-vowel-usage
	//	Check out: https://www.dictionary.com/e/w-vowel/
	//
	// HINT
	//
	//   - You can find the length of an argument using the len function.
	//
	//   - len(os.Args[1]) will give you the length of the 1st argument.
	//
	// BONUS
	//
	//	Use strings.IndexAny function to detect the vowels.
	//	Search on Google for: golang pkg strings IndexAny
	//
	// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
	//
	// EXPECTED OUTPUT
	//
	//	go run main.go
	//	  Give me a letter
	//
	//	go run main.go hey
	//	  Give me a letter
	//
	//	go run main.go a
	//	  "a" is a vowel.
	//
	//	go run main.go y
	//	  "y" is sometimes a vowel, sometimes not.
	//
	//	go run main.go w
	//	  "w" is sometimes a vowel, sometimes not.
	//
	//	go run main.go x
	//	  "x" is a consonant.
	//
	// ---------------------------------------------------------
	argsCount := len(os.Args) - 1
	isHaveArgs := argsCount == 1
	var letter string
	isLetter := false

	const (
		fxName         = "\nVowelOrConsonant():"
		noLetterTxt    = "Give me a letter"
		isVowelTxt     = "\"%s\" is a vowel.\n"
		isSemiVowelTxt = "\"%s\" is sometimes a vowel, sometimes not\n"
		isConsonantTxt = "\"%s\" is consonant.\n"
		isNormalTxt    = "\"%s\" is normal letter.\n"
	)

	fmt.Println(fxName)
	if isHaveArgs {
		charCount := utf8.RuneCountInString(os.Args[1])
		isLetter = charCount == 1
		if isLetter {
			letter = os.Args[1]
		}
	}

	if !isLetter {
		fmt.Println(noLetterTxt)
	}

	isConsonant := strings.Contains(letter, "x")
	isSemiVowel := strings.ContainsAny(letter, "yw")
	isVowel := strings.ContainsAny(letter, "aeiou")

	if isVowel {
		fmt.Printf(isVowelTxt, letter)
	} else if isSemiVowel {
		fmt.Printf(isSemiVowelTxt, letter)
	} else if isConsonant {
		fmt.Printf(isConsonantTxt, letter)
	} else {
		fmt.Printf(isNormalTxt, letter)
	}
}

func (i *ifStatementEx) Challenge1() {
	// ---------------------------------------------------------
	// CHALLENGE #1
	//  Create a user/password protected program.
	//
	// EXAMPLE USER
	//  username: jack
	//  password: 1888
	//
	// EXPECTED OUTPUT
	//  go run main.go
	//    Usage: [username] [password]
	//
	//  go run main.go albert
	//    Usage: [username] [password]
	//
	//  go run main.go hacker 42
	//    Access denied for "hacker".
	//
	//  go run main.go jack 6475
	//    Invalid password for "jack".
	//
	//  go run main.go jack 1888
	//    Access granted to "jack".
	// ---------------------------------------------------------
	const (
		user     = "eiei"
		password = "5555"
	)

	var (
		userArgs     string
		passwordArgs string
	)

	const (
		usage    = "Usage: [username] [password]"
		errUser  = "Access denied for %q.\n"
		errPwd   = "Invalid password for %q.\n"
		accessOK = "Access granted to %q.\n"
	)

	argsCount := len(os.Args) - 1
	inValidUser, inValidPassword := argsCount < 1, argsCount < 2
	inValidArgs := inValidUser || inValidPassword
	if !inValidArgs {
		userArgs = os.Args[1]
		passwordArgs = os.Args[2]
	}

	userMatch := userArgs == user
	passwordMatch := passwordArgs == password

	fmt.Println("Challenge1():")
	if inValidArgs {
		fmt.Println(usage)
		return
	}

	if !userMatch {
		fmt.Printf(errUser, userArgs)
	}

	if !passwordMatch {
		fmt.Printf(errPwd, userArgs)
	}

	if userMatch && passwordMatch {
		fmt.Printf(accessOK, userArgs)
	}

}

func (i *ifStatementEx) Challenge2() {
	// ---------------------------------------------------------
	// CHALLENGE #2
	//  Add one more user to the PassMe program below.
	//
	// EXAMPLE USERS
	//  username: jack
	//  password: 1888
	//
	//  username: inanc
	//  password: 1879
	//
	// EXPECTED OUTPUT
	//  go run main.go
	//    Usage: [username] [password]
	//
	//  go run main.go hacker 42
	//    Access denied for "hacker".
	//
	//  go run main.go jack 1888
	//    Access granted to "jack".
	//
	//  go run main.go inanc 1879
	//    Access granted to "inanc".
	//
	//  go run main.go jack 1879
	//    Invalid password for "jack".
	//
	//  go run main.go inanc 1888
	//    Invalid password for "inanc".
	// ---------------------------------------------------------

	const (
		usage    = "Usage: [username] [password]"
		errUser  = "Access denied for %q.\n"
		errPwd   = "Invalid password for %q.\n"
		accessOK = "Access granted to %q.\n"
		fxname   = "Challenge2():"
	)

	const (
		user1 = "jack"
		pass1 = "1888"
		user2 = "inanc"
		pass2 = "1879"
	)

	args := os.Args

	fmt.Println(fxname)
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	isUser1, isUser2 := u == user1, u == user2
	isNotHaveUser := !isUser1 && !isUser2
	isUser1Valid := isUser1 && p == pass1
	isUser2Valid := isUser2 && p == pass2

	if isNotHaveUser {
		fmt.Printf(errUser, u)
	} else if isUser1Valid || isUser2Valid {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}

}
