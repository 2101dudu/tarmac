package coordinator

import (
	"strconv"
	"tarmac/utils"
)

type ProductQuery struct {
	Country   string `json:"Country"`
	Location  string `json:"Location"`
	DepDate   string `json:"DepDate"`
	Tag       string `json:"Tag"`
	SortBy    string `json:"SortBy"`
	SortOrder string `json:"SortOrder"`
	PriceFrom string `json:"PriceFrom"`
	PriceTo   string `json:"PriceTo"`
	NumDays   string `json:"NumDays"`
	Length    int    `json:"Length" binding:"omitempty,min=1"`
}

func ExtractProductQueryFromMap(raw map[string]any) ProductQuery {
	length := 24
	if l, err := strconv.Atoi(utils.ExtractStringField(raw, "Length")); err == nil && l > 0 {
		length = l
	}

	return ProductQuery{
		Country:   utils.ExtractStringField(raw, "Country"),
		Location:  utils.ExtractStringField(raw, "Location"),
		DepDate:   utils.ExtractStringField(raw, "DepDate"),
		Tag:       utils.ExtractStringField(raw, "Tag"),
		SortBy:    utils.ExtractStringField(raw, "SortBy"),
		SortOrder: utils.ExtractStringField(raw, "SortOrder"),
		PriceFrom: utils.ExtractStringField(raw, "PriceFrom"),
		PriceTo:   utils.ExtractStringField(raw, "PriceTo"),
		NumDays:   utils.ExtractStringField(raw, "NumDays"),
		Length:    length,
	}
}

type ContactInfo struct {
	Name            string `json:"Name"`
	Surname         string `json:"Surname"`
	Email           string `json:"Email"`
	Phone           string `json:"Phone"`
	DateOfBirth     string `json:"DateOfBirth"`
	WantsNewsletter string `json:"WantsNewsletter"`
}

func ExtractContactInfoFromMap(raw map[string]any) ContactInfo {
	return ContactInfo{
		Name:            utils.ExtractStringField(raw, "Name"),
		Surname:         utils.ExtractStringField(raw, "Surname"),
		Email:           utils.ExtractStringField(raw, "Email"),
		Phone:           utils.ExtractStringField(raw, "Phone"),
		DateOfBirth:     utils.ExtractStringField(raw, "DateOfBirth"),
		WantsNewsletter: utils.ExtractStringField(raw, "WantsNewsletter"),
	}
}
