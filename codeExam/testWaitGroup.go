package codeExam

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func g1 () {
	fmt.Println("g1")
	wg.Done()
}

func g2 () {
	fmt.Println("g2")
	wg.Done()
}

func TestWaitgroup() {
	fmt.Println("begin")
	wg.Add(2)

	go g1()
	go g2()

	wg.Wait()

	fmt.Println("end")

}