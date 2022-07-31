package lcgo

import "testing"

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"pivot index is in middle of nums",
			args{
				[]int{1, 7, 3, 6, 5, 6},
			},
			3,
		},
		{
			"pivot index doesn't exist",
			args{
				[]int{1, 2, 3},
			},
			-1,
		},
		{
			"pivot index is 0",
			args{
				[]int{2, 1, -1},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
