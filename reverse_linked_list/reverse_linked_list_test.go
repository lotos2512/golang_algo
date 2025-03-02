package reverse_linked_list

import (
	"reflect"
	"testing"
)

// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	a5 := &ListNode{Val: 5, Next: nil}
	a4 := &ListNode{Val: 4, Next: a5}
	a3 := &ListNode{Val: 3, Next: a4}
	a2 := &ListNode{Val: 2, Next: a3}
	a1 := &ListNode{Val: 1, Next: a2}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: testing.CoverMode(), args: args{head: a1}, want: a5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
