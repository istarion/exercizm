package differenceofsquares

func SquareOfSum(n int) int {
	sum := 0
	for i := range n + 1 {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := range n + 1 {
		if i == 0 {
			continue
		}
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
