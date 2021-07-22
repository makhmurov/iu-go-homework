package makhmurov

// Найдите самую длинную последовательность нулей в двоичном представлении целого числа.

func SolutionBinaryGap(N int) int {
	//fmt.Printf("%#b\n", N)
	var count, maxCount = 0, 0
	for N > 0 {
		bit := N & 0b1

		if bit == 0 {
			count++
		} else {
			// We are count form right to left without leading zeros,
			// therefore last bit before stop will be one.
			if maxCount < count {
				maxCount = count
			}
			count = 0
		}
		N = N >> 1
	}
	return maxCount
}
