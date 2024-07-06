package typeconversion

import "fmt"

type TypeConversion interface {
	ConvertAndFix1()
	ConvertAndFix2()
	ConvertAndFix3()
	ConvertAndFix4()
	ConvertAndFix5()
}

type typeconversion struct {
}

func NewTypeConversion() TypeConversion {
	return &typeconversion{}
}

func (t *typeconversion) ConvertAndFix1() {
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix
	//
	//  Fix the code by using the conversion expression.
	//
	// EXPECTED OUTPUT
	//  15.5
	// ---------------------------------------------------------
	a, b := 10, 5.5
	fmt.Println("Converted1:", (float64(a) + b))

}

func (t *typeconversion) ConvertAndFix2() {
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #2
	//
	//  Fix the code by using the conversion expression.
	//
	// EXPECTED OUTPUT
	//  10.5
	// ---------------------------------------------------------
	a, b := 10, 5.5
	a = int(b)
	fmt.Println("Converted2:", (float64(a) + b))

}

func (t *typeconversion) ConvertAndFix3() {
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #3
	//
	//  Fix the code.
	//
	// EXPECTED OUTPUT
	//  5.5
	// ---------------------------------------------------------
	fmt.Println("Converted3:", (float64(5.5)))

}

func (t *typeconversion) ConvertAndFix4() {
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #4
	//
	//  Fix the code.
	//
	// EXPECTED OUTPUT
	//  9.5
	// ---------------------------------------------------------
	age := 2
	fmt.Println("Converted4:", (float64(7.5) + float64(age)))

}

func (t *typeconversion) ConvertAndFix5() {
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #5
	//
	//  Fix the code.
	//
	// HINTS
	//   maximum of int8  can be 127
	//   maximum of int16 can be 32767
	//
	// EXPECTED OUTPUT
	//  1127
	// ---------------------------------------------------------
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println("Converted5:", (int16(max) + int16(min)))

}
