package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	type res struct {
		len  int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want res
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: res{
				len:  2,
				nums: []int{2, 2},
			},
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: res{
				len:  5,
				nums: []int{0, 1, 4, 0, 3},
			},
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{1},
				val:  1,
			},
			want: res{
				len:  0,
				nums: []int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if got != tt.want.len {
				t.Errorf("removeElement() = %v, want %v", got, tt.want.len)
			}
			if !reflect.DeepEqual(tt.args.nums[:tt.want.len], tt.want.nums) {
				t.Errorf("want %v, got %v", tt.want.nums, tt.args.nums[:tt.want.len])
			}
		})
	}
}
