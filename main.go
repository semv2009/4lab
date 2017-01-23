package main

import "fmt"
import "sync"	

var wg sync.WaitGroup
const recipients = 10

type Token struct {
    data string
    recipient int
}

func main() {
	start := make(chan Token)
	current := start
	wg.Add(1)
	for i := 0; i < recipients + 1; i++ {
		next := make(chan Token)
		go pushTokenBetweenGoroutine(i, current, next)
		current = next
	}
	start <- Token{"Game over", recipients}
	wg.Wait()
}

func pushTokenBetweenGoroutine(currentRecipient int, currentGoroutine <-chan Token, nextGoroutine chan<- Token) {
	for {
		token := <- currentGoroutine
		if currentRecipient == token.recipient {
			fmt.Println(token.data, currentRecipient, "goroutine")
			wg.Done()
        } else {
            fmt.Println("Push token from ", currentRecipient, " goroutine to ", currentRecipient + 1, " goroutine")
           	nextGoroutine <- token
        }
	}
}