package kata

import "math/big"

func FindMult10SF(n int) *big.Int {
	res := big.NewInt(1)
	cnt := 0
	for i := 1; cnt < n; i++ {
		bi := big.NewInt(int64(i))
		twoExp, sixExp := new(big.Int), new(big.Int)
		twoExp.Exp(big.NewInt(2), bi, nil)
		sixExp.Exp(big.NewInt(6), bi, nil)

		res = big.NewInt(3)
		res.Mul(res, twoExp).Add(res, sixExp).Div(res, big.NewInt(4))

		div := new(big.Int)
		div.Mod(res, big.NewInt(10))
		if div.Cmp(big.NewInt(0)) == 0 {
			cnt++
		}
	}
	return res
}
