package main

import "testing"
import "reflect"

func TestParseText(t *testing.T) {
	input := "енот рад рад"
	expected := map[string]int{
		"енот": 1,
		"рад":  2,
	}
	actual := ParseText(input)
	eq := reflect.DeepEqual(actual, expected)
	if !eq {
		t.Errorf("input=%v, actual=%v, expected=%v", input, actual, expected)
	}
}

func TestParseTextWithComma(t *testing.T) {
	input := "енот рад, рад"
	expected := map[string]int{
		"енот": 1,
		"рад":  2,
	}
	actual := ParseText(input)
	eq := reflect.DeepEqual(actual, expected)
	if !eq {
		t.Errorf("input=%v, actual=%v, expected=%v", input, actual, expected)
	}
}

func TestParseTextWithСapital(t *testing.T) {
	input := "енот рад, Рад"
	expected := map[string]int{
		"енот": 1,
		"рад":  2,
	}
	actual := ParseText(input)
	eq := reflect.DeepEqual(actual, expected)
	if !eq {
		t.Errorf("input=%v, actual=%v, expected=%v", input, actual, expected)
	}
}

func TestGetTopWords(t *testing.T) {
	input := map[string]int{
		"енот": 1,
		"рад":  2,
	}
	expected := [10]string{"рад", "енот"}

	actual := GetTopWords(input)
	eq := reflect.DeepEqual(actual, expected)
	if !eq {
		t.Errorf("input=%v, actual=%v, expected=%v", input, actual, expected)
	}
}
