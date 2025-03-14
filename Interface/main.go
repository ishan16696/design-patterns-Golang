package main

import (
	"context"
	"design-patterns-Golang/Interface/common"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.TODO())

	watchDog := common.NewCheckerActionWatchdog(common.ActionFunc(func(ctx context.Context) {
		// can define the action here...
		fmt.Println("Taking some action...")
	}), 2*time.Second, cancel)

	watchDog.Start(ctx)
	fmt.Println("Do something...")
	time.Sleep(6 * time.Second)
	watchDog.Stop()
}
