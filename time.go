package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println("DATE : ", t.Format("12-02-2022"))
	fmt.Println("Time : ", t.Format("10:04:10"))
}
