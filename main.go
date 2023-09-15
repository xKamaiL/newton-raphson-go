package main

import (
	"fmt"
	"math"
	"time"
)

func Do(f, df func(x float64) float64, init, delta float64) (r float64) {
	r = init // initial value
	step := func() float64 {
		return r - f(r)/df(r)
	}
	for rr := step(); math.Abs(rr-r) > delta; {
		r = rr
		rr = step()
	}
	return
}

func main() {
	// define function
	// and derivative of function
	f := func(x float64) float64 {
		return math.Sin(x-0.361) + (math.Sin(0.361) * math.Exp(-x/0.377))
	}
	df := func(x float64) float64 {
		return math.Cos(x-0.361) + (math.Sin(0.367) * (-1 / 0.377) * math.Exp(-x/0.377))
	}

	// measure time
	since := time.Now()
	defer func() {
		fmt.Println(time.Since(since))
	}()
	// initial value is math.Pi
	// precision is 0.0000001
	fmt.Println(
		Do(f, df, 3*math.Pi/4, 0.0000001),
	)
}
