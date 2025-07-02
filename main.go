package main

import (
	"tarmac/api"
	"tarmac/cache"
	"tarmac/db"
	"tarmac/env"
	"tarmac/logger"
	"tarmac/wsdl"

	"github.com/fiorix/wsdl2go/soap"
)

func main() {
	vars := env.LoadEnvFile()

	soapClient := soap.Client{
		URL:       vars.Endpoints.APIEndpoint,
		Namespace: wsdl.Namespace,
	}
	soapService := wsdl.NewWbs_pkt_methodsSoap(&soapClient)

	cacheClient := cache.Client{Addr: vars.Endpoints.CacheEndpoint}
	cacheService := cache.Start(&cacheClient)

	dbClient := db.Client{Addr: vars.Endpoints.DBEndpoint, Username: vars.Credentials.DBUsername, Password: vars.Credentials.DBPassword}
	dbService := db.Start(&dbClient)

	logger.Start(vars.Logs.FilePath, true) // temp, change to flag system
	// defer logger.Log.Close()

	api := api.Api{SoapService: soapService, Credentials: &wsdl.CredentialsStruct{
		System:   &vars.Credentials.System,
		Client:   &vars.Credentials.Client,
		Username: &vars.Credentials.Username,
		Password: &vars.Credentials.Password,
	}, CacheService: cacheService, CacheTimes: vars.CacheTimes, DBService: dbService, AdminHashedPassword: vars.Credentials.AdminHashedPassword}

	api.Start()
}
