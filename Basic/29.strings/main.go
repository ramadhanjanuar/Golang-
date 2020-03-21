package main

import (
	"fmt"
	"strings"
)

func main() {
	var text = "John Wick"
	var isExist = strings.Contains(text, "Wick")
	fmt.Println(isExist)

	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true

	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2

	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2

	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2) // 4

	text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	var str = strings.Repeat("na", 4)
	fmt.Println(str) // "nananana"

	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1) // ["the", "dark", "knight"]

	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // ["b", "a", "t", "m", "a", "n"]

	var data = []string{"banana", "papaya", "tomato"}
	str = strings.Join(data, "-")
	fmt.Println(str) // "banana-papaya-tomato"

	str = strings.ToLower("aLAy")
	fmt.Println(str) // "alay"

	str = strings.ToUpper("eat!")
	fmt.Println(str) // "EAT!"
}
