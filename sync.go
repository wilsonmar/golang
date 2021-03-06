/* sync.go in https://github.com/wilsonmar/sync.go
   Based on https://stackoverflow.com/questions/16020406/better-go-idiomatic-way-of-writing-this-code
   This makes use of sync for goroutines.
   A goroutine is a function that is capable of running concurrently with other functions.
   Channels provide a way for goroutines to communicate with one another so they can synchronize their execution.
   See tutorials https://tour.golang.org/concurrency/1
   and https://www.golang-book.com/books/intro/10 and https://www.golang-book.com/books/intro/13#section9
*/

package main

import (
    "fmt"
    "sync" // to synchronize access to memory shared by Goroutines using the same address space
)

const N = 10  // 10 channels.

func main() {
    // Make a new buffered channel named "ch" using the "chan" type which is an integer for goroutines to communicate:
    ch := make(chan int, N)
    // Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

    // A wait group waits for a collection of goroutines to finish:
    var wg sync.WaitGroup  // see https://golang.org/pkg/sync/#WaitGroup
    for i := 0; i < N; i++ {
        // The main goroutine calls Add to set the number of goroutines to wait for.
        wg.Add(1)
        // Create a goroutine using the keyword go followed by a function invocation:
        go func(n int) {
            // Each of the goroutines runs and calls Done when finished.
            defer wg.Done()
            for i := 0; i < N; i++ {
                // Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
                // Send n to a channel named ch (data flows in the direction of the arrow):
                ch <- n*N + i
            }
        }(i)
    }
    go func() { // At the same time, Wait is used to block until all goroutines have finished.
        wg.Wait()
        close(ch) // Only the sender should close a channel, never the receiver.
        // Sending on a closed channel will cause a panic.
        // Channels aren't like files; you don't usually need to close them.
        // Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
    }()
    for i := range ch {
        fmt.Println(i)
    }
}
