package kata

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func ProperFractions(n int) int {
	fmt.Println(n)

	if n <= 1 {
		return 0
	}

	var (
		wg            sync.WaitGroup
		all           = atomic.Int64{}
		blocks        = runtime.NumCPU()
		partsFraction = float64(n) / float64(blocks)
		parts         = int(math.Ceil(partsFraction))
		computer      = func(a, b int) {
			defer wg.Done()
			if a < 0 {
				a = 1
			}
			if b > n {
				b = n
			}

			var count int64
			for i := a; i <= b; i++ {
				if gcd(i, n) == 1 {
					count++
				}
			}

			all.Add(count)
		}
	)

	if parts == 0 {
		parts = n
	}

	start := 1
	end := parts

	for i := 0; i < blocks; i++ {
		wg.Add(1)
		go computer(start, end)
		if end >= n {
			break
		}
		start = end + 1
		end += parts
	}

	wg.Wait()

	return int(all.Load())
}
