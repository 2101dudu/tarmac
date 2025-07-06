package main

import (
	"tarmac/api"
	"tarmac/cache"
	"tarmac/coordinator"
	"tarmac/db"
	"tarmac/env"
	"tarmac/logger"
	"tarmac/mail"
	"tarmac/wsdl"
)

func main() {
	vars := env.LoadEnvFile()

	wsdlClient := wsdl.Client{URL: vars.Endpoints.APIEndpoint, Namespace: wsdl.Namespace, System: vars.Credentials.System, Client: vars.Credentials.Client, Username: vars.Credentials.Username, Password: vars.Credentials.Password}
	wsdlService := wsdlClient.NewWsdlService()

	cacheClient := cache.Client{Addr: vars.Endpoints.CacheEndpoint, ShortCacheTime: vars.CacheTimes.ShortCacheTime, MediumCacheTime: vars.CacheTimes.MediumCacheTime, LongCacheTime: vars.CacheTimes.LongCacheTime}
	cacheService := cacheClient.NewCacheService()

	dbClient := db.Client{Addr: vars.Endpoints.DBEndpoint, Username: vars.Credentials.DBUsername, Password: vars.Credentials.DBPassword}
	dbService := dbClient.NewDbService()

	logger.Start(vars.Logs.FilePath, true) // temp, change to flag system
	// defer logger.Log.Close()

	mailClient := mail.Client{AgencyEmail: vars.EmailCredentials.AgencyEmail, InternalAgencyEmail: vars.EmailCredentials.InternalAgencyEmail, APIKey: vars.EmailCredentials.APIKey}
	mailService := mailClient.Start()

	coordinatorService := coordinator.Service{WSDLService: wsdlService, CacheService: cacheService, DBService: dbService, MailService: mailService}

	api := api.Api{Coordinator: &coordinatorService, AdminHashedPassword: vars.Credentials.AdminHashedPassword}

	api.Start()
}
