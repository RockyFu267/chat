package channel

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send：", i)
	}
}

func consumer(c <-chan int) {
	for {
		v := <-c
		fmt.Println("receive:", v)
	}
}

func chan2() {
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}
