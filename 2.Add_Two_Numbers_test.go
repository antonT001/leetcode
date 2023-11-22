package leetcode

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Case 1",
			args: args{
				l1: prepereTestData([]int{2, 4, 3}),
				l2: prepereTestData([]int{5, 6, 4}),
			},
			want: prepereTestData([]int{7, 0, 8}),
		},
		{
			name: "Case 2",
			args: args{
				l1: prepereTestData([]int{0}),
				l2: prepereTestData([]int{0}),
			},
			want: prepereTestData([]int{0}),
		},
		{
			name: "Case 3",
			args: args{
				l1: prepereTestData([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: prepereTestData([]int{9, 9, 9, 9}),
			},
			want: prepereTestData([]int{8, 9, 9, 9, 0, 0, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !deepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func deepEqual(got, want *ListNode) bool {
	if got == nil || want == nil {
		return false
	}
	for got != nil && want != nil {
		if got == nil || want == nil {
			return false
		}
		if got.Val != want.Val {
			return false
		}
		got = got.Next
		want = want.Next

	}
	return true
}

func prepereTestData(in []int) *ListNode {
	if len(in) == 0 {
		return nil
	}
	return &ListNode{
		Val:  in[0],
		Next: prepereTestData(in[1:]),
	}
}
