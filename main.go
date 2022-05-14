package main

import (
	"fmt"
	bar "pbar/utils"
	"time"
)

func main() {

	ch := make(chan int)
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {

		i := 0
		for {

			ch <- i
			i++
			time.Sleep(time.Duration(time.Millisecond * 1000))

		}
	}()

	go func() {
		bar.Run(5000, c1, "c1")

	}()

	go func() {
		bar.Run(20000, c2, "c2")
	}()

	go func() {
		bar.Run(5000, c3, "c3")
	}()

	for {

		select {

		case c1msg := <-c1:
			fmt.Printf(c1msg)

		case c2msg := <-c2:
			fmt.Printf(c2msg)

		case c3msg := <-c3:
			fmt.Printf(c3msg)

		case sec := <-ch:
			t := fmt.Sprintf("\033c%vs", sec+1)
			fmt.Println(t)

		}

	}

}
