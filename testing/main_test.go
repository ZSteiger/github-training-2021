package main

import (
	"testing"
)

func Test_toTen(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "toTen Successful",
			args: args{1},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toTen(tt.args.i); got != tt.want {
				t.Errorf("toTen() = %v, want %v", got, tt.want)
			}
		})
	}
}
