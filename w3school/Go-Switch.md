## Go-Switch

I use the switch statement when I have a long list of specific values to check, as it’s much cleaner than writing a dozen else if blocks. I just provide an expression once, and Go compares it against each case until it finds a match.

Go's switch is said to be smarter than in other languages as it only runs the matched case and stops, so one doesn't have to manually type break every single time. If nothing matches, then the default keyword acts as a catch-all to handle unexpected values.

I also find it really handy that I can group multiple values into a single case using commas, like case 1, 3, 5:.. This lets me run the same code for a whole set of possibilities without repeating myself. I just have to make sure my case values are the same data type as my main expression, or the compiler will report an error.

