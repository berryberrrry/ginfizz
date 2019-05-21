package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1

	go func() {
		fmt.Println(a)
	}()
	time.Sleep(1 * time.Second)

}
