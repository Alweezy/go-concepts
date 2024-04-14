package main

import "fmt"

// a mechanism for goroutines to share data
// https://youtu.be/zh7d-9O_Hd0?t=11533

// keyword chan is used to define channels

//func main() {
//	var channelOne chan int
//	fmt.Println("Channel 1 value", channelOne)
//	fmt.Printf("ChanneOne type: %T\n", channelOne)
//
//	channel2 := make(chan int)
//
//	fmt.Println("Channel 2 value", channel2)
//	fmt.Printf("Channel2 type: %T\n", channel2)
//}

// RX and TX data through channels, <- direction of arrow
// RX channelName <- element
// TX element -> channelName

//func main() {
//	fmt.Println("Main function ...")
//	ch := make(chan int)
//	go send(ch)
//
//	go receive(ch)
//
//	time.Sleep(1 * time.Second)
//
//	fmt.Println("Main function end ...")
//}
//
//func send(ch chan int) {
//	ch <- 50
//}
//
//func receive(ch chan int) {
//	fmt.Println(200 + <-ch)
//}

// closing a channel, close function()

//func main() {
//	ch := make(chan int)
//	go send(ch)
//
//	for {
//		value, ok := <-ch
//
//		if ok {
//			fmt.Printf("Channel open: %v Value: %d\n", ok, value)
//			continue
//		}
//		fmt.Printf("Channel opem: %v", ok)
//		break
//	}
//}
//
//func send(ch chan int) {
//	v := 0
//	for v < 4 {
//		ch <- v
//		v++
//	}
//	close(ch)
//}

// Buffered channels
// channelName := make(chan type capacity)
//func main() {
//	ch := make(chan string, 1)
//	ch <- "Golang"
//
//	message := <-ch
//
//	fmt.Println(message)
//}

// length and capacity of buffered channels
//func main() {
//	ch := make(chan int, 5)
//	ch <- 5
//
//	fmt.Printf("Length of channel: %d\n", len(ch))
//	fmt.Printf("Capacity of channel: %d\n", cap(ch))
//
//	ch <- 20
//
//	fmt.Printf("Length of channel: %d\n", len(ch))
//	fmt.Printf("Capacity of channel: %d\n", cap(ch))
//
//	ch <- 50
//
//	fmt.Printf("Length of channel: %d\n", len(ch))
//	fmt.Printf("Capacity of channel: %d\n", cap(ch))
//
//}

// channel as func
//bidir channel

//func main() {
//	ch := make(chan int, 1)
//	go process(ch)
//	time.Sleep(1 * time.Second)
//}
//
//func process(ch chan int) {
//	ch <- 20
//	value := <-ch
//	fmt.Println(value)
//}

// Send only channel
//func main() {
//	ch := make(chan int, 2)
//	go process(ch)
//	fmt.Println(<-ch)
//}
//
//func process(ch chan<- int) {
//	ch <- 2
//}

// Receive only channel

//func main() {
//	ch := make(chan int, 2)
//	ch <- 10
//	go receive(ch)
//	time.Sleep(1 * time.Second)
//}
//
//func receive(ch <-chan int) {
//	value := <-ch
//	fmt.Println(value)
//}

// receive all values in a channel by for range loop
//func main() {
//	ch := make(chan int, 3)
//
//	ch <- 2
//	ch <- 2
//	ch <- 2
//
//	close(ch)
//
//	sum(ch)
//	time.Sleep(200 * time.Microsecond)
//}
//
//func sum(ch chan int) {
//	sum := 0
//
//	for value := range ch {
//		sum += value
//	}
//
//	fmt.Printf("sum Total: %d\n", sum)
//}

// select statement in go
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go hello(ch1)
	go bye(ch2)

	select {
	case value1 := <-ch1:
		fmt.Println(value1)

	case value2 := <-ch2:
		fmt.Println(value2)

	}

}

func hello(ch chan<- string) {
	ch <- "Hello"
}

func bye(ch chan<- string) {
	ch <- "Goodbye"
}
