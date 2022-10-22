package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 14
	actual := Add(3, 4, 7)

	if expected != actual {
		t.Errorf("err: expected value: %d is not equal with actual value: %d\n", expected, actual)
	}
}

func TestSubtract(t *testing.T) {
	expected := 1
	actual := Subtract(12, 3, 7)

	if expected != actual {
		t.Errorf("err: expected value: %d is not equal with actual value: %d\n", expected, actual)
	}
}
