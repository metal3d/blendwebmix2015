//+build ignore // OMIT
package main

import (
	"log"
	"net/http" // HL
)

func Hello(w http.ResponseWriter, r *http.Request) { // HL
	log.Println("Recieved connection")
	w.Write([]byte("Hello Blend Web Mix !"))
}

func main() {
	http.HandleFunc("/", Hello) // HL
	log.Println("Listening")
	http.ListenAndServe(":8000", nil)
}
