package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacc(t *testing.T) {
	call := Fibonacci()

	errorTestCases := []struct {
		name     string
		input    int
		expected []int
	}{
		{name: "for zero", input: 0, expected: []int{}},
		{name: "for one", input: 1, expected: []int{0}},
		{name: "for -one", input: -1, expected: []int{}},
		{name: "for 9", input: 9, expected: []int{0, 1, 2, 3, 5, 8, 13, 21}},
	}
	for _, value := range errorTestCases {

		actualAnswer := call(value.input)

		expectedAnswer := value.expected

		if !reflect.DeepEqual(actualAnswer, expectedAnswer) {
			t.Errorf("got %v, wanted %v", &actualAnswer, expectedAnswer)
		}
	}

} /*a:= type errorTestcases struct {
	description string
	input       call
	expectedAnswer []int
}
for _,test:=range[] errorTestcases{
	description:"djj",
	input : 6,
	expectedAnswer : []int{0,1,1,2,3},
}*/

/*func TestFibonnaci(t *testing.T) {

	actualAnswer := Fibonacci()
	call := actualAnswer(5)
	actualAnswer(6)
	expectedAnswer := []int{0, 1, 1, 2, 3}
	if !reflect.DeepEqual(call, expectedAnswer) {
		t.Errorf("Expected Answer %q is not same as actual Answer %q .", expectedAnswer, call)
	}

}*/
