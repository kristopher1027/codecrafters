Constants are the go to if a variable should have a fixed value that cannot be changed, you can use the const keyword.

The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
Syntax
```go
const CONSTNAME type = value
```
Note: The value of a constant must be assigned when you declare it.

Constant names follow the same naming rules as variables
Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
Constants can be declared both inside and outside of a function

There are two types of constants:

    Typed constants
    Untyped constants

Typed constants are declared with a defined type while undefined constants are not

```go
//typed constant
package main
import ("fmt")

const A int = 1

func main() {
  fmt.Println(A)
}
```

```go
//typed constant
package main
import ("fmt")

const A = 1

func main() {
  fmt.Println(A)

}
```