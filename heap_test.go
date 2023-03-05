package heap

import (
	"fmt"
	"testing"
)

func TestIsHeap(t *testing.T) {
	var tests = []struct {
		heap []int
		want bool
	}{
		{[]int{100, 19, 36, 17, 3, 25, 1, 2, 7}, true},
		{[]int{100, 19, 36, 17, 3, 25, 1, 2, 7, 200}, false},
		{[]int{100}, true},
		{[]int{100, 30}, true},
		{[]int{}, true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("isHeap(%v)", tt.heap)
		t.Run(testname, func(t *testing.T) {
			got := isHeap(tt.heap, 0)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	var tests = []struct {
		values []int
		want   []int
	}{
		{nil, []int{}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{2, 1}, []int{2, 1}},
		{[]int{2, 2}, []int{2, 2}},
		{[]int{3, 7, 1, 8, 2, 5, 9, 4, 6}, []int{9, 8, 5, 7, 2, 3, 1, 4, 6}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("New(%v)", tt.values)
		t.Run(testname, func(t *testing.T) {
			if heap := New(tt.values); !isSame(heap, tt.want) {
				t.Errorf("got %v, want %v", heap, tt.want)
			}
		})
	}
}

func isSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// func TestInsert(t *testing.T) {
// 	heap := []int{7, 100, 2, 19, 1, 36, 25, 17, 3}
// 	Heapify(heap)
// 	if !isHeap(heap, 0) {
// 		t.Errorf("isHeap(%v)=false", heap)
// 	}
// }
//
// func TestExtract(t *testing.T) {
// 	heap := []int{7, 100, 2, 19, 1, 36, 25, 17, 3}
// 	initial := []int{7, 100, 2, 19, 1, 36, 25, 17, 3}
// 	Heapify(heap)
// 	Heapify(initial)
// 	val, _ := Extract(heap)
// 	if val != 100 {
// 		t.Errorf("Extract(%v)=%d", initial, val)
// 	}
// 	if !isHeap(heap, 0) {
// 		t.Errorf("isHeap(%v)=false", heap)
// 	}
// }
