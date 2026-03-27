package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convtodec(hexstr string, base int) (int64, error) {
	return strconv.ParseInt(hexstr, base, 64)
}

func main() {

	fmt.Println("which number did you want to convert between bases")
	for {
		var input1 string
		fmt.Println("Enter Value:")
		fmt.Scan(&input1)

		var operator string
		fmt.Println("operation: dec  bin  hex  other")
		fmt.Scan(&operator)

		switch operator {
		case "bin", "Bin":
			fmt.Println(convtodec(input1, 2))
		case "hex", "Hex":
			fmt.Println(convtodec(input1, 16))
			return
		case "dec", "Dec":
			var num int64
			fmt.Scan(&num)
			fmt.Println(strings.ToUpper(strconv.FormatInt(num, 16)))
			fmt.Println(strconv.FormatInt(num, 2))
			// str is "11111111"

			// str is "ff"

			// fmt.Println(convtodec("1E", 16))
			// fmt.Println(convtodec("10", 2))
			// fmt.Println(convtodec("255", 10))
		}
	}
}
