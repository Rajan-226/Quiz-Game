package problems

type Problem struct {
	Question string
	Solution string
}

func (p *Problem) GetQuestion() string {
	return p.Question
}

func (p *Problem) GetSolution() string {
	return p.Solution
}
