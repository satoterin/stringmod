package main

import (
	"fmt"

	str "github.com/satoterin/stringmod/quotes"
	qt "github.com/satoterin/stringmod/strings"
)

func main() {
	o, e := str.CountOddEven("12345")
	fmt.Println(o, e)

	fmt.Println(qt.GetEmoji("turtle"))
}
