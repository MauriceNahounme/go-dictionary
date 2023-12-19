package main

import (
	"dictionnaire/dictionary"
	"fmt"
)

func main() {
	// Création d'un dictionnaire
	d := dictionary.New()

	// Ajout de mots et de définitions au dictionnaire
	d.Add("go", "A programming language")
	d.Add("python", "Another programming language")

	// Affichage de la définition d'un mot spécifique
	wordToGet := "go"
	definition, found := d.Get(wordToGet)
	if found {
		fmt.Printf("Definition of %s: %s\n", wordToGet, definition)
	} else {
		fmt.Printf("%s not found in the dictionary\n", wordToGet)
	}

	// Suppression d'un mot du dictionnaire
	wordToRemove := "python"
	fmt.Printf("Removing %s from the dictionary...\n", wordToRemove)
	d.Remove(wordToRemove)

	// Affichage de la liste triée des mots et de leurs définitions
	fmt.Println("List of words and definitions:")
	wordList := d.List()
	for _, entry := range wordList {
		fmt.Println(entry)
	}
}
