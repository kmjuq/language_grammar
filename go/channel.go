package main

import (
	"fmt"
	"time"
)

func channel() {
	fmt.Println("-- channels --")
	msgChan := make(chan string)
	go func() {
		msgChan <- "ping"
	}()
	msg := <-msgChan
	fmt.Println(msg)
	// 缓冲 channel
	msgChan2 := make(chan string, 2)
	msgChan2 <- "buffered"
	msgChan2 <- "channel"

	fmt.Println(<-msgChan2)
	fmt.Println(<-msgChan2)
	// 异步channel
	done := make(chan bool, 1)
	go worker(done)
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// select 语法，主要用来处理多个channel的消费问题
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		// select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。
		// 当多个分支都准备好时会随机选择一个执行。
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 当 select 中的其它分支都没有准备好时，default 分支就会执行。
	// 为了在尝试发送或者接收时不发生阻塞，可使用 default 分支：
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
loop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break loop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	// channel close
	jobs := make(chan int, 5)
	doneChan := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				doneChan <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// close 函数会关闭channel，防止继续加入数据到channel中。
	// 只应由发送者关闭信道，而不应由接收者关闭。
	// 向一个已经关闭的信道发送数据会引发程序 panic。
	close(jobs)
	fmt.Println("sent all jobs")
	<-doneChan

	// channel 的 range 用户
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}

// -- channels --
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

// 通过chan关键左右的符号，可以限制参数只能接收或者发送消息
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
