package time_bucket

import (
	"math/rand"
	"testing"
	"time"
)

func randomDuration() time.Duration {
	// random float64 between 0.0 and 5.0 seconds
	sec := rand.Float64() * 5

	// convert to duration with millisecond precision
	return time.Duration(sec * float64(time.Second))
}

func Test_CaptureResponseTime(t *testing.T) {
	CaptureResponseTime(randomDuration())
	CaptureResponseTime(randomDuration())
	t.Log(ListResponseTimes())
}

func Test_CaptureResponseTimeAndListResponseTimes(t *testing.T) {
	for i := 0; i < 100; i++ {
		CaptureResponseTime(randomDuration())
		if i%10 == 0 {
			t.Log(ListResponseTimes())
		}
	}

	t.Log(ListResponseTimes())
}

func Benchmark_CaptureResponseTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CaptureResponseTime(randomDuration())
	}
	b.Log(ListResponseTimes())
}
