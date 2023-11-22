package problems

import (
	"encoding/csv"
	"github.com/myProjects/Quiz-Game/app/errors"
	"github.com/myProjects/Quiz-Game/app/flags"
	"os"
)

type IProblem interface {
	GetQuestion() string
	GetSolution() string
}

type Problem struct {
	Question string
	Solution string
}

func GetProblems() []IProblem {
	file, err := os.Open("app/problems/" + *flags.Get().CsvFileName)
	if err != nil {
		errors.ExitWithMessage("Failed to open csv file")
	}

	r := csv.NewReader(file)
	questions, err := r.ReadAll()
	if err != nil {
		errors.ExitWithMessage("Failed to parse provided csv file")
	}

	var problems []IProblem

	for _, question := range questions {
		problem := &Problem{
			Question: question[0],
			Solution: question[1],
		}

		problems = append(problems, problem)
	}

	return problems
}

func (p *Problem) GetQuestion() string {
	return p.Question
}

func (p *Problem) GetSolution() string {
	return p.Solution
}
