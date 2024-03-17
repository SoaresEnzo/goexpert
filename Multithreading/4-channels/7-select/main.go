package main

import (
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{Id: i, Msg: "Hello from RabbitMQ"}
			//
			c1 <- msg
		}

	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{Id: i, Msg: "Hello from Kafka"}
			c2 <- msg
		}
	}()

	// for i := 0; i < 3; i++ {
	for {
		select {
		case msg := <-c1:
			println("received", msg.Msg, msg.Id)

		case msg := <-c2:
			println("received", msg.Msg, msg.Id)

		case <-time.After(3 * time.Second):
			println("timeout")

			// default:
			// 	println("default")
		}
	}

}
