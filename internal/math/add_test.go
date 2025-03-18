package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2.2)
	expected := 3.2
	if result != expected {
		t.Errorf("result: %v, expected: %v", result, expected)
	}
}

func TestAddTableDriven(t *testing.T) {
	table := []struct {
		name     string
		a, b     float32
		expected float32
	}{
		{"2.2+2.11", 2.2, 2.11, 4.31},
		{"2.2+2", 2.2, 2, 4.2},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			expected := tc.expected
			if result != expected {
				t.Errorf("result: %v, expected: %v", result, expected)
			}
		})
	}
}
