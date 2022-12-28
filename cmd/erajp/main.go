package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mattn/go-erajp"
)

var formats = []struct {
	f string
	t int
}{
	{"2006", 0},
	{"2006/1", 1},
	{"2006/01", 1},
	{"2006/1/2", 2},
	{"2006/01/02", 2},
}

func main() {
	if len(os.Args) == 1 {
		ymd := time.Now()
		if item := erajp.Find(ymd); item != nil {
			y := fmt.Sprint(ymd.Year() - item.Year + 1)
			if y == "1" {
				y = "元"
			}
			fmt.Printf("%s%s年\n", item.Name, y)
			return
		}
	} else if len(os.Args) == 2 {
		for _, format := range formats {
			ymd, err := time.ParseInLocation(format.f, os.Args[1], time.Local)
			if err == nil {
				switch format.t {
				case 0:
					ymd = time.Date(ymd.Year()+1, 1, 1, 0, 0, 0, 0, time.Local).Add(-1)
				case 1:
					ymd = time.Date(ymd.Year(), ymd.Month()+1, 1, 0, 0, 0, 0, time.Local).Add(-1)
				case 2:
					ymd = time.Date(ymd.Year(), ymd.Month(), ymd.Day()+1, 0, 0, 0, 0, time.Local).Add(-1)
				}
				if item := erajp.Find(ymd); item != nil {
					y := fmt.Sprint(ymd.Year() - item.Year + 1)
					if y == "1" {
						y = "元"
					}
					fmt.Printf("%s%s年\n", item.Name, y)
					return
				}
			}
		}
	}
	fmt.Fprintf(os.Stderr, "%v [year or year/month or year/month/day]\n", os.Args[0])
	os.Exit(2)
}
