/*
	création d'une structure permettant de créer un personnage et de méthodes qui agissent dessus
*/

package main

import "fmt"

type Character struct {
	name       string
	lifePoints int
	level      int
	weapon     [2]string
	isDead     bool
}

func main() {
	var char1, char2 Character

	char1.InitChar("Guts", 100, 50, [2]string{"dragonslayer", "behelit"}, false)
	char1.DisplayChar()

	fmt.Printf("\n--------------------------------------------------------------\n")

	char2.InitChar("Griffith", 100, 80, [2]string{"casque", "behelit pourpre"}, false)
	char2.DisplayChar()
}

// Méthode agissant sur le personnage pour initialiser ses valeurs
func (char *Character) InitChar(name string, lp int, lvl int, wp [2]string, isdead bool) {
	char.name = name
	char.lifePoints = lp
	char.level = lvl
	char.weapon = wp
	char.isDead = isdead
}

// Méthode agissant sur le personnage pour afficher ses valeurs
func (char *Character) DisplayChar() {
	fmt.Printf("Nom : %s\nPoints de vie : %d\nNiveau : %d\n", char.name, char.lifePoints, char.level)

	fmt.Printf("Il a dans son inventaire :\n")
	for i := range char.weapon {
		fmt.Printf("\t- %s\n", char.weapon[i])
	}

	if char.isDead == true {
		fmt.Printf("Il est mort\n")
	} else {
		fmt.Printf("Il est en vie\n")
	}
}
