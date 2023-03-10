package main

import "fmt"

func main() {
	input:=5
	for i := 0; i < input; i++ {
		fmt.Println("Nilai i :", i)
	}
	for j := 0; j <= 10; j++ {
		if j == 5 {
			strings := "C A ШA P B O"
			for index, v := range strings {
				if index%2 == 0 || index == 0 {
					fmt.Printf("character %#U starts at byte position %d\n", v, index)
				}
			}
		}
		fmt.Println("Nilai j :", j)
	}
}