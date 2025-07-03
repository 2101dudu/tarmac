package cronjob

import (
	"strconv"
	"tarmac/db"
	"tarmac/logger"
	"tarmac/wsdl"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SyncAllProducts(service wsdl.Wbs_pkt_methodsSoap, creds *wsdl.CredentialsStruct, dbClient *mongo.Client) (error, []string, []string) {
	logger.Log.Log("Starting full product sync...")

	countrySet := make(map[string]struct{})
	locationSet := make(map[string]struct{})

	resp, err := service.SearchProducts(&wsdl.SearchProductRequest{
		Credentials: creds,
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
		oldProduct, _ := db.CheckDBHit[wsdl.ProductWrapper](dbClient, "product_index", codeStr)

		var tags []string
		if oldProduct == nil || oldProduct.Tags == nil {
			tags = []string{}
		} else {
			tags = oldProduct.Tags
		}

		// wrap code into a struct with product-specific tags
		productWrapper := wsdl.ProductWrapper{Product: *product, Tags: tags}

		db.RefreshDB(dbClient, "product_index", codeStr, productWrapper)

		// handle list of countries + locations
		if product.Country != nil {
			countrySet[*product.Country] = struct{}{}
		}
		if product.Location != nil {
			locationSet[*product.Location] = struct{}{}
		}
	}

	var countryCodes, locationCodes []string
	for code := range countrySet {
		countryCodes = append(countryCodes, code)
	}
	for code := range locationSet {
		locationCodes = append(locationCodes, code)
	}

	logger.Log.Log("Finished syncing ", len(resp.ProductArray.Items), " products.")
	return nil, countryCodes, locationCodes
}
