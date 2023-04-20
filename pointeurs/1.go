/*
	Un simple de test avec les pointeurs pour comprendre comment jouer avec les adresses
*/

package main

import "fmt"

func main() {
	var nba int = 10
	var nbb *int = &nba

	fmt.Printf("nba est à l'adresse : %p, nbb à l'adresse : %p et nbb pointe sur l'adresse de nba qui est : %p\n", &nba, &nbb, nbb)
	fmt.Printf("\nnba = %d et nbb = %d\n", nba, *nbb)

	*nbb = 5 //changement de la valeur stocké à l'adresse de la variable nba via le pointeur nbb

	fmt.Printf("nba = %d et nbb = %d\n", nba, *nbb)
}
