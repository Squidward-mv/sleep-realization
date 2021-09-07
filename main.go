package main

import (
	"fmt"
	"time"
)

func sleepForGoroutine(t time.Duration, someFunc func()) {
	select {
		case <- time.After(t * time.Second):
			fmt.Println("Started")
			go someFunc()
	}
}
