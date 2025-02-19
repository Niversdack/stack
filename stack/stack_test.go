package stack

import (
	"reflect"
	"testing"
)

func TestStack_PopInt(t *testing.T) {
	s := Stack[int]{
		a: []int{1, 2, 3},
	}
	if got, _ := s.Pop(); !reflect.DeepEqual(got, 3) {
		t.Errorf("Pop() = %v, want %v", got, 3)
	}
	if got, _ := s.Pop(); !reflect.DeepEqual(got, 2) {
		t.Errorf("Pop() = %v, want %v", got, 2)
	}
	if got, _ := s.Pop(); !reflect.DeepEqual(got, 1) {
		t.Errorf("Pop() = %v, want %v", got, 1)
	}
	if got, _ := s.Pop(); !reflect.DeepEqual(got, 0) {
		t.Errorf("Pop() = %v, want %v", got, 0)
	}
}
