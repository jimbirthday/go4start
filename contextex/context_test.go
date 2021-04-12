package contextex

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCC(T *testing.T) {
	ctx := context.Background()

	ctxx, cancel := context.WithCancel(ctx)
	go A(ctxx)
	time.Sleep(2 * time.Second)
	cancel()
	select {
	case <-ctxx.Done():
		fmt.Println("ctx is done")
	default:
		fmt.Println("ctx waiting")
	}

	time.Sleep(2 * time.Second)
}

func A(ctx context.Context) {
	fmt.Println("ctx is start")
	time.Sleep(10 * time.Second)
	fmt.Println("ctx is end")
}

func TestCC1(T *testing.T) {
	ctx := context.Background()
	ctxx, _ := context.WithTimeout(ctx, 2*time.Second)
	go A(ctxx)
	time.Sleep(3 * time.Second)
	select {
	case <-ctxx.Done():
		fmt.Println("ctx is done")
	default:
		fmt.Println("ctx waiting")
	}

	time.Sleep(2 * time.Second)
}
