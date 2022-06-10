package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")

	//childCtx, cancel := context.WithCancel(valCtx)
	childCtx, cancel := context.WithTimeout(valCtx, 30*time.Second)
	defer cancel()

	wg.Add(1)
	go fn(childCtx, wg)
	wg.Wait()
	fmt.Println("main completed")
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)

	childCtx1, cancel1 := context.WithCancel(ctx)
	defer cancel1()
	go f1(childCtx1, wg)

	//go f1(ctx, wg)

	wg.Add(1)

	childCtx2, cancel2 := context.WithCancel(ctx)
	defer cancel2()
	go f2(childCtx2, wg)

	//go f2(ctx, wg)

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			break LOOP
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println("exiting fn")
}

func f1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			break LOOP
		default:
			fmt.Print("f1")
			time.Sleep(400 * time.Millisecond)
		}
	}
	fmt.Println("exiting f1")
}

func f2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Root value = ", ctx.Value("root-key"))
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			break LOOP
		default:
			fmt.Print("f2")
			time.Sleep(1000 * time.Millisecond)
		}
	}
	fmt.Println("exiting f2")
}
