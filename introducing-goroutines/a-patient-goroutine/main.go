package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job, goGroup *sync.WaitGroup) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	// Telling the waitGroup that this function
	// is done
	goGroup.Done()
}

func main() {

	// Initiating a waitGroup which will
	// allow us to wait for our Goroutines
	// to finish
	goGroup := new(sync.WaitGroup)
	fmt.Println("Starting")

	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3

	world := new(Job)
	world.text = "world"
	world.i = 0
	world.max = 5

	go outputText(hello, goGroup)
	go outputText(world, goGroup)

	// goGroup.Add() lets us specify how
	// many Done() messages we should
	// receive before finishing the
	// function
	goGroup.Add(2)

	// Wait tells our function to wait
	// until the specified amount of
	// goroutines have completed and
	// called Done() on our waitGroup
	goGroup.Wait()
}
