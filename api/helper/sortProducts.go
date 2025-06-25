package helper

import (
	"sort"
	"strconv"
	"strings"
	"tarmac/wsdl"
)

func SortProducts(products []*wsdl.Product, sortBy, sortOrder string) []*wsdl.Product {
	if len(products) == 0 {
		return products
	}

	ascending := strings.ToLower(sortOrder) != "desc"

	switch strings.ToLower(sortBy) {
	case "price":
		sort.SliceStable(products, func(i int, j int) bool {
			p1, p2 := products[i].PriceFrom, products[j].PriceFrom

			// Null checks just in case
			if p1 == nil || p2 == nil {
				return false
			}

			priceI := atof(*p1)
			priceJ := atof(*p2)
			if priceI == nil || priceJ == nil {
				return false
			}

			res := *priceI < *priceJ
			if ascending {
				return res
			}
			return !res
		})
	case "name":
		sort.SliceStable(products, func(i int, j int) bool {
			// Null checks
			if products[i].Name == nil || products[j].Name == nil {
				return false
			}
			if ascending {
				return *products[i].Name < *products[j].Name
			}
			return *products[i].Name > *products[j].Name
		})
	// more sort types can be added here (e.g. popularity)
	default:
		// No-op for unknown sort types
	}

	return products
}

func FilterProducts(products []*wsdl.Product, priceFrom, priceTo, days string) []*wsdl.Product {
	if len(products) == 0 {
		return products
	}
	var priceFromf, priceTof *float64 = nil, nil
	if priceFrom != "" {
		priceFromf = atof(priceFrom)
	}
	if priceTo != "" {
		priceTof = atof(priceTo)
	}
	daysI, _ := strconv.Atoi(days)

	var filtered []*wsdl.Product

	for _, p := range products {
		if p == nil {
			continue
		}

		// Check price filters
		if p.PriceFrom != nil {
			price := atof(*p.PriceFrom)
			if price == nil {
				continue
			}
			if priceFrom != "" && *price < *priceFromf {
				continue
			}
			if priceTo != "" && *price > *priceTof {
				continue
			}
		}

		// Check days filter
		if days != "" && p.BaseDays != nil {
			n, err := strconv.Atoi(*p.BaseDays)
			if err != nil || n != daysI {
				continue
			}
		}

		filtered = append(filtered, p)
	}

	return filtered
}
