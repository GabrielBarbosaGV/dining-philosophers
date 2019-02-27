package main

import (
	"fmt"
	"time"
	"math/rand"
)

const maxSleepTimeMillis, maxEatingTimeMillis int = 200, 400
const debugPrintOuts bool = true

func newPhilosopher(philNum int) func(chan bool, []bool) {
	var philosopherNum int = philNum

	return func(mutex chan bool, forks []bool) {
		var hasForks bool = false

		for {
			// Think
			if debugPrintOuts { fmt.Println("Philosopher", philosopherNum, "has started thinking...") }
			sleepTimeMillis := rand.Intn(maxSleepTimeMillis)
			time.Sleep(time.Duration(sleepTimeMillis) * time.Millisecond)
			if debugPrintOuts { fmt.Println("Philosopher", philosopherNum, "is done thinking.") }

			// Ask for permission to get forks
			if debugPrintOuts { fmt.Println("Philosopher", philosopherNum, "asking for permission to eat...") }
			lock := <- mutex
			if debugPrintOuts { fmt.Println("Philosopher", philosopherNum, "was granted permission!") }
			if forks[philosopherNum] && forks[(philosopherNum + 1) % len(forks)] {
				if debugPrintOuts { fmt.Println("Philosopher", philosopherNum, "had available forks, getting forks...") }
				// It is now FALSE that the left and right forks are available
				forks[philosopherNum] = false
				forks[(philosopherNum + 1) % len(forks)] = false
				hasForks = true
				if debugPrintOuts { fmt.Println("Philosopher", philosopherNum, "is done getting forks.") }
			} else if debugPrintOuts {
				fmt.Println("Philosopher", philosopherNum, "did not have available forks, thinking again...")
			}
			mutex <- lock

			// Eat if philosopher has forks, otherwise think again
			if hasForks {
				if debugPrintOuts { fmt.Println("Philosopher", philosopherNum, "has started eating...") }
				eatingTimeMillis := rand.Intn(maxEatingTimeMillis)
				time.Sleep(time.Duration(eatingTimeMillis) * time.Millisecond)
				if debugPrintOuts { fmt.Println("Philosopher", philosopherNum, "is done eating.") }

				// Put forks down
				if debugPrintOuts { fmt.Println("Philosopher", philosopherNum, "Is putting his forks back down.") }
				forks[philosopherNum] = true
				forks[(philosopherNum + 1) % len(forks)] = true
				hasForks = false
			}
		}
	}
}

func main() {
	var philosopherAmount int = 5
	var mutex chan bool
	mutex = make(chan bool)

	var philosopherRoutines []func (chan bool, []bool) =
		make([]func (chan bool, []bool), philosopherAmount, philosopherAmount)

	var forks []bool = make([]bool, philosopherAmount, philosopherAmount)

	// Initially, all forks are available
	for i := range forks { forks[i] = true }

	for i := range philosopherRoutines {
		philosopherRoutines[i] = newPhilosopher(i)
		go philosopherRoutines[i](mutex, forks)
	}

	mutex <- false

	fmt.Scanln()
}