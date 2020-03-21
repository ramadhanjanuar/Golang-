package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "burger banana soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Println(res1)

	var res2 = regex.FindAllString(text, -1)
	fmt.Println(res2)

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	var str = regex.FindString(text)
	fmt.Println(str)

	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)
	// [0, 6]

	str = text[0:6]
	fmt.Println(str)

	str = regex.ReplaceAllString(text, "potato")
	fmt.Println(str)

	str = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str)
}
