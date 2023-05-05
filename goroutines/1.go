/*
	Pour tester les goroutines
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	debut := time.Now()
	// wg.Add permet d'ajouter une goroutine à attendre
	wg.Add(1)
	go run("routine 1")
	wg.Add(1)
	go run("routine 2")
	wg.Add(1)
	go run("routine 3")
	wg.Wait() // permet d'attendre la fin des routines avant de reprendre la routine principale
	fin := time.Now()
	fmt.Println(fin.Sub(debut))
}

func run(routine string) {
	defer wg.Done() // s'execute même en cas de crash grace au defer
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%s : %d\n", routine, i)
	}
}
