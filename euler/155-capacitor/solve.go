package euler

import (
	"strconv" // For converting floats to string for map key, addressing precision
)

// A small epsilon for floating point comparisons to handle precision issues
// when storing/comparing float64 directly in a map.
// For extremely high precision requirements, consider using a fractional representation (e.g., big.Rat)
const epsilon = 1e-9

// uniqueResults stores and quickly checks for unique capacitance values found so far.
// We use string keys derived from float64 to mitigate precision issues when using
// float64 directly as map keys, as two numerically identical floats might have different
// underlying bit representations due to calculation paths.
var uniqueResults = make(map[string]struct{})

// findCombinations recursively explores unique capacitance values.
//
// currentCap: The current capacitance value formed by a combination.
// availableCapacitors: The original list of capacitors.
// usedMask: A bitmask indicating which original capacitors (by their index)
//
//	have been used to form currentCap.
//
// maxCapacitors: The total number of capacitors available. This is used
//
//	to stop recursion if we exceed the count of physical capacitors.
func findCombinations(
	currentCap float64,
	availableCapacitors []float64, // The original N capacitors
	usedMask int,                  // Bitmask for used original capacitors
	maxCapacitors int,             // Total count of original capacitors
) {
	// Convert currentCap to a string with controlled precision for map key
	// This helps in treating numerically close floats as identical for map lookup.
	// You might need to adjust the precision (e.g., 9) based on your needs.
	key := strconv.FormatFloat(currentCap, 'f', 9, 64)

	if _, found := uniqueResults[key]; found {
		return // This value (or one very close to it) has already been found
	}
	uniqueResults[key] = struct{}{} // Add the new unique value

	// Iterate through all available original capacitors to combine with currentCap
	for i, c := range availableCapacitors {
		// If this capacitor 'c' has already been used in the current combination 'currentCap'
		if (usedMask>>i)&1 == 1 {
			continue // Skip it, as it's already part of the current sub-circuit
		}

		// Calculate how many original capacitors are already used in currentCap.
		// This bitcount helps in preventing combinations from exceeding the total number of physical capacitors.
		usedCount := 0
		for j := 0; j < maxCapacitors; j++ {
			if (usedMask>>j)&1 == 1 {
				usedCount++
			}
		}

		// If combining currentCap (which uses `usedCount` caps) with `c` (1 cap)
		// would exceed the total number of available physical capacitors, skip this combination.
		if usedCount+1 > maxCapacitors {
			continue
		}

		// Create a new mask for the recursive calls, marking 'c' as used.
		newUsedMask := usedMask | (1 << i)

		// 1. Parallel Combination
		// This creates a new 'equivalent capacitor' (currentCap || c)
		parallelCap := currentCap + c
		findCombinations(parallelCap, availableCapacitors, newUsedMask, maxCapacitors)

		// 2. Series Combination
		// This creates a new 'equivalent capacitor' (currentCap in series with c)
		// Avoid division by zero if currentCap or c is zero (though for capacitors, they are usually positive)
		if currentCap != 0 && c != 0 {
			seriesCap := (currentCap * c) / (currentCap + c)
			findCombinations(seriesCap, availableCapacitors, newUsedMask, maxCapacitors)
		}
	}
}

// Solve calculates the number of unique capacitance values that can be formed
// from a given number of identical capacitors (e.g., 60uF each).
func Solve(numCapacitors int) int {
	// Reset global state for each call to Solve
	uniqueResults = make(map[string]struct{})

	// Create a slice representing the identical capacitors
	// (e.g., [60.0, 60.0, 60.0] for numCapacitors=3)
	// Even though values are identical, their indices are distinct for the bitmask.
	capacitors := make([]float64, numCapacitors)
	for i := 0; i < numCapacitors; i++ {
		capacitors[i] = 60.0 // All capacitors are 60uF
	}

	// Start recursion with each individual capacitor as a base combination.
	// Each individual capacitor forms a base combination.
	// Its usedMask will only have its own bit set, and it uses 1 capacitor.
	for i, c := range capacitors {
		findCombinations(c, capacitors, (1 << i), numCapacitors)
	}

	// Optional: Print the actual unique values for debugging/verification
	// var resultSlice []float64
	// for valStr := range uniqueResults {
	// 	val, _ := strconv.ParseFloat(valStr, 64)
	// 	resultSlice = append(resultSlice, val)
	// }
	// sort.Float64s(resultSlice)
	// fmt.Println("Unique values:", resultSlice)

	return len(uniqueResults)
}
