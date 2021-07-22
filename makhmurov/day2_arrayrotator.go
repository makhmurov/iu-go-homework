package makhmurov

/*
# ArrayRotator
Дан массив arr, состоящий из N целых чисел. Вращение массива - это сдвиг
каждого элемента вправо на один индекс, а последний элемент массива
перемещается на первое место. Например, поворот массива arr = [3, 8, 9, 7, 6]
равен [6, 3, 8, 9, 7] (элементы сдвигаются вправо на один индекс,
а 6 перемещается на первое место). Задача сдвинуть массив arr, count раз,
то есть, каждый элемент будет сдвинут вправо count раз

## Пример
arr := [3, 8, 9, 7, 6]
count := 2
res := ArrayRotate(arr, count)
//res = [7, 6, 3, 8, 9]
*/

func ArrayRotate(arr []int, count int) []int {
	high := len(arr)
	if high == 0 {
		return nil
	}
	res := make([]int, high)
	//shift := count % len(arr)
	shift := (high + count%high) % high
	copy(res[shift:], arr[:high-shift])
	copy(res[:shift], arr[high-shift:])
	return res
}

func ArrayRotate2(arr []int, count int) []int {
	high := len(arr)
	if high == 0 {
		return nil
	}
	res := make([]int, high)
	//shift := count % len(arr)
	shift := (high + count%high) % high
	for i := 0; i < high; i++ {
		dsti := (shift + i) % high
		res[dsti] = arr[i]
	}
	return res
}
