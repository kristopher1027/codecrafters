## Maps
are unordered collections that store data in key:value pairs. 
They are highly efficient for data retrieval because they are built on top of hash tables. Unlike arrays or slices, the order of elements in a map is not guaranteed and can change every time the code runs.

Creating a map is best done using the make() function. While it is possible to declare a map variable without it, doing so creates a nil map; attempting to add data to a nil map causes a runtime panic. Map keys must be types that support the equality operator (==), such as strings or integers, while values can be any type.

>Key Operations:
    Access/Update: Use metadata["key"] = value to add or change data.
    Delete: Use the delete(map, "key") function to remove an entry.
    Existence Check: Use the "comma ok" syntax (val, ok := m["key"]) to distinguish between a key that exists with a zero value and a key that is actually missing.
    Iteration: Use a for range loop to grab keys and values, but remember the sequence will be random.

Because maps are reference types, assigning one map variable to another does not create a copy. Both variables will point to the same underlying data, so changing one affects the other. To maintain a specific order during iteration, a separate slice must be used to store and sort the keys manually.

This example below shows how to create maps in Go. Notice the order in the code and in the output.

```go

package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}
```