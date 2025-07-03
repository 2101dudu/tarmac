package cronjob

import (
	"strconv"
	"tarmac/db"
	"tarmac/logger"
	"tarmac/wsdl"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SyncAllProducts(service wsdl.Wbs_pkt_methodsSoap, creds *wsdl.CredentialsStruct, dbClient *mongo.Client) (error, []string, []string) {
	logger.Log.Log("Starting full product sync...")

	collectionName := "product_index"

	countrySet := make(map[string]struct{})
	locationSet := make(map[string]struct{})

	// keep track of products coming from the SOAP API — delete outdated ones
	seenProductCodes := make(map[string]bool)

	depDateString := time.Now().Format(time.DateOnly)

	resp, err := service.SearchProducts(&wsdl.SearchProductRequest{
		Credentials: creds,
		DepDate:     &depDateString,
	})
	if err != nil {
		logger.Log.Log("Failed SOAP fetch: ", err)
		return err, nil, nil
	}

	for _, product := range resp.ProductArray.Items {
		if product.Code == nil {
			logger.Log.Log("Skipping product with nil code")
			continue
		}
		codeStr := *product.Code

		// ensure product code is numeric string
		_, err := strconv.Atoi(codeStr)
		if err != nil {
			logger.Log.Log("Invalid product code (not a number): ", codeStr)
			continue
		}

		// get existing product and store its tags
		oldProduct, _ := db.CheckDBHit[wsdl.ProductWrapper](dbClient, collectionName, codeStr)

		tags := []string{}
		enabled := true
		if oldProduct != nil {
			if oldProduct.Tags != nil {
				tags = oldProduct.Tags
			}
			if oldProduct.Enabled != nil {
				enabled = *oldProduct.Enabled
			}
		}

		// wrap code into a struct with product-specific tags
		productWrapper := wsdl.ProductWrapper{Product: *product, Tags: tags, Enabled: &enabled}

		db.RefreshDB(dbClient, collectionName, codeStr, productWrapper)

		// handle list of countries + locations
		if product.Country != nil {
			countrySet[*product.Country] = struct{}{}
		}
		if product.Location != nil {
			locationSet[*product.Location] = struct{}{}
		}

		seenProductCodes[codeStr] = true
	}

	var countryCodes, locationCodes []string
	for code := range countrySet {
		countryCodes = append(countryCodes, code)
	}
	for code := range locationSet {
		locationCodes = append(locationCodes, code)
	}

	productsList, _ := db.GetAllProducts[wsdl.ProductWrapper](dbClient, collectionName)

	for _, pW := range productsList {
		if pW == nil || pW.Product.Code == nil {
			continue
		}
		code := *pW.Product.Code
		if _, found := seenProductCodes[code]; !found {
			// delete product from DB
			db.DeleteByID(dbClient, collectionName, code)
			logger.Log.Log("Deleted stale product: ", code)
		}
	}

	logger.Log.Log("Finished syncing ", len(resp.ProductArray.Items), " products.")
	return nil, countryCodes, locationCodes
}
