package kata

import (
	"log"
	"math"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

var (
	rotated = map[int64]int64{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}
)

func isReversible(n int64) (bool, int64) {
	var (
		count        int64
		invalidIndex int64
		isValid      = true
	)

	for n != 0 {
		c := n % 10
		if _, ok := rotated[c]; !ok {
			invalidIndex = count
			isValid = false
		}

		n /= 10
		count++
	}

	return isValid, invalidIndex
}

func reverse(n int64) int64 {
	var (
		reversed int64
	)

	for n != 0 {
		c := n % 10
		v := rotated[c]
		reversed *= 10
		reversed += v
		n /= 10
	}

	return reversed
}

func _UpsideDown(n1, n2 string) uint64 {
	var (
		start      = int64(0)
		end        = int64(0)
		count      = atomic.Uint64{}
		partitions = int64(0)
		wg         = sync.WaitGroup{}
		numCpu     = int64(runtime.NumCPU())
	)

	start, _ = strconv.ParseInt(n1, 10, 64)
	end, _ = strconv.ParseInt(n2, 10, 64)
	partitions = (end - start + numCpu) / numCpu

	log.Println("start :", start, "end :", end, "partitions :", partitions, "num-cpu :", numCpu)

	for i := int64(0); i < numCpu; i++ {
		partStart := start + i*partitions
		partEnd := partStart + partitions - 1
		if partEnd > end {
			partEnd = end
		}

		wg.Add(1)
		go func(start, end int64) {
			defer wg.Done()
			ct := uint64(0)
			for i := start; i <= end; {
				if ok, index := isReversible(i); !ok {
					nextIndex := math.Pow(10, float64(index))
					remaining := i % int64(nextIndex)
					next := int64(nextIndex) - remaining
					i += next
					continue
				}

				if reverse(i) == i {
					ct++
				}
				i++
			}

			count.Add(ct)
		}(partStart, partEnd)
	}

	wg.Wait()

	return count.Load()
}
