package helper

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"tarmac/wsdl"
)

type productWithPrice struct {
	product *wsdl.Product
	price   float64
}

func trimEmojis(s string) string {
	re := regexp.MustCompile(`&#\d+;`)
	return re.ReplaceAllString(s, "")
}

func trimPrefix(s string) string {
	re := regexp.MustCompile(`^\d+\s*:\s*`)
	return re.ReplaceAllString(s, "")
}

func OptimizeProductSearch(res *wsdl.SearchProductResponse) error {
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

		price, err := strconv.ParseFloat(*product.PriceFrom, 64)
		if err != nil || price <= 0.0 {
			totalProducts--
			continue
		}

		*product.Name = trimEmojis(trimPrefix(*product.Name))

		product.Continent = nil
		product.Country = nil
		product.Location = nil
		product.Zone = nil

		filtered = append(filtered, productWithPrice{product, price})
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].price < filtered[j].price
	})

	cleaned := original[:0]
	for _, pp := range filtered {
		cleaned = append(cleaned, pp.product)
	}

	*res.TotalProducts = strconv.Itoa(totalProducts)
	res.ProductArray.Items = cleaned

	return nil
}
