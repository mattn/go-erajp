package erajp

import (
	"testing"
	"time"
)

var testcases = []struct {
	year int
	era  string
}{
	{1900, `明治`},
	{1911, `明治`},
	{1912, `大正`},
	{1925, `大正`},
	{1926, `昭和`},
	{1988, `昭和`},
	{1989, `平成`},
	{2016, `平成`},
	{2019, `令和`},
	{2039, `令和`},
}

func TestEra(t *testing.T) {
	year := 2030
	got := Era(year).String()
	expected := "令和"
	if got != expected {
		t.Fatalf("Expected %v for year %v, but %v:", expected, year, got)
	}
	year = 640
	got = Era(year).String()
	expected = "640"
	if got != expected {
		t.Fatalf("Expected %v for year %v, but %v:", expected, year, got)
	}
}

func TestToEra(t *testing.T) {
	for _, testcase := range testcases {
		year := testcase.year
		got := ToEra(year)
		expected := testcase.era
		if got != expected {
			t.Fatalf("Expected %v for year %v, but %v:", expected, year, got)
		}
	}
}

func TestNow(t *testing.T) {
	now := time.Now()
	got := ToEraFromTime(now)
	expected := "令和"
	if got != expected {
		t.Fatalf("Expected %v for year %v, but %v:", expected, now, got)
	}
}

func TestFind(t *testing.T) {
	now := time.Now()
	got := Find(now)
	if got == nil {
		t.Fatalf("Expected non-nil")
	}
	expected := "令和"
	if got.Name != expected {
		t.Fatalf("Expected %s for year %v, but %v:", expected, now, got.Name)
	}
}
