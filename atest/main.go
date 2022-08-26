package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(Schedule("01/02/2006 15:04:05"))

}

func Schedule(date string) time.Time {
	//layout := "2006-01-02 15:04:05"
	t, err := time.Parse("01/02/2006 15:04:05", date)
	if err != nil {
		fmt.Println("error", err)
	}
	return t
}
