package wsdl_to_dto

import (
	"tarmac/dto"
	"tarmac/wsdl"
)

// Helper function to convert a wsdl.Product struct with string pointers into a dto.Product with regular strings
func toProductDTO(p *wsdl.Product) dto.Product {
	if p == nil {
		return dto.Product{}
	}

	return dto.Product{
		Code:          deref(p.Code),
		ProdCode:      deref(p.ProdCode),
		Name:          deref(p.Name),
		ImageUrl:      deref(p.ImageUrl),
		Continent:     deref(p.Continent),
		Country:       deref(p.Country),
		Location:      deref(p.Location),
		Zone:          deref(p.Zone),
		BaseDays:      deref(p.BaseDays),
		OperationFrom: deref(p.OperationFrom),
		OperationTo:   deref(p.OperationTo),
		ShortDesc:     deref(p.ShortDesc),
		PopupInfo:     deref(p.PopupInfo),
		PriceFrom:     deref(p.PriceFrom),
		ProdType:      deref(p.ProdType),
		CanReserve:    deref(p.CanReserve),
		Precalc:       deref(p.Precalc),
		HasFlights:    deref(p.HasFlights),
		ItineraryQTY:  deref(p.ItineraryQTY),
	}
}

// Helper function to convert a wsdl.SearchProductResponse into a dto.SearchProductResponse
func ToSearchProductResponseDTO(resp *wsdl.SearchProductResponse) dto.SearchProductResponse {
	dtoResp := dto.SearchProductResponse{
		TotalProducts: deref(resp.TotalProducts),
	}

	if resp.ProductArray != nil {
		for _, p := range resp.ProductArray.Items {
			dtoResp.Products = append(dtoResp.Products, toProductDTO(p))
		}
	}

	return dtoResp
}
