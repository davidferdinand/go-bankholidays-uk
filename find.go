package go_bankholidays_uk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BankHolidays struct {
	EnglandAndWales struct {
		Division string   `json:"division"`
		Events   []Events `json:"events"`
	} `json:"england-and-wales"`
	Scotland struct {
		Division string   `json:"division"`
		Events   []Events `json:"events"`
	} `json:"scotland"`
	NorthernIreland struct {
		Division string   `json:"division"`
		Events   []Events `json:"events"`
	} `json:"northern-ireland"`
}

type Region struct {
	Division string
	Events   []Events
}

type Events struct {
	Title string `json:"title"`
	Date  string `json:"date"`
	// Notes string `json:"notes"`
}

var layout = "2006-01-02"

func getRegions(b *BankHolidays) []Region {
	regions := make([]Region, 0)
	regions = append(regions, Region{
		Division: "england-and-wales",
		Events:   b.EnglandAndWales.Events,
	})

	regions = append(regions, Region{
		Division: "scotland",
		Events:   b.Scotland.Events,
	})

	regions = append(regions, Region{
		Division: "northern-ireland",
		Events:   b.NorthernIreland.Events,
	})

	return regions
}

func Find(division string) []Events {
	allBankHolidays, err := getBankHolidays()

	fmt.Println(err)

	regions := getRegions(allBankHolidays)

	requestedRegion := getEventsByRegion(regions, division)

	return requestedRegion
}

func getBankHolidays() (*BankHolidays, error) {
	res, err := http.Get("https://www.gov.uk/bank-holidays.json")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var a = new(BankHolidays)

	err = json.Unmarshal(body, &a)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return a, err
}

func FindFrom(division string, from string) ([]Events, error) {
	requestedRegion := Find(division)

	var fromEvents []Events

	fromTime, err := time.Parse(layout, from)
	if err != nil {
		return nil, err
	}

	for _, r := range requestedRegion {
		eventTime, err := time.Parse(layout, r.Date)
		if err != nil {
			return nil, err
		}
		if eventTime.After(fromTime) {
			fromEvents = append(fromEvents, r)
		}
	}
	return fromEvents, err
}

func FindTo(division string, to string) ([]Events, error) {
	requestedRegion := Find(division)

	var toEvents []Events

	toTime, err := time.Parse(layout, to)
	if err != nil {
		return nil, err
	}

	for _, r := range requestedRegion {
		eventTime, err := time.Parse(layout, r.Date)
		if err != nil {
			return nil, err
		}
		if eventTime.Before(toTime) {
			toEvents = append(toEvents, r)
		}
	}
	return toEvents, err
}

func FindBetween(division string, from string, to string) ([]Events, error) {
	requestedRegion := Find(division)

	var betweenEvents []Events

	toTime, err := time.Parse(layout, to)
	if err != nil {
		return nil, err
	}

	fromTime, err := time.Parse(layout, from)
	if err != nil {
		return nil, err
	}

	for _, r := range requestedRegion {
		eventTime, err := time.Parse(layout, r.Date)
		if err != nil {
			return nil, err
		}
		if eventTime.After(fromTime) && eventTime.Before(toTime) {
			betweenEvents = append(betweenEvents, r)
		}
	}
	return betweenEvents, err
}

func getEventsByRegion(regions []Region, name string) []Events {
	for _, r := range regions {
		if r.Division == name {
			return r.Events
		}
	}
	return nil
}
