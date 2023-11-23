package merge_sort

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int
	}{
		{
			name: "test", args: args{
				input: []int{0, 3, 1, 11, 0, 10},
			},
			wantOutput: []int{0, 0, 1, 3, 10, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput := mergeSort(tt.args.input)
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("mergeSort() = %v, want %v", gotOutput, tt.wantOutput)
			}
			gotOutput = insertSort(tt.args.input)
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("mergeSort() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
