package main

import "fmt"

func main() {
	s := 8

	switch s {
	case 10:
		fmt.Println("case 10")
		fallthrough //继续执行下一个分支
	case 5:
		fmt.Println("case 5")
	default:
		fmt.Println("default")
	}

LABEL1:
	for i := 0; i <= 5; i++ {
		//fmt.Println("I:", i)
		for j := 0; j <= 5; j++ {
			//fmt.Println("J:", j)
			if j == 4 {
				// i继续下一个值，j会被重置
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
