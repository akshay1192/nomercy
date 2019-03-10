package helperFunc

import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {
	actual := Shift("xyz", 2)
	expected := "zab"

	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}

func TestShift2(t *testing.T) {
	actual := Shift("xyz", 26)
	expected := "xyz"

	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}
