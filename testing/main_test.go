package main

import (
	"testing"
)

// function to run unit tests
func Test_toTen(t *testing.T) {

	// arguments passed to function being tested, aka "starterInt"
	type args struct {
		i int
	}
	// the group of tests to be ran
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
	// for each of the tests above, run them and compare the inputs
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toTen(tt.args.i); got != tt.want {
				t.Errorf("toTen() = %v, want %v", got, tt.want)
			}
		})
	}
}
