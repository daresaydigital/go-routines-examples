package main

import (
	"log"
	"time"
)

type returnType struct {
	id   int64
	name string
}

func main() {
	var stringThing = make(chan string)
	var typeThing = make(chan returnType)
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("3 seconds have passed")
		stringThing <- "I was returned after 3 seconds!"
	}()
	go func() {
		time.Sleep(7 * time.Second)
		log.Println("7 seconds have passed")
		typeThing <- returnType{
			id:   1,
			name: "Hello, I was created after 7 seconds",
		}
	}()
	log.Printf("First routine: %s\nSecond routine: %v\n", <-stringThing, <-typeThing)
}
