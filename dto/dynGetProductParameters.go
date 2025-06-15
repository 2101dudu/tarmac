package dto

// DynProductParametersResponse represents the response structure for dynamic product parameters.
type DynProductParametersResponse struct {
	Code            string              `json:"code"`                  // Codigo do produto
	Name            string              `json:"name"`                  // Nome do produto
	Localization    string              `json:"localization"`          // Localizacao do produto
	Country         string              `json:"country"`               // Codigo do Pais do produto
	Zone            string              `json:"zone"`                  // Zona do produto
	ExtraNightsFrom string              `json:"extraNightsFrom"`       // Noites extra sem entradas
	ExtraNightsTo   string              `json:"extraNightsTo"`         // Noites extra sem entradas
	MinPaxReserve   string              `json:"minPaxReserve"`         // Minimo passageiros para reservar
	MaxPaxReserve   string              `json:"maxPaxReserve"`         // Maximo passageiros para reservar
	DepartureLocals []DepartureLocalDTO `json:"departureLocals"`       // Lista de Locais de Partida
	DepartureDates  []DepartureDateDTO  `json:"departureDates"`        // Lista de Datas de Partida
	RoomTypes       []RoomTypeDTO       `json:"roomTypes"`             // Lista de Tipos de Quarto
	BaseLocals      []BaseLocalDTO      `json:"baseLocals"`            // Locais Base do Produto
	ExtraLocals     []ExtraLocalDTO     `json:"extraLocals"`           // Locais Extra do Produto
	RequestAges     *RequestAgesDTO     `json:"requestAges,omitempty"` // Modelo de idades (so criancas ou adultos e criancas)
	Errors          *ErrorDTO           `json:"errors,omitempty"`      // Estrutura de Erros, Caso Aplicavel
}

// DepartureLocalDTO represents a departure local with its properties.
type DepartureLocalDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// DepartureDateDTO represents a departure date with its properties.
type DepartureDateDTO struct {
	Date string `json:"date"`
}

// RoomTypeDTO represents a room type with its properties.
type RoomTypeDTO struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	NumAdults    string `json:"numAdults"`
	NumChilds    string `json:"numChilds"`
	ChildAgeFrom string `json:"childAgeFrom"`
	ChildAgeTo   string `json:"childAgeTo"`
}

// DynBaseLocalDTO represents a base local with its properties.
type BaseLocalDTO struct {
	Code            string `json:"code"`            // Codigo do Local
	Name            string `json:"name"`            // Nome do Local
	MinNights       string `json:"minNights"`       // Minimo de Noites
	MaxNights       string `json:"maxNights"`       // Maximo de Noites
	BeforeAfterBase string `json:"beforeAfterBase"` // Verificacao de extensao dos servicos
}

// DynExtraLocalDTO represents an extra local with its properties.
type ExtraLocalDTO struct {
	Code   string         `json:"code"`   // Codigo do Local
	Name   string         `json:"name"`   // Nome do Local
	Locals []BaseLocalDTO `json:"locals"` // Lista de locais base
}

// DynRequestAgesDTO represents the request ages model.
type RequestAgesDTO struct {
	ReqAgeType  string `json:"reqAgeType"`  // Modelo de pedido de idades
	MinAgeAdult string `json:"minAgeAdult"` // Idade minima do adulto
	MaxAgeAdult string `json:"maxAgeAdult"` // Idade maxima do adulto
}

// ErrorDTO represents the error structure.
type ErrorDTO struct {
	HasErrors string     `json:"hasErrors"` // Indica se existem erros
	ErrorList []ErrorObj `json:"errorList"` // Lista de erros
}

// ErrorObj represents an individual error object.
type ErrorObj struct {
	Code string `json:"code"` // Codigo do erro
	Desc string `json:"desc"` // Descricao do erro
}
