package main

import (
	"log"
	"time"
)

func main() {
	var sig = make(chan int)
	go func() {
		time.Sleep(10 * time.Second)
		log.Println("10 seconds have passed")
		sig <- 1
	}()
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("5 seconds have passed")
	}()
	log.Println("Waiting")
	<-sig
	log.Println("Got signal, time to go")
}
