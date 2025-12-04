package arrays

func Sum(numbers []int) int {
	var out int
	for _, number := range numbers {
		out += number
	}
	return out
}

func SumAll(arrs ...[]int) []int {
	out := []int{}
	for _, arr := range arrs {
		out = append(out, Sum(arr))
	}
	return out
}

func GetTailSums(arraysToSum ...[]int) []int {
	out := []int{}
	for _, numbers := range arraysToSum {
		if len(numbers) > 0 {
			out = append(out, Sum(numbers[1:]))
		} else {
			out = append(out, 0)
		}
	}
	return out
}
