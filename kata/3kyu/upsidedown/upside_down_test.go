package kata

import (
	"fmt"
	"math"
	"testing"
)

func TestIsReversible(t *testing.T) {
	table := []struct {
		input        int64
		index        int64
		isReversible bool
	}{
		{
			input:        400,
			index:        2,
			isReversible: false,
		},
		{
			input:        30,
			index:        1,
			isReversible: false,
		},
		{
			input:        2,
			index:        0,
			isReversible: false,
		},
		{
			input:        184992,
			index:        3,
			isReversible: false,
		},
	}

	for _, tt := range table {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			ok, index := isReversible(tt.input)
			if ok != tt.isReversible {
				t.Errorf("got %v, want %v", ok, tt.isReversible)
			}

			next := math.Pow(10, float64(tt.index))
			t.Log("next index :", next)

			if index != tt.index {
				t.Errorf("got %v, want %v", index, tt.index)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	table := []struct {
		input    int64
		expected int64
	}{
		{
			input:    0,
			expected: 0,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    9981,
			expected: 1866,
		},
		{
			input:    872,
			expected: 0,
		},
		{
			input:    98186,
			expected: 98186,
		},
	}

	for _, test := range table {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			output := reverse(test.input)
			if output != test.expected {
				t.Errorf("output: %v != %v", output, test.expected)
			}
		})
	}
}

func TestUpsideDown(t *testing.T) {
	table := []struct {
		start         string
		end           string
		expectedCount uint64
	}{
		{
			start:         "0",
			end:           "10",
			expectedCount: 3,
		},
		{
			start:         "6",
			end:           "25",
			expectedCount: 2,
		},
		{
			start:         "10",
			end:           "100",
			expectedCount: 4,
		},
		{
			start:         "100",
			end:           "1000",
			expectedCount: 12,
		},
		{
			start:         "100000",
			end:           "123456789",
			expectedCount: 1050,
		},
		{
			start:         "100000",
			end:           "12345678900",
			expectedCount: 5650,
		},
	}

	for _, test := range table {
		t.Run(fmt.Sprintf("start :: %s - to :: %s", test.start, test.end), func(t *testing.T) {
			output := UpsideDown(test.start, test.end)
			if output != test.expectedCount {
				t.Errorf("output: %v != %v", output, test.expectedCount)
			}
		})
	}
}
