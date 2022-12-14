package main

import (
	"fmt"
	"time"
)

var Version = "development"
var BuildDate = time.Now().Format("Mon Jan 2 15:04:05")

func main() {
	fmt.Println("hello moto!")
}
