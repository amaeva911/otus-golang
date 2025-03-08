package main

import (
	"fmt"
	"time"
)

func main() {
	callNext()
	timeNow()
}

func callNext() {
	fmt.Println("Docker, ты следующий!!!")
}

func timeNow() {
	t := time.Now()
	fmt.Printf(t.Format(time.RFC3339))
}
