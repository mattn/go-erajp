package era_jp

import (
	"fmt"
	"time"
)

//go:generate go run tool/scrape.go > table.go

type Era int

func (era Era) String() string {
	s := ToEra(int(era))
	if s == "" {
		s = fmt.Sprint(int(era))
	}
	return s
}

func ToEra(y int) string {
	for i := len(eras) - 1; i >= 0; i-- {
		if eras[i].Year <= y {
			return eras[i].Name
		}
	}
	return ""
}

func ToEraFromTime(t time.Time) string {
	loc, _ := time.LoadLocation("JST")
	if loc == nil {
		loc = time.Local
	}
	for i := len(eras) - 1; i >= 0; i-- {
		et := time.Date(eras[i].Year, time.Month(eras[i].Month), eras[i].Day, 0, 0, 0, 0, loc)
		if t.After(et) {
			return eras[i].Name
		}
	}
	return ""
}
