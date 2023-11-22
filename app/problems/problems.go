package problems

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/myProjects/Quiz-Game/app/errors"
	"github.com/myProjects/Quiz-Game/app/flags"
)

type IProblem interface {
	GetQuestion() string
	GetSolution() string
}

func GetProblems() []IProblem {
	questions, err := readFile()
	if err != nil {
		errors.ExitWithMessage("Failed to parse provided csv file")
	}

	return parseProblems(questions)
}

func readFile() ([][]string, error) {
	file, err := os.Open("app/problems/" + *flags.Get().CsvFileName)
	if err != nil {
		errors.ExitWithMessage("Failed to open csv file")
	}

	r := csv.NewReader(file)

	return r.ReadAll()
}

func parseProblems(questions [][]string) []IProblem {
	var problems []IProblem
	for _, question := range questions {
		problem := &Problem{
			Question: strings.TrimSpace(question[0]),
			Solution: strings.TrimSpace(question[1]),
		}

		problems = append(problems, problem)
	}

	return problems
}
