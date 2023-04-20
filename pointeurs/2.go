/*
	Test de modification de valeur de variable externe une fonction sans retour
*/

package main

import "fmt"

func ChangeNumber(a *int, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}

func main() {
	var nba int = 5
	var nbb int = 12

	fmt.Printf("nba = %d et nbb = %d\n", nba, nbb)

	ChangeNumber(&nba, &nbb)

	fmt.Printf("nba = %d et nbb = %d\n", nba, nbb)
}
