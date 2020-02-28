package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
		arr   []string
	)
	wg.Add(4)
	go func() {
		log.Println("Sleeping 3 seconds")
		time.Sleep(3 * time.Second)
		log.Println("Slept 3 seconds")
		mutex.Lock()
		arr = append(arr, "3 seconds")
		mutex.Unlock()
		wg.Done()
	}()
	go func() {
		log.Println("Sleeping 2 seconds")
		time.Sleep(2 * time.Second)
		log.Println("Slept 2 seconds")
		mutex.Lock()
		arr = append(arr, "2 seconds")
		mutex.Unlock()
		wg.Done()
	}()
	go func() {
		log.Println("Sleeping 2 seconds")
		time.Sleep(2 * time.Second)
		log.Println("Slept 2 seconds")
		mutex.Lock()
		arr = append(arr, "2 seconds")
		mutex.Unlock()
		wg.Done()
	}()
	go func() {
		log.Println("Sleeping 2 seconds")
		time.Sleep(2 * time.Second)
		log.Println("Slept 2 seconds")
		mutex.Lock()
		arr = append(arr, "2 seconds")
		mutex.Unlock()
		wg.Done()
	}()
	wg.Wait()
	log.Printf("%v", arr)
}
