package main

import "fmt"

type medias struct {
	Name  string
	Value string
	Type  string
}

func main() {
	data := []medias{
		{
			Name:  "object A",
			Value: "value A",
			Type:  "A",
		},
		{
			Name:  "object B",
			Value: "value B",
			Type:  "B",
		},
		{
			Name:  "object C",
			Value: "value C",
			Type:  "C",
		},
		{
			Name:  "object A2",
			Value: "value A2",
			Type:  "A",
		},
		{
			Name:  "object B2",
			Value: "value B2",
			Type:  "B",
		},
		{
			Name:  "object A3",
			Value: "value A3",
			Type:  "A",
		},
	}

	var mapData map[string][]medias
	mapData = map[string][]medias{}
	for _, item := range data {
		tempData := mapData[item.Type]
		mapData[item.Type] = append(tempData, item)
	}
	fmt.Println(mapData)
}
