package main

import (
	"fmt"
	pubsub "pubsub/internal/pub"
	"time"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()

	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return len(s) > 10
		}
		return false
	})
	p.Publish("hello, world!")
	p.Publish("hello, golang!")
	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()
	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()
	time.Sleep(3 * time.Second)
}
