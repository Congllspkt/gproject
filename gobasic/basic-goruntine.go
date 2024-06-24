package gobasic

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func countNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func RunGoRuntine() {
	go sayHello()    
	go countNumbers() 
	time.Sleep(6 * time.Second)
	fmt.Println("Done")
}