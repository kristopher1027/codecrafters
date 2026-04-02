 ## Go Variables
Variables are labeled jars for storing data. Go is "statically typed," meaning it cares about what kind of data (integers, strings, etc.) you put in them. Below are the Basic data types in Go.

> int- stores integers (whole numbers), such as 123 or -123

> float32- stores floating point numbers, with decimals, such as 19.99 or -19.99

> string - stores text, such as "Hello World". String values are surrounded by double quotes

> bool- stores values with two states: true or false

### Two ways to create them:

    var: Can be used anywhere (even outside functions). You can declare it now and give it a value later.
    :=: The "short-cut." Use this inside functions for quick declarations. Go will guess the type for you based on the value.

## Key Rules:

    Zero Values: If you don't give a variable a value, Go doesn't leave it empty; it initializes it to its default value (like 0 for numbers or "" for strings).

    Scope: You cannot use the := shortcut outside of a function—you must use var there.

    Types: Once a variable is a certain type (like a string), you can't suddenly shove a number into it.

## multiple variables can also be delclared like so
```go
package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
```

## For variable naming 
Names must start with a
letter or underscore, never a digit.
Use only letters, numbers, and underscores; no spaces allowed.

Avoid Go keywords and remember names are case-sensitive (age vs Age).

Camel Case (myVar) is the Go standard for internal variables.

Pascal Case (MyVar) is used when you want a variable to be "exported" (public).

Snake Case (my_var) is allowed but less common in the Go community.

There’s no length limit, but keep names descriptive yet concise.