package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 1, 2},
			},
			want: res{
				len:  2,
				nums: []int{1, 2},
			},
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: res{
				len:  5,
				nums: []int{0, 1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want.len {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.nums[:tt.want.len], tt.want.nums) {
				t.Errorf("want %v, got %v", tt.want.nums, tt.args.nums[:tt.want.len])
			}
		})
	}
}
