package main

import (
	"fmt"

	"training.go/hangman/hangman"
)

func main() {

	g := hangman.New(8, "Golang")
	fmt.Println(g)
}
