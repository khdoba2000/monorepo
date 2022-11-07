package hello

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	expected := "Hello, Earth!"
	actual := Greet("Earth")
	if expected != actual {
		fmt.Println("expected", expected)
		fmt.Println("actual", string(actual))
		t.Fail()
	}
}
