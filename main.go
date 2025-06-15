package main

import (
	"log"
	"os"
	"tarmac/api"
	"tarmac/wsdl"

	"github.com/fiorix/wsdl2go/soap"
	"github.com/joho/godotenv"
)

func loadCredentials() *wsdl.CredentialsStruct {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	system := os.Getenv("SYSTEM")
	client := os.Getenv("CLIENT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	return &wsdl.CredentialsStruct{
		System:     &system,
		Client:     &client,
		Username:   &username,
		Password:   &password,
		OprChannel: nil,
		Lang:       nil,
	}
}

func main() {
	credentials := loadCredentials()

	cli := soap.Client{
		URL:       "https://www3.optitravel.net/optitravel/export/wbs/pkt/wbs_server.php",
		Namespace: wsdl.Namespace,
	}
	soapService := wsdl.NewWbs_pkt_methodsSoap(&cli)

	api := api.Api{SoapService: soapService, Credentials: credentials}

	api.Start()
}
