package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// Function to pick a random word from a predefined list
func getRandomWord() string {
	words := []string{"hangman", "golang", "programming", "developer", "challenge"}
	index := rand.Intn(len(words))
	return words[index]
}

// Function to display the current state of the word with blanks for unknown letters
func displayWord(word string, guessedLetters []string) string {
	display := ""
	for _, letter := range word {
		if contains(guessedLetters, string(letter)) {
			display += string(letter)
		} else {
			display += "_"
		}
		display += " "
	}
	return display
}

// Function to check if a letter is in the guessed letters list
func contains(slice []string, letter string) bool {
	for _, l := range slice {
		if l == letter {
			return true
		}
	}
	return false
}

func main() {
	// Set up the Hangman game
	maxAttempts := 6
	wordToGuess := getRandomWord()
	guessedLetters := []string{}
	attempts := 0

	fmt.Println("Welcome to Hangman!")

	for {
		// Display the current state of the word
		fmt.Println(displayWord(wordToGuess, guessedLetters))

		// Display guessed letters
		fmt.Printf("Guessed letters: %s\n", strings.Join(guessedLetters, ", "))

		// Get user input
		fmt.Print("Enter a letter: ")
		var guess string
		fmt.Scanln(&guess)

		// Check if the guessed letter is in the word
		if strings.Contains(wordToGuess, guess) {
			fmt.Println("Correct guess!")
		} else {
			fmt.Println("Incorrect guess.")
			attempts++
		}

		// Add the guessed letter to the list
		guessedLetters = append(guessedLetters, guess)

		// Check if the player has won or lost
		if displayWord(wordToGuess, guessedLetters) == wordToGuess {
			fmt.Println("Congratulations! You guessed the word:", wordToGuess)
			break
		} else if attempts == maxAttempts {
			fmt.Println("Sorry, you've run out of attempts. The correct word was:", wordToGuess)
			break
		}
	}
}
