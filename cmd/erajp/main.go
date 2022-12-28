package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mattn/go-erajp"
)

var formats = []string{
	"2006",
	"2006/1",
	"2006/01",
	"2006/1/2",
	"2006/01/02",
}

func main() {
	if len(os.Args) == 2 {
		for _, format := range formats {
			ymd, err := time.Parse(format, os.Args[1])
			if err == nil {
				fmt.Println(erajp.ToEraFromTime(ymd))
				return
			}
		}
	}
	fmt.Fprintf(os.Stderr, "%v [year or year/month or year/month/day]\n", os.Args[0])
	os.Exit(2)
}
