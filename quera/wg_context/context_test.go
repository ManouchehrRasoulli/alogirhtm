package wg_context

import (
	"context"
	"testing"
	"time"
)

func TestWithWaitGroup(t *testing.T) {
	var cf func()
	ctx := context.Background()
	ctx = WithWaitGroup(ctx)
	ctx, cf = context.WithCancel(ctx)

	action := func(id int, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				t.Log("action :: ", id, " exiting ....")
				time.Sleep(5 * time.Second)
				DoneContext(ctx)
				return
			default:
				t.Log("action :: ", id, " running ...")
				time.Sleep(time.Second)
			}
		}
	}

	for i := 0; i < 3; i++ {
		AddContext(ctx, 1)
		go action(i, ctx)
	}

	time.Sleep(10 * time.Second)

	// trigger the exit
	cf()

	WaitContext(ctx)
}
