## Go-Conditions
 Conditions are used to make decisions. I rely on standard math symbols like ==, !=, <, and > to check if something is true or false, and I combine them with && (AND) or || (OR) for more complex logic.

When I need to run code only if a specific rule is met, I use an if statement. If that rule fails, I use else as a fallback. When I have multiple "what-if" scenarios, I chain them with else if, but I keep in mind that Go will only execute the first true condition it hits and skip the rest.

I learnt to remember to keep the syntax tight—Go is picky and requires the else to be on the same line as the closing bracket (} else {).