/*
	Pour tester les goroutines
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	debut := time.Now()
	run("routine 1")
	run("routine 2")
	run("routine 3")
	fin := time.Now()
	fmt.Println(fin.Sub(debut))
}

func run(routine string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%s : %d\n", routine, i)
	}
}
