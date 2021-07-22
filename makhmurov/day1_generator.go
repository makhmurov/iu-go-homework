package makhmurov

// Напишите генератор квадратов натуральных чисел.

func SolutionSquareGenerator(start int, n int) []int {
	if n <= 0 {
		return nil
	}
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = start * start
		start++
	}
	return res
}
