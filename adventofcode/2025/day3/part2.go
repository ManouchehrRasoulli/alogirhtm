package day3

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func MaxJoltagePart2(input string) string {
	var (
		maxValueJoltage = func(bank string) *big.Int {
			replaceMinFromLeft := func(v int, x string) string {
				currentV, _ := strconv.Atoi(x)
				for i := len(x) - 1; i >= 0; i-- {
					if int(x[i]-'0') > v {
						continue
					}

					f := fmt.Sprintf("%d%s%s", v, x[:i], x[i+1:])
					fV, _ := strconv.Atoi(f)
					if fV > currentV {
						currentV = fV
					}
				}

				return fmt.Sprintf("%d", currentV)
			}
			current := bank[len(bank)-12:]
			currentValue := new(big.Int)
			replacedValue := new(big.Int)

			for i := len(bank) - 13; i >= 0; i-- {
				replaced := replaceMinFromLeft(int(bank[i]-'0'), current)
				currentValue, _ = currentValue.SetString(current, 10)
				replacedValue, _ = replacedValue.SetString(replaced, 10)

				if replacedValue.Cmp(currentValue) > 0 {
					current = replaced
				}
			}

			currentValue, _ = currentValue.SetString(current, 10)

			return currentValue
		}

		lines = strings.Split(input, "\n")
		sum   = big.NewInt(0)
	)

	for _, line := range lines {
		sum.Add(sum, maxValueJoltage(line))
	}

	return sum.String()
}
