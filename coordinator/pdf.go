package coordinator

import (
	"tarmac/pdf"
	"tarmac/wsdl"
)

func fillPDF(simul *wsdl.DynGetSimulationResponse, name, surname string, quotationNumber int) pdf.PDFData {
	return pdf.PDFData{
		GeneralInfo: pdf.GeneralInfo{
			QuotationNumber: quotationNumber,
			CustomerName:    name + " " + surname,
			ProdName:        "mykonos",
			DateIn:          "2025-06-08",
			DateOut:         "2025-06-12",
		},
		Services:     fillServices(simul.Services.Items),
		Prices:       fillPrices(simul.Calcs.Items),
		TotalPrice:   extractPointer(simul.TotalPrice),
		Policies:     fillPolicies(simul.Policies.Items),
		Conditions:   extractPointer(simul.Conditions),
		Description:  extractPointer(simul.Conditions),
		Program:      "BIG ASS TExTO SOBRE Program",
		Observations: "BIG ASS TExTO SOBRE Observations",
	}
}

func fillServices(simul []*wsdl.DynResServices) []pdf.Service {
	var res []pdf.Service
	for _, s := range simul {
		if s == nil {
			continue
		}
		p := pdf.Service{
			Icon:   extractPointer(s.Type),
			Desc:   extractPointer(s.Description),
			Qty:    extractPointer(s.Quant),
			Status: extractPointer(s.Status),
			From:   extractPointer(s.DateFrom),
			To:     extractPointer(s.DateTo),
		}
		res = append(res, p)
	}
	return res
}

func fillPrices(simul []*wsdl.DynResCalcs) []pdf.Price {
	var res []pdf.Price
	for _, s := range simul {
		if s == nil {
			continue
		}
		p := pdf.Price{
			Desc:       extractPointer(s.Description),
			Qty:        extractPointer(s.Quant),
			UnitGross:  extractPointer(s.GrossUnitVal),
			TotalGross: extractPointer(s.GrossTotalVal),
		}
		res = append(res, p)
	}
	return res
}

func fillPolicies(simul []*wsdl.DynPolicies) []pdf.Policy {
	var res []pdf.Policy
	for _, s := range simul {
		if s == nil || s.Type == nil || *s.Type == "" {
			continue
		}
		p := pdf.Policy{
			Kind:    extractPointer(s.Type),
			Service: extractPointer(s.Service),
			From:    extractPointer(s.DateFrom),
			To:      extractPointer(s.DateTo),
			Price:   extractPointer(s.Value),
		}
		res = append(res, p)
	}
	return res
}
