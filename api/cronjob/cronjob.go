package cronjob

import (
	"strconv"
	"tarmac/db"
	"tarmac/logger"
	"tarmac/wsdl"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SyncAllProducts(service wsdl.Wbs_pkt_methodsSoap, creds *wsdl.CredentialsStruct, dbClient *mongo.Client) error {
	logger.Log.Log("Starting full product sync...")

	resp, err := service.SearchProducts(&wsdl.SearchProductRequest{
		Credentials: creds,
	})
	if err != nil {
		logger.Log.Log("Failed SOAP fetch: ", err)
		return err
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

		if oldProduct.Tags == nil {
			oldProduct.Tags = []string{}
		}

		// wrap code into a struct with product-specific tags
		productWrapper := wsdl.ProductWrapper{Product: *product, Tags: oldProduct.Tags}

		db.RefreshDB(dbClient, "product_index", codeStr, productWrapper)
	}

	logger.Log.Log("Finished syncing ", len(resp.ProductArray.Items), " products.")
	return nil
}
