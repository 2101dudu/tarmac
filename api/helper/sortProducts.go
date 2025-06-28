package helper

import (
	"sort"
	"strconv"
	"strings"
	"tarmac/logger"
	"tarmac/wsdl"
	"time"
)

func SortProducts(products []*wsdl.Product, sortBy, sortOrder string) []*wsdl.Product {
	defer logger.Log.TrackTime()()
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
	defer logger.Log.TrackTime()()
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

func ApplyQueryToData(products []*wsdl.Product, country, location, dateFrom string) []*wsdl.Product {
	defer logger.Log.TrackTime()()

	var canonicalizedDate time.Time

	// Normalize dep date
	if dateFrom != "" {
		// convert strings like "2025-1-9" into "2025-01-09"
		// to avoid differnt hash values for the same query
		canonicalizedDate, _ = time.Parse(time.DateOnly, dateFrom)
	}

	var queried []*wsdl.Product

	for _, product := range products {
		if product == nil {
			continue
		}
		match := true
		if country != "" {
			if product.Country == nil || *product.Country != country {
				match = false
			}
		}
		if location != "" {
			if product.Location == nil || *product.Location != location {
				match = false
			}
		}
		if dateFrom != "" {
			if product.OperationFrom == nil || product.OperationTo == nil {
				match = false
			} else {
				// Parse the dates as time.Time for accurate comparison
				layout := time.DateOnly
				of, err2 := time.Parse(layout, *product.OperationFrom)
				ot, err3 := time.Parse(layout, *product.OperationTo)
				if err2 != nil || err3 != nil {
					match = false
				} else if canonicalizedDate.Before(of) || canonicalizedDate.After(ot) {
					match = false
				}
			}

		}
		if match {
			queried = append(queried, product)
		}
	}

	logger.Log.Log("Filtered products. Now have ", len(queried), " products")

	return queried
}
