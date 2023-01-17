package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonnaci(t *testing.T) {

	actualAnswer := Fibonacci()
	call := actualAnswer(5)
	actualAnswer(6)
	expectedAnswer := []int{0, 1, 1, 2, 3}
	if !reflect.DeepEqual(call, expectedAnswer) {
		t.Errorf("Expected Answer %q is not same as actual Answer %q .", expectedAnswer, call)
	}

}
