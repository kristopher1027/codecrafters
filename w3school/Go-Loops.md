## Go-Loops

Go only has the for loop. It usually takes three parts: a counter start, a condition to check, and an update step (like i++). If the middle condition stays true, the code keeps running.

Inside the loop, continue skips the current turn and jumps to the next one, while break kills the loop entirely. Loops can also be put inside other loops, which is called nesting.

For collections like arrays or slices, the range keyword is the easiest way to go because gives back the index and the value for every item and the index or value isn't needed, a simple underscore (_) can be used to ignore it. example below


```go 

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}
```

