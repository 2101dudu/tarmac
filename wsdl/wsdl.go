package wsdl

import (
	"errors"
	"strconv"
	"time"

	"github.com/fiorix/wsdl2go/soap"
)

type Client struct {
	URL       string
	Namespace string
	System    string
	Client    string
	Username  string
	Password  string
}

type ClientList struct {
	Clients []Client
}

type Service struct {
	soapClient      *wbs_pkt_methodsSoap
	soapCredentials *CredentialsStruct
}

type ServiceList struct {
	Services map[int]*Service
}

func (c *ClientList) NewWsdlService() *ServiceList {
	services := make(map[int]*Service)

	for i, cli := range c.Clients {
		soap := NewWbs_pkt_methodsSoap(&soap.Client{
			URL:       cli.URL,
			Namespace: Namespace,
		}).(*wbs_pkt_methodsSoap)

		cred := &CredentialsStruct{
			System:   &cli.System,
			Client:   &cli.Client,
			Username: &cli.Username,
			Password: &cli.Password,
		}

		fullService := &Service{soapClient: soap, soapCredentials: cred}

		services[i] = fullService
	}

	return &ServiceList{
		Services: services,
	}
}

func (s *ServiceList) SearchProductMasterData() (*SearchProductMasterDataResponse, error) {
	final := &SearchProductMasterDataResponse{
		SearchProductMasterDataArray: &SearchProductMasterData{
			CountriesArray: &CountryArray{Items: []*Country{}},
			LocationsArray: &LocationArray{Items: []*Location{}},
		},
	}
	for _, service := range s.Services {
		res, err := service.soapClient.SearchProductMasterData(&SearchProductMasterDataRequest{ // TODO: check if master data is the same for all
			Credentials: service.soapCredentials,
		})
		if err != nil {
			return nil, err
		}

		final.SearchProductMasterDataArray.CountriesArray.Items = append(final.SearchProductMasterDataArray.CountriesArray.Items, res.SearchProductMasterDataArray.CountriesArray.Items...)
		final.SearchProductMasterDataArray.LocationsArray.Items = append(final.SearchProductMasterDataArray.LocationsArray.Items, res.SearchProductMasterDataArray.LocationsArray.Items...)
	}

	return final, nil
}

func (s *ServiceList) DynGetProductParameters(prodCodeRaw string, serviceID int) (*DynProductParametersResponse, error) {
	return s.Services[serviceID].soapClient.DynGetProductParameters(&DynProductParametersRequest{
		Credentials: s.Services[serviceID].soapCredentials,
		Productcode: &prodCodeRaw,
	})
}

func (s *ServiceList) DynSearchProductAvailableServices(in DynProductAvailableServicesRequest, serviceID int) (*DynProductAvailableServicesResponse, error) {
	in.Credentials = s.Services[serviceID].soapCredentials
	return s.Services[serviceID].soapClient.DynSearchProductAvailableServices(&in)
}

func (s *ServiceList) DynSetServicesSelected(in DynServicesSelectedRequest, serviceID int) (*DynServicesSelectedResponse, error) {
	in.Credentials = s.Services[serviceID].soapCredentials
	return s.Services[serviceID].soapClient.DynSetServicesSelected(&in)
}

func (s *ServiceList) DynGetOptionalsSelected(in DynGetOptionalsSelectedRequest, serviceID int) (*DynGetOptionalsSelectedResponse, error) {
	in.Credentials = s.Services[serviceID].soapCredentials
	return s.Services[serviceID].soapClient.DynGetOptionalsSelected(&in)
}

func (s *ServiceList) DynGetSimulation(in DynGetSimulationRequest, serviceID int) (*DynGetSimulationResponse, error) {
	in.Credentials = s.Services[serviceID].soapCredentials
	return s.Services[serviceID].soapClient.DynGetSimulation(&in)
}

func (s *ServiceList) SearchProductsWithBodyNow(in SearchProductRequest) (*SearchProductResponse, error) {
	depDateString := time.Now().Format(time.DateOnly)
	in.DepDate = &depDateString

	var prodArray []*Product
	tp := 0
	for i, service := range s.Services {
		in.Credentials = service.soapCredentials

		resp, err := service.soapClient.SearchProducts(&in)
		if err != nil {
			return nil, errors.New("Invalid service: " + *service.soapCredentials.Client + ": " + err.Error())
		}
		t, err := strconv.Atoi(*resp.TotalProducts)
		if err != nil {
			return nil, errors.New("Invalid number: " + *resp.TotalProducts)
		}
		tp += t
		prodArray = append(prodArray, resp.ProductArray.Items...)

		for _, prod := range resp.ProductArray.Items {
			*prod.Code += "-" + strconv.Itoa(i)
		}
	}

	totalProductsString := strconv.Itoa(tp)
	return &SearchProductResponse{TotalProducts: &totalProductsString, ProductArray: &ProductArray{
		Items: prodArray,
	},
	}, nil
}
