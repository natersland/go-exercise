package main

import (
	variable "github.com/natersland/go-exercise/pkg/01-variable"
	typeconversion "github.com/natersland/go-exercise/pkg/02-type-conversion"
	cmdargs "github.com/natersland/go-exercise/pkg/03-cmd-args"
)

func main() {
	variableExercise := variable.NewVariableExercise()

	variableExercise.MakeItBlue()
	variableExercise.VarToVar()
	variableExercise.AssignWithExpression()
	variableExercise.FindTheRetangle()
	variableExercise.MultiAssign()
	variableExercise.MultiAssign2()
	variableExercise.MultiShortTime()
	variableExercise.Swapper()
	variableExercise.Swapper2()
	variableExercise.DiscardTheFile()

	typeconversionExercise := typeconversion.NewTypeConversion()

	typeconversionExercise.ConvertAndFix1()
	typeconversionExercise.ConvertAndFix2()
	typeconversionExercise.ConvertAndFix3()
	typeconversionExercise.ConvertAndFix4()
	typeconversionExercise.ConvertAndFix5()

	cmdArgsExercise := cmdargs.NewCommandLineArgs()

	cmdArgsExercise.CountTheArgs()
	cmdArgsExercise.PrintThePath()
	cmdArgsExercise.PrintYourName()
	cmdArgsExercise.GreetMorePeople()
	cmdArgsExercise.GreetFivePeople()

}
