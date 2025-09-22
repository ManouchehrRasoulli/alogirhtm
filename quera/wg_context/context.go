package wg_context

import (
	"context"
	"sync"
)

func WithWaitGroup(parent context.Context) context.Context {
	return context.WithValue(parent, "wg", &sync.WaitGroup{})
}

func DoneContext(parent context.Context) {
	wa := parent.Value("wg")
	if wa != nil {
		wg := wa.(*sync.WaitGroup)
		wg.Done()
	}
}

func AddContext(parent context.Context, value int) {
	wa := parent.Value("wg")
	if wa != nil {
		wg := wa.(*sync.WaitGroup)
		wg.Add(value)
	}
}

func WaitContext(parent context.Context) {
	wa := parent.Value("wg")
	if wa != nil {
		wg := wa.(*sync.WaitGroup)
		wg.Wait()
	}
}
