package utils

import (
	"testing"
)

func Test_Set(t *testing.T) {
	set := NewSet[string]()
	set.Add("one")
	if set.Size() != 1 {
		t.Fail()
	}
	set.Add("one")
	if set.Size() != 1 {
		t.Fail()
	}
	set.Add("two")
	if set.Size() != 2 {
		t.Fail()
	}
	if !set.Exists("one") {
		t.Fail()
	}
}
