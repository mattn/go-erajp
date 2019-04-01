package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	out, err := os.Create("table.go")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	resp, err := http.Get("http://www.kumamotokokufu-h.ed.jp/kumamoto/bungaku/nengoui.html")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	r := transform.NewReader(resp.Body, japanese.EUCJP.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(out, `package erajp

var eras = []struct {
	Name  string
	Ruby  string
	Year  int
	Month int
	Day   int
}{
`)

	hasReiwa := false

	var prevy, prevm, prevd string
	doc.Find(`table[border="1"] tr`).Each(func(n int, sel *goquery.Selection) {
		cells := sel.Children().Filter("td")
		tdnum := cells.Length()
		if _, ok := cells.Eq(1).Attr("rowspan"); ok {
			return
		}

		var name, ruby, fromto, day string
		if tdnum == 6 {
			name = cells.Eq(1).Text()
			ruby = cells.Eq(2).Text()
			fromto = cells.Eq(3).Text()
			day = cells.Eq(4).Text()
		} else if tdnum == 7 {
			name = cells.Eq(2).Text()
			ruby = cells.Eq(3).Text()
			fromto = cells.Eq(4).Text()
			day = cells.Eq(5).Text()
		} else {
			return
		}
		y := strings.Split(fromto, "～")
		day = strings.Replace(day, "閏", "", -1)
		md := strings.Split(day, "/")
		if md[1] == "？" {
			md[1] = "1"
		}

		if y[0] == prevy && md[0] == prevm && md[1] == prevd {
			return
		}
		fmt.Fprintf(out, "\t{Name: %q, Ruby: %q, Year: %v, Month: %v, Day: %v},\n", name, ruby, y[0], md[0], md[1])
		prevy = y[0]
		prevm = md[0]
		prevd = md[1]

		if !hasReiwa && name == "令和" {
			hasReiwa = true
		}
	})
	if !hasReiwa {
		fmt.Fprintf(out, "\t{Name: %q, Ruby: %q, Year: %v, Month: %v, Day: %v},\n", "令和", "れいわ", 2019, 5, 1)
	}
	fmt.Fprintln(out, "}")
}
