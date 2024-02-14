package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Marx", rightFork: 4, leftFork: 0},
	{name: "Kant", rightFork: 0, leftFork: 1},
	{name: "Turing", rightFork: 1, leftFork: 2},
	{name: "Descartes", rightFork: 2, leftFork: 3},
	{name: "Wittgenstein", rightFork: 3, leftFork: 4},
}

var hunger = 3 // Number of times each philosopher will eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	// print out a welcome message
	fmt.Println("Welcome to the philosophers' dinner!")
	fmt.Println("-----------------------------------")
	fmt.Println("The table is empty.")

	// start the meal
	dine()

	// print out finished message
	fmt.Println("-----------------------------------")
	fmt.Println("The philosophers have finished dining.")

}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// forks is a map of all five forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
}
