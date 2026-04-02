## Go output functions

* Print(): The basic one. It dumps text as-is unless you manually add \n for new lines or " " for spaces between strings.

* Println(): The "standard" choice. It automatically adds a space between arguments and a new line at the end.

* Printf(): The "formatter." Use this when you want to inject variables into a specific string template.

## Go Formatting Verbs
* Common Verbs: Use %v for the value, %T for the data type, and %d for integers.

* Strings & Floats: Use %s for plain text, %q for quoted text, and %.2f to limit decimals to two places.

* Alignment: You can pad outputs with spaces (e.g., %8s) or zeros (%04d) to keep things lined up.

* Base Conversion: It can easily show numbers in binary (%b), hex (%x), or octal (%o).


