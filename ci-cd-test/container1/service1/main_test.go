package main

import "testing"

func TestOneEqualsOne(t *testing.T) {
	if 1 != 1 {
		t.Fatalf("expected 1 to equal 1")
	}
}

func TestTwoEqualsTwo(t *testing.T) {
	if 2 != 2 {
		t.Fatalf("expected 2 to equal 2")
	}
}

func TestThreeEqualsThree(t *testing.T) {
	if 3 != 3 {
		t.Fatalf("expected 3 to equal 3")
	}
}
