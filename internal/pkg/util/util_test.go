package util

import "testing"

func TestReverseInt(t *testing.T) {
	number, err := ReverseInt(123)
	if err != nil {
		t.Errorf("Out of range")
	}
	if number != 321 {
		t.Errorf("Not Reverse number")
	}
	NegNumber, err1 := ReverseInt(-649)
	if err1 != nil {
		t.Errorf("Out of range")
	}
	if NegNumber != -946 {
		t.Errorf("Not Reverse number")
	}
	ZeroNumber, err2 := ReverseInt(0)
	if err2 != nil {
		t.Errorf("Out of range")
	}
	if ZeroNumber != 0 {
		t.Errorf("Not Reverse number")
	}
	_, err3 := ReverseInt(2147483648)
	if err3 != nil {
		t.Errorf("Out of range")
	}
	_, err4 := ReverseInt(2.1)
	if err4 == nil {
		t.Errorf("Not int")
	}
}
