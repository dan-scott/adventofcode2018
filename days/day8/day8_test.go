package day8

import "testing"

func TestSumMetadata(t *testing.T) {
	input := "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

	expected := 138

	actual := sumMetadata(input)

	if actual != expected {
		t.Fatalf("Expected to get %v but got %v", expected, actual)
	}
}
