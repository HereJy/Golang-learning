package main

import (
	"fmt"
)

func main() {
	canal := make(chan string) // créé un channel de string
	go run(canal, "Nom 1")
	go run(canal, "Nom 2")
	fmt.Println(<-canal) // récupère la valeur dans le canal
	fmt.Println(<-canal)
}

func run(c chan string, name string) {
	c <- name // envoie la valeur name dans le canal c
}
