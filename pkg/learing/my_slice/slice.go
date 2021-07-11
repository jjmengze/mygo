package my_slice

func popFront(input []int) (int, []int) {
	return input[0], input[1:]
}

func pushToFront(e int, input []int) []int {
	return append([]int{e}, input...)
}

func pushToEnd(e int, input []int) []int {
	return append(input, e)
}

func PopEnd(input []int) (int, []int) {
	return input[len(input)-1], input[:len(input)-1]
}
