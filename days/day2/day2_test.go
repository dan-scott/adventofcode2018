package day2

import "testing"

func TestFindAdjIds(t *testing.T) {
	input := `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`

	expected := "fgij"

	actual := findAdjIds(input)

	if expected != actual {
		t.Fatalf("Expected to get %s but got %s", expected, actual)
	}
}
