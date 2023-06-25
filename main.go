package main

import (
	"bufio"
	f "fmt"
	"math/rand"
	"os"
	"strings"
)

/*
	Text based hangman game.
			+__
			0  |
		  - | -|
		   / \ |
		   	=======

*/

func namePicker(name [10]string) string {
	min := 0
	max := 10
	pickerNumber := rand.Intn(max-min) + min

	return name[pickerNumber]
}

func gameStarter() {
	f.Print(`Text based hangman game.
			+--
			 0  |
			-|- |
			/ \ |
		=======`)
}

func nameBlanker(name string) string {
	return strings.Repeat("_", len(name))
}

func main() {
	lives := 6
	nameArray := [10]string{
		"monkey",
		"apple",
		"banana",
		"software",
		"golang",
		"python",
		"computer",
		"anaconda",
		"turkey",
		"mersin",
	}

	f.Println("Welcome to hangman game")
	secretWord := namePicker(nameArray)
	blanker := nameBlanker(secretWord)
	blankWord := []rune(blanker)
	gameStarter()

	for i := 0; i < lives; i++ {
		f.Println(string(blankWord))
		reader := bufio.NewReader(os.Stdin)
		f.Println("Guess a character or the word directly")
		char, err := reader.ReadString('\n')
		char = strings.TrimSpace(char)

		if err != nil {
			return
		}

		if char == secretWord {
			f.Println("you won")
			return
		}
		for i, v := range secretWord {
			if strings.ToLower(string(v)) == char {
				blankWord[i] = rune(v)

			}

		}

		if !strings.Contains(string(blankWord), "_") {
			break
		}
	}
	f.Printf("Word was %s", secretWord)
}
