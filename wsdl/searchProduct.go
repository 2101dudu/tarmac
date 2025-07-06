package wsdl

import (
	"errors"
	"strconv"
	"tarmac/utils"
)

type productWithPrice struct {
	product *Product
	price   float64
}

func OptimizeProductSearch(res *SearchProductResponse) error {
	totalProducts, err := strconv.Atoi(*res.TotalProducts)
	if err != nil {
		return errors.New("invalid TotalProducts")
	}
	original := res.ProductArray.Items

	filtered := make([]productWithPrice, 0, len(original))

	for _, product := range original {
		if product == nil || product.PriceFrom == nil {
			totalProducts--
			continue
		}

		price := atof(*product.PriceFrom)
		if price == nil || *price <= 0.0 {
			totalProducts--
			continue
		}

		*product.Name = utils.SimplifyString(*product.Name)

		product.Continent = nil
		product.Country = nil
		product.Location = nil
		product.Zone = nil

		filtered = append(filtered, productWithPrice{product, *price})
	}

	cleaned := original[:0]
	for _, pp := range filtered {
		cleaned = append(cleaned, pp.product)
	}

	*res.TotalProducts = strconv.Itoa(totalProducts)
	res.ProductArray.Items = cleaned

	return nil
}
