package bench

import (
	"bytes"
	"compress/gzip"
	"testing"
)

// % go test -bench . -benchmem                                                                                                                                                            (arn:aws:eks:ap-northeast-1:354265976659:cluster/diarkis/default)
// goos: darwin
// goarch: arm64
// pkg: go-study.example.com/bench
// Benchmark_globalBuf-10             20610             61700 ns/op          814029 B/op         17 allocs/op
// Benchmark_localBuf-10              31335             40385 ns/op          813986 B/op         19 allocs/op

var globalBuf = &bytes.Buffer{}
var str = []byte("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

func Benchmark_globalBuf(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		gz := gzip.NewWriter(globalBuf)
		gz.Write(str)
		gz.Flush()
		gz.Close()
		gz.Reset(globalBuf)
	}
}

func Benchmark_localBuf(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var buff bytes.Buffer
		gz := gzip.NewWriter(&buff)
		gz.Write(str)
		gz.Flush()
		gz.Close()
	}
}
