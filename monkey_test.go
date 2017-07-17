package main

import "testing"

func TestBuildMonkey(t *testing.T) {
	monkey, _ := BuildMonkey("Foo", 5, 0)
	if monkey.Name != "Foo" {
		t.Error(monkey.Name)
	}
	if monkey.age != 5 {
		t.Error(monkey.age)
	}
	if monkey.Gender != 0 {
		t.Error(monkey.Gender)
	}
}
