package kata

import (
	"math/big"
)

var memory []*big.Int

func init() {
	// f(0)=0, f(1)=1, f(2)=1, f(3)=2, f(4)=4
	memory = []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(4),
	}
}

func fn(n int) *big.Int {
	if n < len(memory) {
		return memory[n]
	}

	for len(memory) <= n {
		// f(n) = f(n-1) + f(n-2) - f(n-3) + f(n-4) - f(n-5)
		a := memory[len(memory)-1]
		b := memory[len(memory)-2]
		c := memory[len(memory)-3]
		d := memory[len(memory)-4]
		e := memory[len(memory)-5]

		// Calculate: a + b - c + d - e
		val := new(big.Int).Add(a, b)
		val.Sub(val, c)
		val.Add(val, d)
		val.Sub(val, e)

		memory = append(memory, val)
	}
	return memory[n]
}

func KthLastDigitPrime(k int) [2]int {
	found := 0
	n := 0

	mod := big.NewInt(1_000_000_000)

	for {
		val := fn(n)
		if val.Cmp(mod) == 1 {
			last9 := new(big.Int).Mod(val, mod)
			if last9.ProbablyPrime(10) {
				found++
				if found == k {
					return [2]int{n + 1, int(last9.Int64())}
				}
			}
		}
		n++
	}
}
