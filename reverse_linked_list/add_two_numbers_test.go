package reverse_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	a3 := &ListNode{Val: 3, Next: nil}
	a2 := &ListNode{Val: 4, Next: a3}
	a1 := &ListNode{Val: 2, Next: a2}
	// 342

	b3 := &ListNode{Val: 4, Next: nil}
	b2 := &ListNode{Val: 6, Next: b3}
	b1 := &ListNode{Val: 5, Next: b2}
	// 464

	c3 := &ListNode{Val: 8, Next: nil}
	c2 := &ListNode{Val: 0, Next: c3}
	c1 := &ListNode{Val: 7, Next: c2}
	// 342 + 464 = 807

	fmt.Println(a1, b1, c1)

	d7 := &ListNode{Val: 9, Next: nil}
	d6 := &ListNode{Val: 9, Next: d7}
	d5 := &ListNode{Val: 9, Next: d6}
	d4 := &ListNode{Val: 9, Next: d5}
	d3 := &ListNode{Val: 9, Next: d4}
	d2 := &ListNode{Val: 9, Next: d3}
	d1 := &ListNode{Val: 9, Next: d2}

	e4 := &ListNode{Val: 9, Next: nil}
	e3 := &ListNode{Val: 9, Next: e4}
	e2 := &ListNode{Val: 9, Next: e3}
	e1 := &ListNode{Val: 9, Next: e2}

	f8 := &ListNode{Val: 1, Next: nil}
	f7 := &ListNode{Val: 0, Next: f8}
	f6 := &ListNode{Val: 0, Next: f7}
	f5 := &ListNode{Val: 0, Next: f6}
	f4 := &ListNode{Val: 9, Next: f5}
	f3 := &ListNode{Val: 9, Next: f4}
	f2 := &ListNode{Val: 9, Next: f3}
	f1 := &ListNode{Val: 8, Next: f2}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		//{
		//	name: testing.CoverMode(), args: args{
		//		l1: a1,
		//		l2: b1,
		//	},
		//	want: c1,
		//},
		{
			name: testing.CoverMode(), args: args{
				l1: d1,
				l2: e1,
			},
			want: f1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
