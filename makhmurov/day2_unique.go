package makhmurov

// # Unique
// Дан массив целых чисел, необходимо посчитать, количество уникальных значений

/*
## Пример
arr := [3, 8, 9, 7, 3, 5, 8, 9, 6]
res := GetUniqCount(arr)
//res = 3
*/

func GetUniqCount(arr []int) int {
	var unique = 0
	var count = make(map[int]int)
	for _, value := range arr {
		count[value]++
	}
	for _, used := range count {
		if used == 1 {
			unique++
		}
	}
	return unique
}
