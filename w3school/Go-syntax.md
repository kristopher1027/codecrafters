```go
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
```

A .go file or program consists mainly of the following parts:

    Package declaration
    Import packages
    Functions
    Statements and expressions


and in Go, every program is part of a package. We define this using the package keyword. In the example i gave above, the program belongs to the main package.

import ("fmt") lets us import files included in the fmt package as shown in line 3 above.

Line 4: A blank line. Go ignores white space. Having white spaces in code makes it more readable.

Line 5: func main() {} is a function. Any code inside its curly brackets {} will be executed.

Line 6: fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!". 


## Go Statements

fmt.Println("Hello World!") is a statement.

In Go, statements are separated by ending a line (hitting the Enter key) or by a semicolon ";".

Hitting the Enter key adds ";" to the end of the line implicitly (does not show up in the source code).

The left curly bracket { cannot come at the start of a line.