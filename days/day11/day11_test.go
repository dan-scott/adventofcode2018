package day11

import (
	"fmt"
	"testing"
)

func TestPowerLevel(t *testing.T) {
	cases := []struct{ x, y, serial, level int }{

		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Cell {%v, %v}, serial %v", tc.x, tc.y, tc.serial), func(t *testing.T) {
			power := powerLvl(tc.x, tc.y, tc.serial)
			if power != tc.level {
				t.Fatalf("Expected to get power level %v but got %v", tc.level, power)
			}
		})
	}
}

func TestMaxLevel(t *testing.T) {
	cases := []struct{ x, y, serial int }{
		{33, 45, 18},
		{21, 61, 42},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("Serial %v", tc.serial), func(t *testing.T) {
			loc := findMax(tc.serial)
			if loc.x != tc.x && loc.y != tc.y {
				t.Fatalf("Expected location to be %v, %v but got %v", tc.x, tc.y, loc)
			}
		})
	}
}

func TestMaxDynamic(t *testing.T) {
	cases := []struct{ x, y, size, serial int }{
		{90, 269, 16, 18},
		{232, 251, 12, 42},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("Serial %v", tc.serial), func(t *testing.T) {
			loc, size := findMaxDynamic(tc.serial)
			if loc.x != tc.x && loc.y != tc.y && size != tc.size {
				t.Fatalf("Expected location to be %v, %v with size %v but got %v, %v", tc.x, tc.y, tc.size, loc, size)
			}
		})
	}
}
