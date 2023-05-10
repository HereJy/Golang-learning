/*
	utilisation d'un channel avec buffer
*/

package main

import (
	"fmt"
)

func main() {

	canal := make(chan string, 2) // buffer de taille 2

	go func() {
		defer close(canal) // on indique à notre compilateur qu'on a finit d'écrire sur le channel
		canal <- "test"
	}()

	for elem := range canal {
		fmt.Println(elem)
	} // permet d'afficher toutes les valeurs du canal

}
