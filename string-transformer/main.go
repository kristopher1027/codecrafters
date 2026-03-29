package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	small := map[string]bool{"a": true, "an": true, "the": true, "of": true, "in": true, "on": true, "and": true, "to": true}

	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE {sample: upper <text>  lower <text> ")

	for {
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
		cmd := strings.ToLower(parts[0])

		if cmd == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			return
		}

		if len(parts) < 2 {
			fmt.Printf("✗ No text provided. Usage: %s <text>\n", cmd)
			continue
		}

		text := strings.Join(parts[1:], " ")

		var out string

		switch cmd {

		case "upper":
			out = strings.ToUpper(text)

		case "lower":
			out = strings.ToLower(text)

		case "cap":
			words := strings.Fields(text)
			for i, w := range words {
				words[i] = strings.ToUpper(string(w[0])) + w[1:]
			}
			out = strings.Join(words, " ")

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
			out = strings.Join(words, " ")

		case "snake":
			var r []rune
			for _, c := range text {
				if unicode.IsLetter(c) || unicode.IsDigit(c) {
					r = append(r, unicode.ToLower(c))
				} else if unicode.IsSpace(c) {
					r = append(r, '_')
				}
			}
			out = strings.Trim(strings.ReplaceAll(string(r), "__", "_"), "_")

		case "reverse":
			words := strings.Fields(text)
			for i, w := range words {
				r := []rune(w)
				for l, h := 0, len(r)-1; l < h; l, h = l+1, h-1 {
					r[l], r[h] = r[h], r[l]
				}
				words[i] = string(r)
			}
			out = strings.Join(words, " ")

		default:
			fmt.Printf("✗ Unknown command: \"%s\"\nValid commands: upper, lower, cap, title, snake, reverse, exit\n", cmd)
			continue
		}

		fmt.Println("→", out)
	}
}
