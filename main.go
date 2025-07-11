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

	var clients []wsdl.Client
	for i, _ := range vars.SOAPServices {
		clients = append(clients, wsdl.Client{URL: vars.SOAPServices[i].APIEndpoint, Namespace: wsdl.Namespace, System: vars.SOAPServices[i].System, Client: vars.SOAPServices[i].Client, Username: vars.SOAPServices[i].Username, Password: vars.SOAPServices[i].Password})
	}
	wsdlClient := wsdl.ClientList{Clients: clients}
	wsdlServices := wsdlClient.NewWsdlService()

	cacheClient := cache.Client{Addr: vars.Endpoints.CacheEndpoint, ShortCacheTime: vars.CacheTimes.ShortCacheTime, MediumCacheTime: vars.CacheTimes.MediumCacheTime, LongCacheTime: vars.CacheTimes.LongCacheTime}
	cacheService := cacheClient.NewCacheService()

	dbClient := db.Client{Addr: vars.Endpoints.DBEndpoint, Username: vars.Credentials.DBUsername, Password: vars.Credentials.DBPassword}
	dbService := dbClient.NewDbService()

	logger.Start(vars.Logs.FilePath, true) // temp, change to flag system
	// defer logger.Log.Close()

	mailClient := mail.Client{AgencyEmail: vars.EmailCredentials.AgencyEmail, InternalAgencyEmail: vars.EmailCredentials.InternalAgencyEmail, APIKey: vars.EmailCredentials.APIKey}
	mailService := mailClient.Start()

	coordinatorClient := coordinator.Client{WSDLService: wsdlServices, CacheService: cacheService, DBService: dbService, MailService: mailService}
	coordinatorService := coordinatorClient.NewCoordinatorService()

	api := api.Api{Coordinator: coordinatorService, AdminHashedPassword: vars.Credentials.AdminHashedPassword}

	api.Start()
}
