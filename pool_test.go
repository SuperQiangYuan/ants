package ants

import (
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	ch := make(chan int, 0)

	go func() {
		for s := range ch {
			time.Sleep(time.Second)
			t.Log(s)
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		t.Log("finish")
	}

	select {}
}
