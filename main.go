package main

import (
	"github.com/myProjects/Quiz-Game/app/flags"
	"github.com/myProjects/Quiz-Game/app/problems"
)

func main() {
	flags.Initialize()

	problems.GetProblems()
}
