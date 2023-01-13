package pkg

type Algorithm struct {
}

func (a *Algorithm) Add(num1, num2 int) int {
	return num1 + num2
}

func NewAlgorithm() *Algorithm {
	return &Algorithm{}
}
