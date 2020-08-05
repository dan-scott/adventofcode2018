package day9

import (
	"fmt"
	"testing"
)

func TestWinningScore(t *testing.T) {
	cases := []struct {
		players int
		worth   int
		score   int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("Winning score for %v players, last marble %v", tc.players, tc.worth), func(t *testing.T) {
			score := getWinningScore(tc.players, tc.worth)
			if score != tc.score {
				t.Fatalf("Expected score to be %v but was %v", tc.score, score)
			}
		})
	}
}
