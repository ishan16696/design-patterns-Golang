package common

import (
	"context"
	"fmt"
	"time"
)

// Action is a context-aware action.
type Action interface {
	// Do performs an action.
	Do(ctx context.Context)
}

// ActionFunc is a function that implements Action.
type ActionFunc func(ctx context.Context)

// Do performs an action.
func (f ActionFunc) Do(ctx context.Context) {
	f(ctx)
}

type checkerActionWatchdog struct {
	action     Action
	interval   time.Duration
	cancelFunc context.CancelFunc
}

// Watchdog manages a goroutine that can be started and stopped.
type Watchdog interface {
	// Start starts a goroutine using the given context.
	Start(ctx context.Context)
	// Stop stops the goroutine started by Start.
	Stop()
}

func NewCheckerActionWatchdog(action Action, interval time.Duration, cancel context.CancelFunc) Watchdog {
	return &checkerActionWatchdog{
		action:     action,
		interval:   interval,
		cancelFunc: cancel,
	}
}

func (c *checkerActionWatchdog) Start(ctx context.Context) {

	go func() {
		fmt.Println("WatchDog starts..")
		for {
			select {
			case <-time.After(c.interval):
				c.action.Do(ctx)

			case <-ctx.Done():
				fmt.Println("Shutting down...")
				return
			}
		}
	}()
}

func (c *checkerActionWatchdog) Stop() {
	fmt.Println("Stopping the watchDog")
	c.cancelFunc()
}
