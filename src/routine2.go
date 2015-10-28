//+build ignore // OMIT
package main

import (
	"fmt"
	"time"
)

func InfiniteLoop(stop chan int) chan int {
	c := make(chan int)
	i := 0
	go func() {
		for ; ; i++ {
			select {
			case <-stop:
				fmt.Println("On me demande de couper, je ferme")
				close(c)
				return
			default:
				c <- i
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	return c
}

func main() {
	stop := make(chan int)
	c := InfiniteLoop(stop)

	fmt.Println("Je me met en attente")

	go func() {
		time.Sleep(2 * time.Second)
		stop <- 1
	}()
	for v := range c {
		fmt.Println("Recu", v)
	}

}
