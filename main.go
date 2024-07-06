package main

import variable "github.com/natersland/go-exercise/pkg/01-variable"

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

}
