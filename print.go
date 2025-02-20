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
	for _, i := range slice {
		for _, j := range i {
			fmt.Println(j)
		}
	}
}
