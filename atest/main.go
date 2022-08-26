package main

import (
	"fmt"
)

type Chessboard map[string]File
type File []bool

func main() {
	type File []bool
	cb := (map[string]bool{"A": false, "B": false, "C": true, "D": false, "E": true, "F": false, "G": true})
	count := 0
	for _, file := range cb {

		count += len(file)
	}
	fmt.Println(count)
}

// // AddItem adds an item to customer bill.
// func AddItem(bill, units map[string]int, item, unit string) bool {

// }
