package main

import "fmt"

func main() {
	var point = 9

	if point == 10 {
		fmt.Println("Wah, kamu hebat")
	} else if point > 7 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Dungu")
	}

	//variable temporary
	var poin = 8840.0
	if percent := poin / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	//switch case
	point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("good")
	default:
		fmt.Println("dungu!!")
	}

	//Switch case dengan banyak kondisi
	point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("good")
	default:
		fmt.Println("dungu!!")
	}

	//Switch dengan gaya if-else
	point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("good")
	default:
		{
			fmt.Println("dungu!!")
			fmt.Println("Belajar ya!!")
		}
	}

	//Penggunaan keyword fallthrough
	fmt.Println("")
	point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("good")
		fallthrough
	default:
		{
			fmt.Println("dungu!!")
			fmt.Println("Belajar ya!!")
		}
	}

	//nested seleksi kondisi
	fmt.Println("")
	point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder")
			}
		}
	}
}
