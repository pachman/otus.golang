package main

import "testing"

func TestUnpack1(t *testing.T) {
	input := "a4bc2d5e"
	expected := "aaaabccddddde"

	actual, _ := Unpack(input)
	if actual != expected {
		t.Errorf("input=%v, actual=%v, expected=%v", input, actual, expected)
	}
}

func TestUnpack2(t *testing.T) {
	input := "abcd"
	expected := "abcd"

	actual, _ := Unpack(input)
	if actual != expected {
		t.Errorf("input=%v, actual=%v, expected=%v", input, actual, expected)
	}
}

func TestUnpack3(t *testing.T) {
	input := "45"

	_, err := Unpack(input)
	if err == nil {
		t.Error("Error not set")
	}
}

func TestUnpack4(t *testing.T) {
	input := `qwe\4\5`
	expected := "qwe45"

	actual, _ := Unpack(input)
	if actual != expected {
		t.Errorf("input=%v, actual=%v, expected=%v", input, actual, expected)
	}
}

func TestUnpack5(t *testing.T) {
	input := `qwe\45`
	expected := "qwe44444"

	actual, _ := Unpack(input)
	if actual != expected {
		t.Errorf("input=%v, actual=%v, expected=%v", input, actual, expected)
	}
}

func TestUnpack6(t *testing.T) {
	input := `qwe\\5`
	expected := `qwe\\\\\`

	actual, _ := Unpack(input)
	if actual != expected {
		t.Errorf("input=%v, actual=%v, expected=%v", input, actual, expected)
	}
}
