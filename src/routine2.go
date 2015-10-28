//+build ignore // OMIT
package main

import (
	"fmt"
	"time"
)

// START C OMIT
func Calc(i int, c chan int) { // HL
	fmt.Printf("Je suis la routine %d et je calcule\n", i)
	time.Sleep(1 * time.Second)
	c <- i // HL
}

func main() {
	c := make(chan int) // HL
	for i := 1; i < 5; i++ {
		go Calc(i, c) // HL
	}
	for i := 1; i < 5; i++ {
		fmt.Println("Fin de la tâche", <-c) // HL
	}	
}
// END C OMIT
