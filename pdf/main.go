package main

import "fmt"

func main() {
	s := Service{
		Icon:   "bed",
		Desc:   "isto é uma descrição",
		Qty:    12,
		Status: "OK",
		From:   "2025-06-08",
		To:     "2025-06-12",
	}
	p := Price{

		Desc:       "Preço base",
		Qty:        2,
		UnitGross:  708.12,
		TotalGross: 709.24,
	}
	pol := Policy{
		Kind:    "CANCEL",
		Service: "AVIÃO",
		From:    "2025-06-08",
		To:      "2025-06-12",
		Price:   "0 EUR",
	}
	pdf := PDFData{
		GeneralInfo: GeneralInfo{
			QuotationNumber: 123345,
			CustomerName:    "dudu",
			ProdName:        "mykonos",
			DateIn:          "2025-06-08",
			DateOut:         "2025-06-12",
		},
		Services:     []Service{s},
		Prices:       []Price{p},
		TotalPrice:   709.24,
		Policies:     []Policy{pol},
		Conditions:   "BIG ASS TExTO SOBRE Conditions",
		Description:  "BIG ASS TExTO SOBRE Description",
		Program:      "BIG ASS TExTO SOBRE Program",
		Observations: "BIG ASS TExTO SOBRE Observations",
	}

	err := pdf.GeneratePDF()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("PDF generated")
}
