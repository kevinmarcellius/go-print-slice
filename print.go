package sliceprinter

import (
	"fmt"
)

func PrintSlice(slice []int) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

func Print2DSlice(slice [][]int) {
	for i := range slice {
		for j := range slice[i] {
			fmt.Print(slice[i][j], " ")
		}
		fmt.Println()
	}
}
