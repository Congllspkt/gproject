package codeExam

import (
	"fmt"
	"sync"
)

func TestOnce1() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestOnce2() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	
	once.Do(onceBody)
	once.Do(onceBody)
	once.Do(onceBody)
	once.Do(onceBody)
	once.Do(onceBody)
	once.Do(onceBody)
	once.Do(onceBody)
	once.Do(onceBody)
	once.Do(onceBody)
	once.Do(onceBody)
	fmt.Println("hihihi")
}