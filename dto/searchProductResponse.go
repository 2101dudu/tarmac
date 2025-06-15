package dto

type Product struct {
	Code                   string   `json:"code"`                   // Indice do produto
	ProdCode               string   `json:"prodCode"`               // Codigo do produto
	Name                   string   `json:"name"`                   // Nome do produto
	ImageUrl               string   `json:"imageUrl"`               // Url para a foto
	Continent              string   `json:"continent"`              // Code Continent
	Country                string   `json:"country"`                // Code Country
	Location               string   `json:"location"`               // Code Location
	Zone                   string   `json:"zone"`                   // Code Zone
	BaseDays               string   `json:"baseDays"`               // Dias base
	OperationFrom          string   `json:"operationFrom"`          // Data de inicio da operacao
	OperationTo            string   `json:"operationTo"`            // Data final da operacao
	ShortDesc              string   `json:"shortDesc"`              // Descricao curta
	PopupInfo              string   `json:"popupInfo"`              // Templating de alertas do produto
	PriceFrom              string   `json:"priceFrom"`              // Preco a partir de
	ProdType               string   `json:"prodType"`               // Tipo de produto
	TextContentsArray      []string `json:"textContentsArray"`      // TextContentArray
	PhotoArray             []string `json:"photoArray"`             // Lista de imagens
	GeneralConditionsArray []string `json:"generalConditionsArray"` // GeneralConditionsArray
	CanReserve             string   `json:"canReserve"`             // Produto permite reservas
	Precalc                string   `json:"precalc"`                // Precalculo
	HasFlights             string   `json:"hasFlights"`             // Se o produto tem voos
	ItineraryQTY           string   `json:"itineraryQty"`           // Quantidade de itinerarios
}

type SearchProductResponse struct {
	TotalProducts string    `json:"totalProducts"`
	Products      []Product `json:"products"`
}
