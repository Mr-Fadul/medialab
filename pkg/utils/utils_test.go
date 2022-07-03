package utils

import (
	"testing"
)

func TestGreet(t *testing.T) {
	expected := "Hello, Earth!"
	actual := Greet("Earth")
	if expected != actual {
		t.Fail()
	}
}
