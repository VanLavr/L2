package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	//orChan recieves all the done channels and returns one channel, that closes when at least one recieved channel sent
	<-orChan(
		sig(2*time.Hour),
		sig(5*time.Second),
		sig(20*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}

func orChan(channels ...<-chan any) <-chan any {
	done := make(chan any)
	if len(channels) == 1 {
		return channels[0]
	}

	go func() {
		defer close(done)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-orChan(channels[1:]...):
			}
		}
	}()

	return done
}
