package go_bankholidays_uk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Title   string `json:"title"`
	Date    string `json:"date"`
	Notes   string `json:"notes"`
	Bunting bool   `json:"bunting"`
}

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
	res, err := http.Get("https://www.gov.uk/bank-holidays.json")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	a, err := getBankHolidays([]byte(body))

	regions := getRegions(a)

	requestedRegion := getEventsByRegion(regions, division)

	return requestedRegion
}

func getEventsByRegion(regions []Region, name string) []Events {
	for _, r := range regions {
		if r.Division == name {
			return r.Events
		}
	}
	return nil
}

func getBankHolidays(body []byte) (*BankHolidays, error) {
	var a = new(BankHolidays)
	err := json.Unmarshal(body, &a)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return a, err
}
