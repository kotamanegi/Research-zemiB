package main

import (
	"fmt"
	"sync"

	"github.com/golang-design/lockfree"
)

const eachLen = 100
const numOfThread = 10

var queue *lockfree.Queue

func producer() {
	for i := 0; i < eachLen; i++ {
		queue.Enqueue(i)
	}
}

func consumer() {
	cnt := 0
	for cnt < numOfThread*eachLen {
		hoge := queue.Dequeue()
		if hoge != nil {
			fmt.Printf("%d %d\n", hoge, cnt)
			cnt++
		}
	}
}

func main() {
	var wg sync.WaitGroup
	queue = lockfree.NewQueue()
	for i := 0; i < numOfThread; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			producer()
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		consumer()
	}()
	wg.Wait()
}
