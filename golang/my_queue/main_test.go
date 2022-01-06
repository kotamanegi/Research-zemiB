package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

type directItem struct {
	next unsafe.Pointer
	v    interface{}
}

func loaditem(p *unsafe.Pointer) *directItem {
	return (*directItem)(atomic.LoadPointer(p))
}
func casitem(p *unsafe.Pointer, old, new *directItem) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}

type Queue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
	len  uint64
}

// NewQueue creates a new lock-free queue.
func NewQueue() *Queue {
	head := directItem{next: nil, v: nil} // allocate a free item
	return &Queue{
		tail: unsafe.Pointer(&head), // both head and tail points
		head: unsafe.Pointer(&head), // to the free item
	}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *Queue) Enqueue(v interface{}) {
	i := &directItem{next: nil, v: v} // allocate new item
	var last, lastnext *directItem
	for {
		last = loaditem(&q.tail)
		lastnext = loaditem(&last.next)
		if lastnext == nil { // was tail pointing to the last node?
			if casitem(&last.next, lastnext, i) { // try to link item at the end of linked list
				//atomic.AddUint64(&q.len, 1)
				return
			}
		} else { // tail was not pointing to the last node
			casitem(&q.tail, last, lastnext) // try swing tail to the next node
		}
	}
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *Queue) Dequeue() interface{} {
	var first, last, firstnext *directItem
	for {
		first = loaditem(&q.head)
		last = loaditem(&q.tail)
		firstnext = loaditem(&first.next)
		if first == last { // is queue empty?
			if firstnext == nil { // queue is empty, couldn't dequeue
				return nil
			}
			casitem(&q.tail, last, firstnext) // tail is falling behind, try to advance it
		} else { // read value before cas, otherwise another dequeue might free the next node
			v := firstnext.v
			if casitem(&q.head, first, firstnext) { // try to swing head to the next node
				//atomic.AddUint64(&q.len, ^uint64(0))
				return v // queue was not empty and dequeue finished.
			}
		}
	}
}

var queue *Queue

func producer() {
	for i := 0; i < eachLen; i++ {
		queue.Enqueue(i)
	}
}

func consumer() {
	cnt := 0
	for cnt < (numOfProducerThread*eachLen)/numOfConsumerThread {
		hoge := queue.Dequeue()
		if hoge != nil {
			//fmt.Printf("%d %d\n", hoge, cnt)
			cnt++
		}
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
	queue = NewQueue()

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
