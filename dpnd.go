package main

import (
	//"fmt"
	"time"
	"math/rand"
)

const maxSleepTimeMillis, maxEatingTimeMillis int = 200, 400

func newPhilosopher(philNum int) func(chan bool, []bool) {
	var philosopherNum int = philNum

	return func(mutex chan bool, forks []bool) {
		var hasForks bool = false

		for {
			// Think
			sleepTimeMillis := rand.Intn(maxSleepTimeMillis)
			time.Sleep(time.Duration(sleepTimeMillis) * time.Millisecond)

			// Ask for permission to get forks
			lock := <- mutex
			if forks[philosopherNum] && forks[(philosopherNum + 1) % len(forks)] {

				// It is now FALSE that the left and right forks are available
				forks[philosopherNum] = false
				forks[(philosopherNum + 1) % len(forks)] = false
				hasForks = true
			}
			mutex <- lock

			// Eat if philosopher has forks, otherwise think again
			if hasForks {
				eatingTimeMillis := rand.Intn(maxEatingTimeMillis)
				time.Sleep(time.Duration(eatingTimeMillis) * time.Millisecond)

				// Put forks down
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

	for {}
}