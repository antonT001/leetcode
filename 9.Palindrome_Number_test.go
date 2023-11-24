package leetcode

import "testing"

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				x: 1121,
			},
			want: false,
		},
		{
			name: "Case 3",
			args: args{
				x: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
