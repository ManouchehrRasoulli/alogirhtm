package kata

import "testing"

func TestOrder(t *testing.T) {
	t.Log(Order("is2 Thi1s T4est 3a"))               // "Thi1s is2 3a T4est"
	t.Log(Order("4of Fo1r pe6ople g3ood th5e the2")) // "Fo1r the2 g3ood 4of th5e pe6ople"
	t.Log(Order(""))                                 // ""
}
