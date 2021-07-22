package makhmurov

import "testing"

func TestGetUniqCount(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Test cases.
		{"Sample test 1", args{[]int{3, 8, 9, 7, 3, 5, 8, 9, 6}}, 3},
		{"All unique", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 9},
		{"All non-unique", args{[]int{1, 2, 3, 4, 1, 2, 3, 4}}, 0},
		{"All the same", args{[]int{4, 4, 4, 4, 4, 4, 4, 4}}, 0},
		{"One unique", args{[]int{1, 2, 3, 4, 5, 1, 2, 3, 4}}, 1},
		{"2U3R", args{[]int{1, 2, 3, 1, 1}}, 2},
		{"3U2R", args{[]int{1, 2, 3, 4, 1}}, 3},
		{"2U4R", args{[]int{1, 2, 2, 4, 4, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUniqCount(tt.args.arr); got != tt.want {
				t.Errorf("GetUniqCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
