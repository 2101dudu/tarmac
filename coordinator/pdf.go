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
		Insurances:   fillInsurances(simul.Services.Items, simul.Calcs.Items),
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

func fillInsurances(services []*wsdl.DynResServices, calcs []*wsdl.DynResCalcs) []pdf.Insurance {
	var res []pdf.Insurance

	// Build a map of insurance pricing from calcs
	insurancePricing := make(map[string]string)
	for _, c := range calcs {
		if c == nil || c.Service == nil {
			continue
		}
		service := extractPointer(c.Service)
		// Check if this is an insurance service (contains "Seguro" or "Insurance")
		if containsIgnoreCase(service, "Seguro") || containsIgnoreCase(service, "Insurance") {
			desc := extractPointer(c.Description)
			price := extractPointer(c.GrossTotalVal)
			insurancePricing[desc] = price
		}
	}

	// Extract insurance services from services list
	for _, s := range services {
		if s == nil || s.Type == nil {
			continue
		}
		serviceType := extractPointer(s.Type)
		// Check if this is an insurance service
		if containsIgnoreCase(serviceType, "Seguro") || containsIgnoreCase(serviceType, "Insurance") {
			desc := extractPointer(s.Description)
			// Try to get price from pricing map, fallback to empty
			price := insurancePricing[desc]
			if price == "" {
				price = "0.00"
			}

			ins := pdf.Insurance{
				Desc:  desc,
				Type:  serviceType,
				From:  extractPointer(s.DateFrom),
				To:    extractPointer(s.DateTo),
				Price: price,
			}
			res = append(res, ins)
		}
	}

	return res
}

func containsIgnoreCase(s, substr string) bool {
	sLower := ""
	substrLower := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			sLower += string(r + 32)
		} else {
			sLower += string(r)
		}
	}
	for _, r := range substr {
		if r >= 'A' && r <= 'Z' {
			substrLower += string(r + 32)
		} else {
			substrLower += string(r)
		}
	}
	return len(sLower) >= len(substrLower) && func() bool {
		for i := 0; i <= len(sLower)-len(substrLower); i++ {
			if sLower[i:i+len(substrLower)] == substrLower {
				return true
			}
		}
		return false
	}()
}
