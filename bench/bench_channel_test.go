package bench

import (
	"sync"
	"testing"
	"time"
)

// % go test -bench . -benchmem
// goos: darwin
// goarch: arm64
// Benchmark_WaitGroup-10               230           5228995 ns/op             328 B/op            8 allocs/op
// Benchmark_Channel-10                 229           5242579 ns/op             430 B/op            8 allocs/op

type Async struct {
	sync.WaitGroup
}

func Benchmark_WaitGroup(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = doWithWaitGroup()
	}
}

func Benchmark_Channel(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = doWithChannel()
	}
}

func doWithWaitGroup() []byte {
	var res []byte
	async := new(Async)
	async.Add(1)
	go func() {
		time.Sleep(5 * time.Millisecond)
		if async == nil {
			return
		}
		res = []byte("hello")
		async.Done()
	}()
	go func() {
		time.Sleep(15 * time.Millisecond)
		if async == nil {
			return
		}
		res = []byte("world")
		async.Done()
	}()
	async.Wait()
	async = nil
	return res
}

func doWithChannel() []byte {
	var res []byte
	done := make(chan struct{}, 1)
	timeoutChan := time.After(15 * time.Millisecond)
	go func() {
		time.Sleep(5 * time.Millisecond)
		res = []byte("hello")
		done <- struct{}{}
		close(done)
	}()
	select {
	case <-done:
		return res
	case <-timeoutChan:
		return []byte("world")
	}
}
