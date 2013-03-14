package main

import (
	"runtime"
	"testing"
)

func Benchmark_calcs1(b *testing.B) {
	runtime.GOMAXPROCS(1)
	w := 8000
	h := 6000
	Calculate(w, h)
}

func Benchmark_calcs4(b *testing.B) {
	runtime.GOMAXPROCS(CPUS)
	w := 8000
	h := 6000
	Calculate(w, h)
	runtime.GOMAXPROCS(1)
}
