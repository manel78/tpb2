package b2

func Ft_missing(num []int) int {
	n := len(num)

	expectedSum := n * (n + 1) / 2
	actualSum := 0
	for _, val := range num {
		actualSum += val
	}
	return expectedSum - actualSum
}
