# Go structs

A struct is basically a custom container used to bundle different data types into one record.

While an array is stuck with just one type, a struct lets me mix things like strings and integers to represent real-world objects, like a "Person" or a "Car."

To set one up, I use the "type" and "struct" keywords to define the blueprint. Once I’ve created a variable from that blueprint, I use the dot operator (.) to assign or grab values from its members. They are also super easy to pass into functions as arguments, which helps keep my code organized when I'm dealing with complex data.

I personally find structs to be the backbone of any serious Go project. They make the code way more readable seeing pers1.name is much more relatable than trying to remember which index in a slice holds a person's name. Below is an example of a struct in go.

```go
package main
import ("fmt")

type Person struct {
  name string
  age int
  job string
  salary int
}

func main() {
  var pers1 Person
  var pers2 Person

  // Pers1 specification
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 6000

  // Pers2 specification
  pers2.name = "Cecilie"
  pers2.age = 24
  pers2.job = "Marketing"
  pers2.salary = 4500

  // Access and print Pers1 info
  fmt.Println("Name: ", pers1.name)
  fmt.Println("Age: ", pers1.age)
  fmt.Println("Job: ", pers1.job)
  fmt.Println("Salary: ", pers1.salary)

  // Access and print Pers2 info
  fmt.Println("Name: ", pers2.name)
  fmt.Println("Age: ", pers2.age)
  fmt.Println("Job: ", pers2.job)
  fmt.Println("Salary: ", pers2.salary)
}
```