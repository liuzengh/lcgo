package lcgo

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"test two lists are nil",
			args{[]int{}, []int{}},
			nil,
		},
		{
			"test first list is nil",
			args{[]int{}, []int{1, 2, 3}},
			[]int{1, 2, 3},
		},
		{
			"test second list is nil",
			args{[]int{1, 2, 3}, []int{}},
			[]int{1, 2, 3},
		},
		{
			"test2",
			args{[]int{2, 4}, []int{5, 6}},
			[]int{7, 0, 1},
		},
		{
			"test3",
			args{[]int{0}, []int{0}},
			[]int{0},
		},
		{
			"test4",
			args{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
			[]int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(IntSlice2List(tt.args.l1), IntSlice2List(tt.args.l2)); !reflect.DeepEqual(List2IntSlice(got), tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// IntSlice2List convert []int to List
func IntSlice2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

// List2IntSlice convert List to []int
func List2IntSlice(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	var res []int
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
