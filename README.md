# go_bankholidays_uk

This package enables the easy querying of the UK goverments bank holiday feed. (https://www.gov.uk/bank-holidays.json)

## Installation

To get the package run the following
```cli
go get github.com/davidferdinand/go_bankholidays_uk
```

## Usage

This package allows you to get the bank holidays for a country in the UK in one of four ways, by calling the relevant function.
- All bank holidays for the given country from the feed. 
```go
holidays, err := go_bankholidays_uk.Find("scotland")
```
- Bank holidays for the given country before a inputed date.
```go
holidays, err := go_bankholidays_uk.FindTo("scotland", "2022-01-01")
```
- Bank holidays for the given country after a inputed date.
```go
holidays, err := go_bankholidays_uk.FindFrom("scotland", "2022-01-01")
```
- Bank holidays for the given country between two inputed dates.
```go
holidays, err := go_bankholidays_uk.FindBetween("scotland", "2020-01-01", "2022-01-01")
```

The countries are split as follows:
- `england-and-wales`
- `scotland`
- `northern-ireland`

#### Example Usage

```go
package main

import (
	"fmt"
	"github.com/davidferdinand/go_bankholidays_uk"
)

func main() {
    holidays, err := go_bankholidays_uk.FindFrom("scotland", "2022-01-01")
	if err != nil {
		panic(err)
	}
	fmt.Println(holidays)
}
```