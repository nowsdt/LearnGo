package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Point struct {
		X, Y int
	}

	type Circle struct {
		point  Point
		radius int
	}

	var circle Circle
	circle.point.X = 10
	circle.point.Y = 15
	circle.radius = 5

	fmt.Println(circle)

	data, err := json.Marshal(circle.point)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("json1: %s\n", data)

	var dataMap = map[string]int{
		"tom":  15,
		"lucy": 20,
	}

	marshal, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", marshal)
}
