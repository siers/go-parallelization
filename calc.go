package main

import "fmt"

var (
	zoom  [2]float32
	iters = 100
)

func calc(x, h int, c chan int, d chan bool) {
	var (
		tmp            float32
		zR, zI, cR, cI float32
		m              int
	)

	for i := 0; i < h; i++ {
		y := i - h/2
		cR, cI = float32(x)/zoom[0], float32(y)/zoom[1]
		zR, zI = 0.0, 0.0

		for m = 0; m < iters && 25.0 > zR*zR+zI*zI; m++ {
			tmp = zR*zR - zI*zI
			zI = 2*zR*zI + cI
			zR = tmp + cR
		}

		c <- m
	}

	d <- true
}

func Calculate(w, h int) []chan int {
	cs := make([]chan int, w)
	jobs := make(chan bool, w)
	skip := 0

	for i := 0; i < w; i++ {
		cs[i] = make(chan int, h)
		go calc(i-2*w/3, h, cs[i], jobs)

		if skip < CPUS && (w-i) > CPUS {
			skip += 1
			continue
		} else {
			<-jobs
			skip -= 1
		}
	}

	for ; skip != 0; skip-- {
		fmt.Printf("Waiting for job on queue. (%d)\n", skip)
		<-jobs
	}

	return cs
}
