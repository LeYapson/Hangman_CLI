package main

import "fmt"

func PrintHang(n int) {
	hang := ImportHang()

	if n == 9 {
		for i := 1; i <= 7; i++ {
			fmt.Println(hang[i])
		}
	} else if n == 8 {
		for i := 7; i < 15; i++ {
			fmt.Println(hang[i])
		}
	} else if n == 7 {
		for i := 15; i < 23; i++ {
			fmt.Println(hang[i])
		}

	} else if n == 6 {
		for i := 23; i < 31; i++ {
			fmt.Println(hang[i])
		}
	} else if n == 5 {
		for i := 31; i < 39; i++ {
			fmt.Println(hang[i])
		}
	} else if n == 4 {
		for i := 39; i < 47; i++ {
			fmt.Println(hang[i])
		}
	} else if n == 3 {
		for i := 47; i < 55; i++ {
			fmt.Println(hang[i])
		}
	} else if n == 2 {
		for i := 55; i < 63; i++ {
			fmt.Println(hang[i])
		}
	} else if n == 1 {
		for i := 63; i < 71; i++ {
			fmt.Println(hang[i])
		}
	} else if n == 0 {
		for i := 71; i < 79; i++ {
			fmt.Println(hang[i])
		}
	}

}
