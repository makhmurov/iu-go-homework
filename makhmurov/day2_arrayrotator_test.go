package makhmurov

import (
	"reflect"
	"testing"
)

func isEqual(v1, v2 []int) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i, el := range v1 {
		if el != v2[i] {
			return false
		}
	}
	return true
}

func TestArrayRotate(t *testing.T) {
	arr := []int{3, 8, 9, 7, 6}
	count := 2
	want := []int{7, 6, 3, 8, 9}
	got := ArrayRotate(arr, count)
	if !isEqual(got, want) {
		t.Errorf("ArrayRotate() = %v, want %v", got, want)
	}
}

func genSlice(length, shift int) []int {
	shift = (length + shift%length) % length
	arr := make([]int, length)
	for inc := range arr {
		index := (inc + shift) % length
		arr[index] = inc
	}
	return arr
}

func TestArrayRotate2(t *testing.T) {
	type args struct {
		arr   []int
		count int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// Test cases.
		{name: "sample1", args: args{[]int{3, 8, 9, 7, 6}, 2}, want: []int{7, 6, 3, 8, 9}},
		{name: "sample2", args: args{[]int{1, 2, 3, 4, 5}, 0}, want: []int{1, 2, 3, 4, 5}},
		{name: "sample3", args: args{[]int{1, 2, 3, 4, 5}, 6}, want: []int{5, 1, 2, 3, 4}},
		{name: "sample4", args: args{genSlice(5, 0), 5}, want: genSlice(5, 5)},
		//{name: "sample5", args: args{[]int{1, 2, 3, 4, 5}, -1}, want: []int{1, 2, 3, 4, 5}},
		{name: "sample6", args: args{genSlice(5, 0), -1}, want: genSlice(5, -1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayRotate2(tt.args.arr, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayRotate2(%v) = %v, want %v", tt.args.arr, got, tt.want)
			}
		})
	}
}
