package main

import "github.com/Kamva/mgm/v2"

type Result struct {
	Success bool `json:"success"`
	Results []struct {
		CountryName           string `json:"countryName"`
		ProvinceShortName     string `json:"provinceShortName"`
		UpdateTime            int    `json:"updateTime"`
		LocationId            int    `json:"locationId"`
		DeadCount             int    `json:"deadCount"`
		ConfirmedCount        int    `json:"confirmedCount"`
		SuspectedCount        int    `json:"suspectedCount"`
		CuredCount            int    `json:"curedCount"`
		CurrentConfirmedCount int    `json:"currentConfirmedCount"`
		Cities                []struct {
			CityName              string `json:"cityName"`
			LocationId            int    `json:"locationId"`
			DeadCount             int    `json:"deadCount"`
			ConfirmedCount        int    `json:"confirmedCount"`
			SuspectedCount        int    `json:"suspectedCount"`
			CuredCount            int    `json:"curedCount"`
			CurrentConfirmedCount int    `json:"currentConfirmedCount"`
		} `json:"cities"`
	} `json:"results"`
}

type NCovInfo struct {
	mgm.DefaultModel   `bson:",inline"`
	Region             string      `json:"region" bson:"region"`
	LocID              int         `json:"locID" bson:"locID"`
	Date               string      `json:"date" bson:"date"`
	Dead               int         `json:"dead" bson:"dead"`
	Confirmed          int         `json:"confirmed" bson:"confirmed"`
	Suspected          int         `json:"suspected" bson:"suspected"`
	Cured              int         `json:"cured" bson:"cured"`
	RemainingConfirmed int         `json:"remainingConfirmed" bson:"remainingConfirmed"`
	Cities             []*CityInfo `json:"cities" bson:"cities"`
}

type CityInfo struct {
	Name               string `json:"name" bson:"name"`
	LocID              int    `json:"locID" bson:"locID"`
	Dead               int    `json:"dead" bson:"dead"`
	Confirmed          int    `json:"confirmed" bson:"confirmed"`
	Suspected          int    `json:"suspected" bson:"suspected"`
	Cured              int    `json:"cured" bson:"cured"`
	RemainingConfirmed int    `json:"remainingConfirmed" bson:"remainingConfirmed"`
}

func (n *NCovInfo) CollectionName() string {
	return "ncovinfos"
}
