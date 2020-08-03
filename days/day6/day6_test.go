package day6

import "testing"

func TestGetMaxField(t *testing.T) {
	in := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

	max := findMaxField(in)

	if max != 17 {
		t.Fatalf("Expected max area to be 17 but got %v", max)
	}

}

func TestFindCloseField(t *testing.T) {
	dist := 32
	in := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

	close := findCloseField(in, dist)

	if close != 16 {
		t.Fatalf("Expected max area to be 16 but got %v", close)

	}

}
