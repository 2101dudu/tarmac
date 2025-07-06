package pdf

import (
	"fmt"
	"os"
	"os/exec"
)

type PDFData struct {
	FlightDesc  string
	FlightPrice float32
	HotelDesc   string
	HotelPrice  float32
}

func NewPDF(flightDesc, hotelDesc string, flightPrice, hotelPrice float32) *PDFData {
	return &PDFData{
		FlightDesc:  flightDesc,
		FlightPrice: flightPrice,
		HotelDesc:   hotelDesc,
		HotelPrice:  hotelPrice,
	}
}

func (p *PDFData) fillPDF(f *os.File) {
	inputStr :=
		`
	#let getFlightDesc() = {
	  "Voo de Ida: %s"
	}

	#let getFlightPrice() = {
		"%.2f€"
	}

	#let getHotelDesc() = {
		"Hotel %s 5 estrelas"
	}

	#let getHotelPrice() = {
		"%.2f€"
	}
	`

	formatString := fmt.Sprintf(inputStr,
		p.FlightDesc,
		p.FlightPrice,
		p.HotelDesc,
		p.HotelPrice)

	_, err := f.WriteString(formatString)
	if err != nil {
		panic(err)
	}
}

func (p *PDFData) GeneratePDF() {
	filePath := "input.typ"

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer os.Remove(filePath)

	p.fillPDF(f)
}

func (p *PDFData) GetPDFFile() *os.File {
	err := exec.Command("typst", "compile", "template.typ").Run()
	if err != nil {
		panic(err)
	}

	f, err := os.Open("template.typ")
	if err != nil {
		panic(err)
	}
	return f
}
