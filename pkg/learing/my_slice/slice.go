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

//from index i to index j cut
func cutSlice(i, j int, input []int) []int {
	return append(input[:i], input[j+1:]...)
}

func deleteSlice(i int, input []int) []int {
	return append(input[:i], input[i+1:]...)
	return input[:i+copy(input[i:], input[i+1:])]
}

func keepSlice(keepFunc func(i int) bool, input []int) []int {
	n := 0
	for _, x := range input {
		if keepFunc(x) {
			input[n] = x
			n++
		}
	}
	input = input[:n]

	return input
}
