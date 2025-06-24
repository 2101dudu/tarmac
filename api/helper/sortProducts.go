package helper

import (
	"sort"
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
			// Null checks just in case
			if products[i].PriceFrom == nil || products[j].PriceFrom == nil {
				return false
			}
			if ascending {
				return *products[i].PriceFrom < *products[j].PriceFrom
			}
			return *products[i].PriceFrom > *products[j].PriceFrom
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
