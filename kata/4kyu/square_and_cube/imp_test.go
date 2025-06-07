package kata

import "testing"

func TestSqCubRevPrime(t *testing.T) {
	t.Log(SqCubRevPrime(220)) // 42982
	t.Log(SqCubRevPrime(4))   // 328
	t.Log(SqCubRevPrime(110)) // 19631
	t.Log(SqCubRevPrime(5))   // 890
	t.Log(SqCubRevPrime(100)) // 19271
	t.Log(SqCubRevPrime(150)) // 31928
	t.Log(SqCubRevPrime(110)) // 19631
	t.Log(SqCubRevPrime(190)) // 41429
}
