package main

import (
	"github.com/Rajan-226/Quiz-Game/app/flags"
	"github.com/Rajan-226/Quiz-Game/app/problems"
)

func main() {
	flags.Initialize()

	problems.GetProblems()
}
