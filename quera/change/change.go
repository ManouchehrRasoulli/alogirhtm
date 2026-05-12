package change

func change(coins []int, amount int) []int {
	coinCount := make([]int, len(coins))
	for i := 0; i < len(coins); i++ {
		for amount >= coins[i] {
			coinCount[i]++
			amount -= coins[i]
		}
	}
	return coinCount
}
