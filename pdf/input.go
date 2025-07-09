package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type GeneralInfo struct {
	QuotationNumber int
	CustomerName    string
	ProdName        string
	DateIn          string
	DateOut         string
}

type Service struct {
	Icon   string
	Desc   string
	Qty    int
	Status string
	From   string
	To     string
}

type Price struct {
	Desc       string
	Qty        int
	UnitGross  float32
	TotalGross float32
}

type Policy struct {
	Kind    string
	Service string
	From    string
	To      string
	Price   string
}

type PDFData struct {
	GeneralInfo GeneralInfo
	Services    []Service
	Prices      []Price
	TotalPrice  float32
	Policies    []Policy

	Conditions   string
	Description  string
	Program      string
	Observations string
}

func (p *PDFData) fillPDF(f *os.File) {
	input := bufio.NewWriter(f)

	input.WriteString(`#import "utils.typ": *`)
	input.WriteString(p.fillGeneralInfo())

	input.WriteString(p.fillAllServices())
	input.WriteString(p.fillAllPrices())
	input.WriteString(p.fillAllPolicies())

	input.WriteString(p.fillConditions())
	input.WriteString(p.fillDescription())
	input.WriteString(p.fillProgram())
	input.WriteString(p.fillObservations())

	err := input.Flush()
	if err != nil {
		panic(err)
	}
}

func (p *PDFData) fillGeneralInfo() string {
	s := `
#let quotationNumber = {
  "%d"
}

#let customerName = {
  "%s"
}

#let prodName = {
  "%s"
}

#let today = {
  "%s"
}

#let dateIn = {
  "%s"
}

#let dateOut = {
  "%s"
}
	`
	return fmt.Sprintf(s,
		p.GeneralInfo.QuotationNumber,
		p.GeneralInfo.CustomerName,
		p.GeneralInfo.ProdName,
		time.Now().Format(time.DateTime),
		p.GeneralInfo.DateIn,
		p.GeneralInfo.DateOut,
	)
}

func (p *PDFData) fillAllServices() string {
	s := `
#let allServices = {
`
	for _, entry := range p.Services {
		newS := `("%s", "%s", "%d", "%s", "%s", "%s")`
		newS = fmt.Sprintf(newS, entry.Icon, entry.Desc, entry.Qty, entry.Status, entry.From, entry.To)

		s += "\n" + newS
	}
	s += `
}
`
	return s
}

func (p *PDFData) fillAllPrices() string {
	s := `
#let allPrices = {
`
	for _, entry := range p.Prices {
		newS := `("%s", "%d", "%.2f", "%.2f")`
		newS = fmt.Sprintf(newS, entry.Desc, entry.Qty, entry.UnitGross, entry.TotalGross)

		s += "\n" + newS
	}
	s += `
}


#let totalPrice = {
  "%.2f"
}
`
	return fmt.Sprintf(s, p.TotalPrice)
}

func (p *PDFData) fillAllPolicies() string {
	s := `
#let allPolicies = {
`
	for _, entry := range p.Policies {
		newS := `("%s", "%s", "%s", "%s", "%s")`
		newS = fmt.Sprintf(newS, entry.Kind, entry.Service, entry.From, entry.To, entry.Price)

		s += "\n" + newS
	}
	s += `
}
`
	return s
}

func (p *PDFData) fillConditions() string {
	s := "#let conditions = { ```\n%s\n``` }\n"
	return fmt.Sprintf(s, p.Conditions)
}

func (p *PDFData) fillDescription() string {
	s := "#let description = { ```\n%s\n``` }\n"
	return fmt.Sprintf(s, p.Description)
}

func (p *PDFData) fillProgram() string {
	s := "#let program = { ```\n%s\n``` }\n"
	return fmt.Sprintf(s, p.Program)
}

func (p *PDFData) fillObservations() string {
	s := "#let observations = { ```\n%s\n``` }\n"
	return fmt.Sprintf(s, p.Observations)
}
