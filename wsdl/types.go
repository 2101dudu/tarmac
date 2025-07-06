package wsdl

type ProductWithExtraData struct {
	Data             DynProductParametersResponse `json:"data"`
	DescriptionArray TextContentArray             `json:"descriptionArray"`
	PhotoArray       PhotoArray                   `json:"photoArray"`
	Price            string                       `json:"price"`
}
