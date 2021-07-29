package main

import (
	"fmt"
	"time"
)

func main() {
	layout1 := "03:04:05PM"
	layout2 := "15:04:05"
	t, err := time.Parse(layout1, "10:12:45PM")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Format(layout1))
	fmt.Println(t.Format(layout2))
}
