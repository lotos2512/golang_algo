package find_anagrams

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: testing.CoverMode(), args: struct {
				s string
				p string
			}{
				s: "cbaebabacd", p: "abc",
			},
			want: []int{0, 6},
		},
		{
			name: testing.CoverMode(), args: struct {
				s string
				p string
			}{
				s: "abab", p: "ab",
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
