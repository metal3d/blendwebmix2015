//+build ignore // OMIT
package main

import (
	"fmt"
	"time"
)

func Calc(i int) {
	fmt.Printf("Je suis la routine %d et je calcule\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d à terminé\n", i)
}

func main() {
	for i := 1; i < 5; i++ {
		Calc(i)
	}
	// time.Sleep(1.1e9)
}
