package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "Case 2",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "Case 3",
			args: args{
				x: 120,
			},
			want: 21,
		},
		{
			name: "Case 4",
			args: args{
				x: 1534236469,
			},
			want: 0, 
		},

		{
			name: "Case 5",
			args: args{
				x: 1563847412,
			},
			want: 0, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
