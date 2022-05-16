package main

import "fmt"

//输出10以内的素数
func main() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Printf("ch中有值 %d\n", x)
		case ch <- i:
		}
	}
}
