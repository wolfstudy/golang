package main

import (
	"time"
	"fmt"
)

func main() {

	//get now timestamp
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	//format string
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05"))

	//format timestamp
	tm2, _ := time.Parse("01/02/2006 03:04:05", "01/02/2018 12:00:00")
	fmt.Println(tm2.Unix())

	//1514678400---->"12/31/2017 00:00:00"
	//1514894400---->"01/02/2018 12:00:00"
}