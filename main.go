package main

import (
	"fmt"
	"os"
	"strings"
)

// Function to convert a character to Morse code
func charToMorse(c rune) string {
	morseCode := map[rune]string{
		'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
		'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
		'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
		'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
		'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
		'Z': "--..",
		'0': "-----", '1': ".----", '2': "..---", '3': "...--",
		'4': "....-", '5': ".....", '6': "-....", '7': "--...",
		'8': "---..", '9': "----.",
	}

	if code, ok := morseCode[c]; ok {
		return code
	}
	return ""
}

// Function to convert a word to Morse code
func wordToMorse(word string) string {
	var morse strings.Builder
	for _, c := range strings.ToUpper(word) {
		morse.WriteString(charToMorse(c) + " ")
	}
	return strings.TrimSpace(morse.String())
}

func main() {
	// Check for command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <string>")
		os.Exit(1)
	}

	input := strings.Join(os.Args[1:], " ")
	morseCode := wordToMorse(input)
	fmt.Println("Morse Code:", morseCode)
}
