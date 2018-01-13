package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func showNumber(num int) {
	tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(num, tstamp)
	time.Sleep(time.Millisecond * 10)
}

// runtime.Gosched() temporarily yields the processor and then
// returns to the current goroutine.
func main() {
	runtime.GOMAXPROCS(2)
	iterations := 10

	for i := 0; i <= iterations; i++ {
		go showNumber(i)
	}

	fmt.Println("Goodbye!")
	runtime.Gosched()
}
