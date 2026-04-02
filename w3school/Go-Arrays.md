## Go Arrays

I look at Go arrays as rigid containers. When I declare one, I’m telling the computer exactly how many slots I need, which is something i can’t change that later. If I create an array for 5 items, it stays 5 items; to add a 6th, I’d have to build an entirely new array.

## on array sizes
I can define the size manually, like [3]int, or let Go count my initial values for me using [...]int. Either way, once it's set, it's final. If I leave some slots empty, Go doesn't leave them "null" but fills them with "zero values" like 0 for numbers or an empty string for text.

## modification
To grab or change my data, I use an index starting at 0. I can even skip around and only fill specific spots, like saying "put 50 at index 2," and Go will leave the rest as defaults. Because they are so fixed, I mainly use arrays as the foundation for Slices, which are the flexible version I'll actually use most of the time.

