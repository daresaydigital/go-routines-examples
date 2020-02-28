# Parallelism in Go
Short examples of parallelism in Go. All examples uses go routines to create parallelism but
uses different methods for handling the locking and waiting mechanisms.

## Signaling
A simple signal is used to tell the parent thread when the child is finished.

To run the example: `$go run signaling/main.go`

## Channels
Channels are used to wait for all return values to be set before the program can continue.

To run the example: `$go run channels/main.go`

## Mutex and Wait Groups
A shared waitgroup is used in 4 go routines and the parent thread will wait until all are done.
A shared mutex is used to prevent `append()` to overwrite the same slice and let all functions 
write their results to the same dataset without interferring with each other, even when executing
at the same time.

To run the example: `$go run mutexAndWaitGroup/main.go`