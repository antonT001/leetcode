package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Case 1",
			args: args{
				list1: prepereTestData([]int{1, 2, 4}),
				list2: prepereTestData([]int{1, 3, 4}),
			},
			want: prepereTestData([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "Case 2",
			args: args{
				list1: prepereTestData([]int{}),
				list2: prepereTestData([]int{}),
			},
			want: prepereTestData([]int{}),
		},
		{
			name: "Case 3",
			args: args{
				list1: prepereTestData([]int{}),
				list2: prepereTestData([]int{0}),
			},
			want: prepereTestData([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
