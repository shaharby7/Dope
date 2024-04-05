package main

import (
	"context"
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	controllerName := os.Getenv("DOPE_CONTROLLER")
	controller, ok := controllers[controllerName]
	if !ok {
		panic(fmt.Sprintf("could not find controller:%s", controllerName))
	}
	err := controller.Start(context.Background(), &wg)
	if err != nil {
		panic(fmt.Sprintf("could not initiate controller:%s", err.Error()))
	}
	wg.Wait()
}
