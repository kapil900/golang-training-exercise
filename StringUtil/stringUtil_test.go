package stringUtil

import "testing"

func TestCommonString1(t *testing.T) {

	input1 := "ahelloworld"
	input2 := "hello"
	actualString := CommonString(input1, input2)

	expectedString := "hello"
	if actualString != expectedString {
		t.Errorf("Expected String %s is not same as"+"actual string %s .", expectedString, actualString)
	}
}
func TestCommonString2(t *testing.T) {

	input1 := "my"
	input2 := "MY"
	actualString := CommonString(input1, input2)

	expectedString := "my"
	if actualString != expectedString {
		t.Errorf("Expected String %s is not same as"+"actual string %s .", expectedString, actualString)
	}
}
func TestCommonString3(t *testing.T) {

	input1 := ""
	input2 := "he"
	actualString := CommonString(input1, input2)

	expectedString := ""
	if actualString != expectedString {
		t.Errorf("Expected String %s is not same as"+"actual string %s .", expectedString, actualString)
	}
}
func TestCommonString4(t *testing.T) {

	input1 := ""
	input2 := ""
	actualString := CommonString(input1, input2)

	expectedString := "l"
	if actualString != expectedString {
		t.Errorf("Expected String %s is not same as"+"actual string %s .", expectedString, actualString)
	}
}
func TestCommonString5(t *testing.T) {

	input1 := "ahellow"
	input2 := "hello"
	actualString := CommonString(input1, input2)

	expectedString := "ahellow"
	if actualString != expectedString {
		t.Errorf("Expected String %s is not same as"+"actual string %s .", expectedString, actualString)
	}
}
func TestCommonString6(t *testing.T) {

	input1 := "done"
	input2 := "done"
	actualString := CommonString(input1, input2)

	expectedString := "done"
	if actualString != expectedString {
		t.Errorf("Expected String %s is not same as"+"actual string %s .", expectedString, actualString)
	}
}
