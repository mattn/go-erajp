package era_jp

//go:generate go run tool/scrape.go > table.go

func ToEra(y int) string {
	for i := len(eras) - 1; i >= 0; i-- {
		if eras[i].Year <= y {
			return eras[i].Name
		}
	}
	return ""
}
