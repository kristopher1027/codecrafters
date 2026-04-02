## Go-Functions

Functions are blocks of code that only run when called. They are created with the func keyword, a name, and parentheses. Names must start with a letter and follow the same case-sensitive rules as variables.

### function parameters
To pass data in, parameters are defined with types (like x int). When the function is called, the actual values sent are called arguments. Unlike many languages, Go functions can return multiple values, and return variables can even be named in the function signature for a "naked" return.

If a return value isn't needed, the underscore (_) is used to discard it. Go also supports recursion, where a function calls itself. This is elegant for math problems like factorials, but it requires a solid "stop condition" to avoid crashing the program or eating up all the memory.

### return
I find the multiple return values a total lifesaver—it’s so much easier than creating a whole struct just to get two pieces of data back. Recursion is cool too, but I usually double-check my logic twice so I don't accidentally create an infinite loop.

