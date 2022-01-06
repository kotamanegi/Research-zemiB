package main

import (
	"sync"
	"testing"
)

var mu sync.Mutex
var queue []int

func producer() {
	for i := 0; i < eachLen; i++ {
		mu.Lock()
		queue = append(queue, i)
		mu.Unlock()
	}
}

func consumer() {
	cnt := 0
	for cnt < numOfProducerThread*eachLen/numOfConsumerThread {
		mu.Lock()
		if len(queue) != 0 {
			//fmt.Printf("%d %d\n", queue[0], cnt)
			queue = queue[1:]
			cnt++
		}
		mu.Unlock()
	}
}

var eachLen = 10000
var numOfProducerThread = 1
var numOfConsumerThread = 1

func runTest(len, prod, cons int, b *testing.B) {
	eachLen = len
	numOfProducerThread = prod
	numOfConsumerThread = cons
	var wg sync.WaitGroup
	queue = make([]int, 0)
	b.ResetTimer()
	for i := 0; i < numOfProducerThread; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			producer()
		}()
	}
	for i := 0; i < numOfConsumerThread; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			consumer()
		}()
	}
	wg.Wait()
}

func Benchmark_N_100_Thread_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(100, 1, 1, b)
	}
}

func Benchmark_N_100_Thread_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(100, 10, 10, b)
	}
}

func Benchmark_N_100_Thread_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(100, 100, 100, b)
	}
}
func Benchmark_N_100_Thread_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(100, 1000, 1000, b)
	}
}
func Benchmark_N_100_Thread_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(100, 10000, 10000, b)
	}
}

func Benchmark_N_1000_Thread_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(1000, 1, 1, b)
	}
}

func Benchmark_N_1000_Thread_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(1000, 10, 10, b)
	}
}

func Benchmark_N_1000_Thread_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(1000, 100, 100, b)
	}
}
func Benchmark_N_1000_Thread_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(1000, 1000, 1000, b)
	}
}
func Benchmark_N_1000_Thread_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(1000, 10000, 10000, b)
	}
}

func Benchmark_N_10000_Thread_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(10000, 1, 1, b)
	}
}
func Benchmark_N_10000_Thread_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(10000, 10, 10, b)
	}
}
func Benchmark_N_10000_Thread_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(10000, 100, 100, b)
	}
}
func Benchmark_N_10000_Thread_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(10000, 1000, 1000, b)
	}
}
func Benchmark_N_10000_Thread_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(10000, 10000, 10000, b)
	}
}
func Benchmark_N_100000_Thread_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(100000, 1, 1, b)
	}
}

func Benchmark_N_100000_Thread_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(100000, 10, 10, b)
	}
}

func Benchmark_N_100000_Thread_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(100000, 100, 100, b)
	}
}

func Benchmark_N_100000_Thread_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(100000, 1000, 1000, b)
	}
}

func Benchmark_N_100000_Thread_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(100000, 10000, 10000, b)
	}
}
