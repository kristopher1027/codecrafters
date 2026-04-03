// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [CHRISTOPHER OKOH]
// Squad:  [THE INTERFACE]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Transformation struct {
	Command string
	Input   string
	Output  string
}

var History []Transformation

func showHistory() {
	start := len(History) - 5
	if start < 0 {
		start = 0
	}
	fmt.Println("--- Last 5 Transformations ---")
	for _, t := range History[start:] {
		fmt.Printf("Command: %s, In: %s, Out: %s\n", t.Command, t.Input, t.Output)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	small := map[string]bool{"a": true, "an": true, "the": true, "of": true, "in": true, "on": true, "and": true, "to": true, "but": true, "or": true, "for": true, "nor": true, "at": true, "by": true, "up": true, "as": true, "is": true, "it": true}

	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE  ")
	for {
		fmt.Println("Valid commands: *upper <text>  *lower <text>  *cap <text>  *title <text>  *snake <text>  *reverse <text>  *count <text>  *palindrome <text>  *history  *exit")

		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			main()
			return
		}

		parts := strings.Fields(input)
		command := strings.ToLower(parts[0])

		if command == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			return
		}
		if command == "history" {
			showHistory()
			continue
		}

		if len(parts) < 2 {
			fmt.Println("ERROR:No command and text provided")
			continue
		}

		text := strings.Join(parts[1:], " ")

		var output string

		switch command {

		case "upper":
			output = strings.ToUpper(text)

		case "lower":
			output = strings.ToLower(text)

		case "cap":
			words := strings.Fields(text)
			for i, w := range words {
				words[i] = strings.ToUpper(string(w[0])) + w[1:]
			}
			output = strings.Join(words, " ")

		case "title":
			words := strings.Fields(text)
			for i, w := range words {
				lw := strings.ToLower(w)
				if i == 0 || !small[lw] {
					words[i] = strings.ToUpper(string(lw[0])) + lw[1:]
				} else {
					words[i] = lw
				}
			}
			output = strings.Join(words, " ")

		case "snake":
			var r []rune
			for _, c := range text {
				if unicode.IsLetter(c) || unicode.IsDigit(c) {
					r = append(r, unicode.ToLower(c))
				} else if unicode.IsSpace(c) {
					r = append(r, '_')
				}
			}
			output = strings.Trim(strings.ReplaceAll(string(r), "__", "_"), "_")

		case "reverse":
			words := strings.Fields(text)
			for i, w := range words {
				r := []rune(w)
				for x, y := 0, len(r)-1; x < y; x, y = x+1, y-1 {
					r[x], r[y] = r[y], r[x]
				}
				words[i] = string(r)
			}
			output = strings.Join(words, " ")

		case "palindrome":

			clean := strings.ToLower(strings.ReplaceAll(text, " ", ""))
			runes := []rune(clean)
			isPal := true

			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				if runes[i] != runes[j] {
					isPal = true
					break
				}
			}

			if isPal {
				output = fmt.Sprintln("is a palindrome!")
			} else {
				output = fmt.Sprintln(" is not a palindrome.")
			}

		case "count":

			var letters, spaces int
			for _, r := range text {
				if unicode.IsLetter(r) {
					letters++
				}
				if unicode.IsSpace(r) {
					spaces++
				}
			}
			output = fmt.Sprintf("Chars: %d, Letters: %d, Words: %d, Spaces: %d",
				len(text), letters, len(strings.Fields(text)), spaces)

		default:
			fmt.Println(" Unknown command")
			continue

		}

		History = append(History, Transformation{Command: command, Input: text, Output: output})
		fmt.Println("→", output)
	}
}
