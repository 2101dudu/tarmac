package wsdl

import (
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

type Service struct {
	soapClient      *wbs_pkt_methodsSoap
	soapCredentials *CredentialsStruct
}

func (c *Client) NewWsdlService() *Service {
	return &Service{
		soapClient: NewWbs_pkt_methodsSoap(&soap.Client{
			URL:       c.URL,
			Namespace: Namespace,
		}).(*wbs_pkt_methodsSoap),
		soapCredentials: &CredentialsStruct{
			System:   &c.System,
			Client:   &c.Client,
			Username: &c.Username,
			Password: &c.Password,
		},
	}
}

func (s *Service) SearchProducts() (*SearchProductResponse, error) {
	return s.soapClient.SearchProducts(&SearchProductRequest{
		Credentials: s.soapCredentials,
	})
}

func (s *Service) SearchProductMasterData() (*SearchProductMasterDataResponse, error) {
	return s.soapClient.SearchProductMasterData(&SearchProductMasterDataRequest{
		Credentials: s.soapCredentials,
	})
}

func (s *Service) DynGetProductParameters(prodCodeRaw string) (*DynProductParametersResponse, error) {
	return s.soapClient.DynGetProductParameters(&DynProductParametersRequest{
		Credentials: s.soapCredentials,
		Productcode: &prodCodeRaw,
	})
}

func (s *Service) DynSearchProductAvailableServices(in DynProductAvailableServicesRequest) (*DynProductAvailableServicesResponse, error) {
	in.Credentials = s.soapCredentials
	return s.soapClient.DynSearchProductAvailableServices(&in)
}

func (s *Service) DynSetServicesSelected(in DynServicesSelectedRequest) (*DynServicesSelectedResponse, error) {
	in.Credentials = s.soapCredentials
	return s.soapClient.DynSetServicesSelected(&in)
}

func (s *Service) DynGetOptionalsSelected(in DynGetOptionalsSelectedRequest) (*DynGetOptionalsSelectedResponse, error) {
	in.Credentials = s.soapCredentials
	return s.soapClient.DynGetOptionalsSelected(&in)
}

func (s *Service) DynGetSimulation(in DynGetSimulationRequest) (*DynGetSimulationResponse, error) {
	in.Credentials = s.soapCredentials
	return s.soapClient.DynGetSimulation(&in)
}

func (s *Service) SearchProductsWithBody(in SearchProductRequest) (*SearchProductResponse, error) {
	in.Credentials = s.soapCredentials
	return s.soapClient.SearchProducts(&in)
}

func (s *Service) SearchProductsWithBodyNow(in SearchProductRequest) (*SearchProductResponse, error) {
	in.Credentials = s.soapCredentials
	depDateString := time.Now().Format(time.DateOnly)
	in.DepDate = &depDateString
	return s.soapClient.SearchProducts(&in)
}
