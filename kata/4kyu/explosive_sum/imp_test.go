package kata

import "testing"

func Test_ExpSum(t *testing.T) {
	var input uint64
	input = 4
	output := ExpSum(input)
	t.Logf("ExpSum(%d) = %d", input, output)

	input = 10
	output = ExpSum(input)
	t.Logf("ExpSum(%d) = %d", input, output)

	input = 100
	output = ExpSum(input)
	t.Logf("ExpSum(%d) = %d", input, output) // 190569292

	input = 400
	output = ExpSum(input)
	t.Logf("ExpSum(%d) = %d", input, output) // 6727090051741041926

	input = 1000
	output = ExpSum(input)
	t.Logf("ExpSum(%d) = %d", input, output) // 13207224549170234103
}
