package hash

import "testing"

func Test_singleNumber1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		wantMinV int
	}{
		{name: testing.CoverMode(), args: struct{ nums []int }{nums: []int{4, 1, 2, 1, 2}}, wantMinV: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMinV := singleNumber1(tt.args.nums); gotMinV != tt.wantMinV {
				t.Errorf("singleNumber1() = %v, want %v", gotMinV, tt.wantMinV)
			}
		})
	}
}

func Test_singleNumber2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		wantMinV int
	}{
		{name: testing.CoverMode(), args: struct{ nums []int }{nums: []int{4, 1, 2, 1, 2}}, wantMinV: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMinV := singleNumber2(tt.args.nums); gotMinV != tt.wantMinV {
				t.Errorf("singleNumber2() = %v, want %v", gotMinV, tt.wantMinV)
			}
		})
	}
}
