package main

import (
	"tarmac/api"
	"tarmac/db"
	"tarmac/env"
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

	dbClient := db.Client{Addr: vars.Endpoints.DBEndpoint}
	dbService := db.Start(&dbClient)

	api := api.Api{SoapService: soapService, Credentials: &wsdl.CredentialsStruct{
		System:   &vars.Credentials.System,
		Client:   &vars.Credentials.Client,
		Username: &vars.Credentials.Username,
		Password: &vars.Credentials.Password,
	}, DBService: dbService, CacheTimes: vars.CacheTimes}

	api.Start()
}
