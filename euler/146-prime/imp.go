package _46_prime

import (
	"log"
	"math"
	"sync"
	"sync/atomic"
)

func isPrime(num int) bool {
	if num%2 == 0 || num%3 == 0 {
		return false
	}

	sqrtN := int(math.Sqrt(float64(num)))
	for i := 5; i <= sqrtN; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}

	return true
}

func Count(till int) int {
	var (
		sum       atomic.Int64
		bucket    = 10_000
		wg        sync.WaitGroup
		goroutine = make(chan struct{}, 100)
	)

	for i := 0; i < 100; i++ {
		goroutine <- struct{}{}
	}

	for i := 0; i < till; i += bucket {
		<-goroutine
		wg.Add(1)

		go func(start, end int) {
			defer func() {
				wg.Done()
				goroutine <- struct{}{}
			}()

			localSum := 0
			for i := start; i < end; i += 2 {
				if i%3 == 0 || i%7 == 0 || i%13 == 0 || i%10 != 0 {
					continue
				}

				sq := i * i

				if !isPrime(sq+1) ||
					!isPrime(sq+3) ||
					!isPrime(sq+7) ||
					!isPrime(sq+9) ||
					!isPrime(sq+13) ||
					isPrime(sq+21) ||
					isPrime(sq+25) ||
					!isPrime(sq+27) {
					continue
				}

				log.Println(i)
				localSum += i
			}

			if start%1_000_000 == 0 {
				log.Println("progress --- ", i)
			}

			sum.Add(int64(localSum))
		}(i, i+bucket)
	}

	wg.Wait()

	return int(sum.Load())
}
