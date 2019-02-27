package main

import (
	"fmt"
	"time"
)

func eat(left chan int, right chan int, id int) {

	for {
		fork1 := <-left
		fmt.Printf("Philosopher %d gets fork %d\n", id, fork1)
		fork2 := <-right
		fmt.Printf("Philosopher %d gets fork %d\n", id, fork2)
	
		fmt.Printf("PHILOSOPHER %d WILL EAT FOR 2 SECONDS\n", id)
		time.Sleep(2 * time.Second)
	
		fmt.Printf("Philosopher %d puts down fork %d\n", id, fork1)
		left <- fork1
	
		fmt.Printf("Philosopher %d puts down fork %d\n", id, fork2)
		right <- fork2
	}
	
}

func main() {

	fmt.Println("DINNER STARTS\n")

	ch12 := make(chan int)
	ch23 := make(chan int)
	ch34 := make(chan int)
	ch45 := make(chan int)
	ch51 := make(chan int)

	go eat(ch12, ch23, 2)
	go eat(ch23, ch34, 3)
	go eat(ch34, ch45, 4)
	go eat(ch45, ch51, 5)
	go eat(ch51, ch12, 1)

	ch12 <- 1
	ch23 <- 2
	ch34 <- 3
	ch45 <- 4
	ch51 <- 5

	//close(ch12)
	//close(ch23)
	//close(ch34)
	//close(ch45)
	//close(ch51)

	fmt.Scanln()
	
}
