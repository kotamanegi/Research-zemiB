package main

import (
	"fmt"
	"sync"
)

const eachLen = 100
const numOfThread = 10

var queue []int

func producer() {
	for i := 0; i < eachLen; i++ {
		queue = append(queue, i)
	}
}

func consumer() {
	cnt := 0
	for cnt < numOfThread*eachLen {
		if len(queue) != 0 {
			fmt.Printf("%d %d\n", queue[0], cnt)
			queue = queue[1:]
			cnt++
		}
	}
}

func main() {
	var wg sync.WaitGroup
	queue = make([]int, 0)
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
