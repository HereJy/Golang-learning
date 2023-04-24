package main

import "fmt"

func main() {
	var week = []string{"Lundi", "Mardi", "Mercredi", "Vendredi", "Samedi"}

	fmt.Printf("%s\n", week)
	week = append(week, "Dimanche")
	fmt.Printf("%s\n", week)

	var slice = make([]int, 10)

	fmt.Printf("\n%d\n", slice)

	for i := range slice {
		slice[i] = i + 1
	}

	fmt.Printf("%d\n", slice)

	// copie d'un slice
	var week2 = make([]string, len(week))
	copy(week2, week)
	fmt.Printf("\nweek2 : %s\n", week2)

	// Suppression d'un jour (le mardi)
	// le tab[:1] permet d'afficher tous les éléments jusqu'à l'indice donné (celui-ci exclu)
	week2 = append(week2[:1], week2[2:]...)
	fmt.Printf("\nweek2 après suppression du lunid : %s\n", week2)
}
