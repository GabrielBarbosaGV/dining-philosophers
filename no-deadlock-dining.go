package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Create a Philosopher
type Philosopher struct {
	name     string
	fork     int
	neighbor *Philosopher
}

func createPhilosopher(name string, neighbor *Philosopher) *Philosopher {
	thinkerMan := &Philosopher{name, 1, neighbor}
	return thinkerMan
}

//Philo get fork
func (thinkerMan *Philosopher) getFork() {
	if thinkerMan.fork == 1 && thinkerMan.neighbor.fork == 1 {
		thinkerMan.neighbor.fork = 0
		thinkerMan.fork = 0
		fmt.Printf("%v pegou o garfo dele e o garfo do %v", thinkerMan.name, thinkerMan.neighbor.name)
		thinkerMan.eat()
	} else {
		thinkerMan.think()
	}
}

//Philo eat
func (thinkerMan *Philosopher) eat() {
	fmt.Printf("%v está comendo", thinkerMan.name)
	time.Sleep(time.Duration(rand.Int63n(5000-1000) + 1000))
	fmt.Printf("%v terminou de comer", thinkerMan.name)

}

//Philo thinks
func (thinkerMan *Philosopher) think() {
	fmt.Printf("%v está pensando...", thinkerMan.name)
	time.Sleep(time.Duration(rand.Int63n(5000-1000) + 1000))
	thinkerMan.getFork()
}

//Philo returns forks
func (thinkerMan *Philosopher) returnFork() {
	thinkerMan.fork = 1
	thinkerMan.neighbor.fork = 1
	fmt.Printf("%v devolveu o garfo dele e o garfo do %v", thinkerMan.name, thinkerMan.neighbor.name)
	thinkerMan.think()
}

func main() {

}
