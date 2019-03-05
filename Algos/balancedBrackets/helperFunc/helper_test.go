package helperFunc

import (
	"testing"
)

func TestCheckBalancedBrackets(t *testing.T) {
	actual := CheckBalancedBrackets("{}")
	expected := true

	if actual != expected {
		t.Fail()
	}
}

func TestCheckBalancedBrackets2(t *testing.T) {
	actual := CheckBalancedBrackets("{]")
	expected := false

	if actual != expected {
		t.Fail()
	}
}
