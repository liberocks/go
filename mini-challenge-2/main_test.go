package main

import (
	"fmt"
	"testing"
)

func compareMaps(map1, map2 map[string]int) bool {
	// Check if the maps have the same length
	if len(map1) != len(map2) {
		return false
	}

	// Check each key-value pair
	for key, value1 := range map1 {
		value2, exists := map2[key]
		if !exists {
			return false
		}
		if value1 != value2 {
			return false
		}
	}

	return true
}

func TestCalculateWord(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"mini challenge", map[string]int{" ": 1, "a": 1, "c": 1, "e": 2, "g": 1, "h": 1, "i": 2, "l": 2, "m": 1, "n": 2}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input_%s", test.input), func(t *testing.T) {
			result := calculateWord(test.input)

			if !compareMaps(result, test.expected) {
				t.Errorf("For input %s, expected:\n%v\nbut got:\n%v", test.input, test.expected, result)
			}
		})
	}
}
