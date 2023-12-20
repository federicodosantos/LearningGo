package main

import (
	"fmt"
	"strings"
)

/*
Simpelnya, kita dapat menjadikan sebuah fungsi menjadi parameter pada sebuah fungsi
*/

func main() {
	data := []string{"Riko", "santos", "robert"}

	var dataContainsO = filter(data, func(s string) bool {
		return strings.Contains(s, "o")
	})

	var datalenght4 = filter(data, func(s string) bool {
		return len(s) == 4
	})

	fmt.Println("Data: ", data)
	fmt.Println("Data contains o: ", dataContainsO)
	fmt.Println("Data lenght = 4 : ", datalenght4)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
