package round

import "testing"

func TestToEven(t *testing.T) {
	tests := []struct {
		in       float64
		decimals int
		expected float64
	}{
		{
			in:       3.5,
			decimals: 0,
			expected: 4,
		},
		{
			in:       2.8,
			decimals: 0,
			expected: 3,
		},
		{
			in:       2.5,
			decimals: 0,
			expected: 2,
		},
		{
			in:       2.1,
			decimals: 0,
			expected: 2,
		},
		{
			in:       -2.1,
			decimals: 0,
			expected: -2,
		},
		{
			in:       -2.5,
			decimals: 0,
			expected: -2,
		},
		{
			in:       -2.8,
			decimals: 0,
			expected: -3,
		},
		{
			in:       -3.5,
			decimals: 0,
			expected: -4,
		},
		{
			in:       1.5,
			decimals: 0,
			expected: 2,
		},
		{
			in:       2.5,
			decimals: 0,
			expected: 2,
		},
		{
			in:       3.5,
			decimals: 0,
			expected: 4,
		},
		{
			in:       -3.5,
			decimals: 0,
			expected: -4,
		},
		{
			in:       -3.12,
			decimals: 0,
			expected: -3,
		},
		{
			in:       1.2345,
			decimals: 1,
			expected: 1.2,
		},
		{
			in:       1.2345,
			decimals: 2,
			expected: 1.23,
		},
		{
			in:       1.2345,
			decimals: 3,
			expected: 1.234,
		},
		{
			in:       7.5,
			decimals: 0,
			expected: 8,
		},
		{
			in:       8.5,
			decimals: 0,
			expected: 8,
		},
		{
			in:       -1.2345,
			decimals: 3,
			expected: -1.234,
		},
		{
			in:       -1.5,
			decimals: 0,
			expected: -2,
		},
		// From https://en.wikipedia.org/wiki/IEEE_floating_point#Rounding_rules
		{
			in:       11.5,
			decimals: 0,
			expected: 12,
		},
		{
			in:       12.5,
			decimals: 0,
			expected: 12,
		},
		{
			in:       -11.5,
			decimals: 0,
			expected: -12,
		},
		{
			in:       -12.5,
			decimals: 0,
			expected: -12,
		},
		{
			in:       1.55,
			decimals: 1,
			expected: 1.6,
		},
		{
			in:       2.55,
			decimals: 1,
			expected: 2.6,
		},
		{
			in:       2.5555555,
			decimals: 3,
			expected: 2.556,
		},
		{
			in:       -2.5555555,
			decimals: 2,
			expected: -2.56,
		},
	}

	for _, test := range tests {
		actual := ToEven(test.in, test.decimals)

		if actual != test.expected {
			t.Errorf("Rounding %v to %d decimal places. Expected %v, but got %v.", test.in, test.decimals, test.expected, actual)
		}
	}
}
