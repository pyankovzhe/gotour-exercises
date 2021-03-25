package intstack

import "testing"

func TestPush(t *testing.T) {
	input := []int{1, 2, 3, 4}
	st := IntStack{}

	for _, v := range input {
		st.Push(v)
	}

	if !sliceEqual(st.stack, input) {
		t.Errorf("Expected %v, got %v", input, st.stack)
	}
}

func TestPop(t *testing.T) {
	output := []int{4, 3, 2, 1}
	st := IntStack{
		stack: []int{1, 2, 3, 4},
	}

	for i := range st.stack {
		if v := st.Pop(); output[i] != v {
			t.Fatalf("Expected %v at %v, got %v", output[i], i, v)
		}
	}
}

func sliceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
