package main

import (
	"fmt"
	homework "iu-go-homework/makhmurov"
)

func main() {
	fmt.Println("Missing test")
	task1_res := homework.SolutionSquareGenerator(-10, 10)
	fmt.Println("# SolutionSquareGenerator:", task1_res)
	task1_res = nil

	task2_res := homework.SolutionBinaryGap(68421398)
	fmt.Println("# SolutionBinaryGap:", task2_res)

	task4_res := homework.GetUniqCount([]int{3, 8, 9, 7, 3, 5, 8, 9, 6})
	fmt.Println("# GetUniqCount:", task4_res)
}
