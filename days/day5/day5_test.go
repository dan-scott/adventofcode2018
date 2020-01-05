package day5

import (
	"testing"
)

func TestCompress(t *testing.T) {
	input := "dabAcCaCBAcCcaDA"

	expectd := "dabCBAcaDA"

	actual := compress(input)

	if actual != expectd {
		t.Fatalf("Expected to get %s but got %s", expectd, actual)
	}
}

func TestOptimize(t *testing.T) {
	input := "dabAcCaCBAcCcaDA"
	expected := 4

	actual := optimize(input)

	if actual != expected {
		t.Fatalf("Expected to get %d but got %d", expected, actual)
	}
}
