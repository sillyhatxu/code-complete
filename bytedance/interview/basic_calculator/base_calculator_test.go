package basic_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"1+2*3", 7},
		{"11+22", 33},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
		{"1+2+3", 6},
		{"2147483647", 2147483647},
		{"1+2*3", 7},
		{"(1+(4+5+2)-3)+(6+8)", 23},
		{"2*(5+5*2)/3+(6/2+8)", 21},
		{"3+2*2", 7},
		{"3/2", 1},
		{"3+5/2", 5},
		{"(2+6*3+5-(3*14/7+2)*5)+3", -12},
		{"1*2-3/4+5*6-7*8+9/10", -24},
		{"10+20*3", 70},
		{"1-2*3+4/2", -3},
		{"(2+3)*4", 20},
		{"(5+5*2)/2", 7},
		{"(3+4*5)/2", 11},
		{"2*(3+4)", 14},
		{"(3+5/2)*2", 10},
		{"((3+5)/2)*2", 8},
		{"((2+3)*(5-2))/2", 7},
		{"(6+8)/2*3-4", 17},
		{"(2*3+4)/2", 5},
		{"2*3+4/2*5-6", 10},
	}
	for _, tc := range testCases {
		assert.EqualValues(t, tc.expected, calculate(tc.input), "Failed for input: %s", tc.input)
	}
}
