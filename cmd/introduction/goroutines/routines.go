package main

import (
	"fmt"
	"sync"
	"time"
)

func printHello() {
	fmt.Println("Hello from go routine")
}

func bufferedChannel() {
	charChannel := make(chan string, 3)
	charSlice := []string{"a", "b", "c"}

	for _, char := range charSlice {
		select {
		case charChannel <- char:
		}
	}
	close(charChannel)

	for char := range charChannel {
		fmt.Println(char)
	}
}

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Stop Doing work")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Doing Work")

		}
	}
}

func main() {
	start := time.Now()
	//Simple go routine
	// go printHello()
	// time.Sleep(time.Second * 1)
	// fmt.Println("Printing last statement from main")

	channel := make(chan string)
	channel2 := make(chan string)

	go func() {
		channel <- "Message sent from go routine"
	}()

	go func() {
		channel2 <- "Message from channel 2"
	}()

	//msg := <-channel
	//fmt.Println(msg)

	select {
	case msg1 := <-channel:
		fmt.Println(msg1)
	case msg2 := <-channel2:
		fmt.Println(msg2)
	}

	took := time.Since(start)
	fmt.Println("Execution Took: ", took)

	bufferedChannel()

	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 5)
	close(done)

	// Using Waitgroup.
	printOut := func(x int) {
		fmt.Println("Number", x)
	}

	arr := []int{1, 2, 3, 4, 5}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))

	for _, ele := range arr {
		go func(x int) {
			defer wg.Done()
			printOut(x)
		}(ele)
	}

	wg.Wait()

}
