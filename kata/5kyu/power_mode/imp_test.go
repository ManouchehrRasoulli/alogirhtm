package kata

import "testing"

func TestModPow(t *testing.T) {
	/*
	  Expect(ModPow(2, 3, 5)).To(Equal(uint64(3)))
	       Expect(ModPow(4, 12, 3)).To(Equal(uint64(1)))
	       Expect(ModPow(11, 10, 300)).To(Equal(uint64(1)))
	       Expect(ModPow(11, 100000, 49)).To(Equal(uint64(32)))
	       Expect(ModPow(5, 100000000, 19)).To(Equal(uint64(5)))
	*/

	t.Log(ModPow(2, 3, 5))          // 3
	t.Log(ModPow(4, 12, 3))         // 1
	t.Log(ModPow(11, 10, 300))      // 1
	t.Log(ModPow(11, 100000, 49))   // 32
	t.Log(ModPow(5, 100000000, 19)) // 5

	t.Log(ModPow(29, 1013293125, 492))    // 161
	t.Log(ModPow(31, 902938423012, 1023)) // 961
}
