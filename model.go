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

type NCovOverallInfo struct {
	mgm.DefaultModel       `bson:",inline"`
	ModifyTime             int64 `json:"modifyTime" bson:"-"`
	Time                   int32 `bson:"time"`
	Dead                   int   `json:"deadCount" bson:"dead"`
	DeadIncr               int   `json:"deadIncr" bson:"deadIncr"`
	Confirmed              int   `json:"confirmedCount" bson:"confirmed"`
	ConfirmedIncr          int   `json:"confirmedIncr" bson:"confirmedIncr"`
	Suspected              int   `json:"suspectedCount" bson:"suspected"`
	SuspectedIncr          int   `json:"suspectedIncr" bson:"suspectedIncr"`
	Cured                  int   `json:"curedCount" bson:"cured"`
	CuredIncr              int   `json:"curedIncr" bson:"curedIncr"`
	RemainingConfirmed     int   `json:"currentConfirmedCount" bson:"remainingConfirmed"`
	RemainingConfirmedIncr int   `json:"currentConfirmedIncr" bson:"remainingConfirmedIncr"`
	Serious                int   `json:"seriousCount" bson:"serious"`
	SeriousIncr            int   `json:"seriousIncr" bson:"seriousIncr"`
	GlobalStatistics       struct {
		Dead                   int `json:"deadCount" bson:"dead"`
		DeadIncr               int `json:"deadIncr" bson:"deadIncr"`
		Confirmed              int `json:"confirmedCount" bson:"confirmed"`
		ConfirmedIncr          int `json:"confirmedIncr" bson:"confirmedIncr"`
		Cured                  int `json:"curedCount" bson:"cured"`
		CuredIncr              int `json:"curedIncr" bson:"curedIncr"`
		RemainingConfirmed     int `json:"currentConfirmedCount" bson:"remainingConfirmed"`
		RemainingConfirmedIncr int `json:"currentConfirmedIncr" bson:"remainingConfirmedIncr"`
	} `json:"globalStatistics" bson:"global"`
}

func (n *NCovInfo) CollectionName() string {
	return "ncovinfos"
}

func (n *NCovOverallInfo) CollectionName() string {
	return "ncov_overall_infos"
}
