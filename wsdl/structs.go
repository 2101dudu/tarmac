package wsdl

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://www3.optitravel.net/optitravel/export/wbs/pkt/"

// NewWbs_pkt_methodsSoap creates an initializes a Wbs_pkt_methodsSoap.
func NewWbs_pkt_methodsSoap(cli *soap.Client) Wbs_pkt_methodsSoap {
	return &wbs_pkt_methodsSoap{cli}
}

// Wbs_pkt_methodsSoap was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Wbs_pkt_methodsSoap interface {
	// CancelReserve was auto-generated from WSDL.
	CancelReserve(input_arr *CancelReserveRequest) (*CancelReserveResponse, error)

	// DynCreateReserve was auto-generated from WSDL.
	DynCreateReserve(input_arr *DynCreateReserveRequest) (*DynCreateReserveResponse, error)

	// DynGetOptionalsSelected was auto-generated from WSDL.
	DynGetOptionalsSelected(input_arr *DynGetOptionalsSelectedRequest) (*DynGetOptionalsSelectedResponse, error)

	// DynGetProductParameters was auto-generated from WSDL.
	DynGetProductParameters(input_arr *DynProductParametersRequest) (*DynProductParametersResponse, error)

	// DynGetSimulation was auto-generated from WSDL.
	DynGetSimulation(input_arr *DynGetSimulationRequest) (*DynGetSimulationResponse, error)

	// DynPreBooking was auto-generated from WSDL.
	DynPreBooking(input_arr *DynPreBookingRequest) (*DynPreBookingResponse, error)

	// DynPriceAvailResume was auto-generated from WSDL.
	DynPriceAvailResume(input_arr *DynPriceAvailResumeRequest) (*DynPriceAvailResumeResponse, error)

	// DynProductCache was auto-generated from WSDL.
	DynProductCache(input_arr *DynProductCacheRequest) (*DynProductCacheResponse, error)

	// DynProductOptionals was auto-generated from WSDL.
	DynProductOptionals(input_arr *DynProductOptionalsRequest) (*DynProductOptionalsResponse, error)

	// DynRemoveOptionals was auto-generated from WSDL.
	DynRemoveOptionals(input_arr *DynRemoveOptionalsRequest) (*DynRemoveOptionalsResponse, error)

	// DynSearchProductAvailableServices was auto-generated from WSDL.
	DynSearchProductAvailableServices(input_arr *DynProductAvailableServicesRequest) (*DynProductAvailableServicesResponse, error)

	// DynSetServicesSelected was auto-generated from WSDL.
	DynSetServicesSelected(input_arr *DynServicesSelectedRequest) (*DynServicesSelectedResponse, error)

	// GetClients was auto-generated from WSDL.
	GetClients(input_arr *GetClientsRequest) (*GetClientsResponse, error)

	// GetHotelDetails was auto-generated from WSDL.
	GetHotelDetails(input_arr *GetHotelDetailsRequest) (*GetHotelDetailsResponse, error)

	// GetInfoPax was auto-generated from WSDL.
	GetInfoPax(input_arr *GetPaxInfoRequest) (*GetPaxInfoResponse, error)

	// GetPaxAdditionalInfo was auto-generated from WSDL.
	GetPaxAdditionalInfo(input_arr *GetPaxAdditionalInfoRequest) (*GetPaxAdditionalInfoResponse, error)

	// GetStores was auto-generated from WSDL.
	GetStores(input_arr *GetStoresRequest) (*GetStoresResponse, error)

	// ReserveList was auto-generated from WSDL.
	ReserveList(input_arr *ReserveListRequest) (*ReserveListResponse, error)

	// RetrieveReserve was auto-generated from WSDL.
	RetrieveReserve(input_arr *RetrieveReserveRequest) (*RetrieveReserveResponse, error)

	// SearchProductByCalendar was auto-generated from WSDL.
	SearchProductByCalendar(input_arr *SearchProductByCalendarRequest) (*SearchProductByCalendarResponse, error)

	// SearchProductMasterData was auto-generated from WSDL.
	SearchProductMasterData(input_arr *SearchProductMasterDataRequest) (*SearchProductMasterDataResponse, error)

	// SearchProductTypes was auto-generated from WSDL.
	SearchProductTypes(input_arr *SearchProductTypesRequest) (*SearchProductTypesResponse, error)

	// SearchProducts was auto-generated from WSDL.
	SearchProducts(input_arr *SearchProductRequest) (*SearchProductResponse, error)

	// SetBookingOffline was auto-generated from WSDL.
	SetBookingOffline(input_arr *SetBookingOfflineRequest) (*SetBookingOfflineResponse, error)

	// SetPaxAdditionalInfo was auto-generated from WSDL.
	SetPaxAdditionalInfo(input_arr *SetPaxAdditionalInfoRequest) (*SetPaxAdditionalInfoResponse, error)

	// SetReserveHotels was auto-generated from WSDL.
	SetReserveHotels(input_arr *SetReserveHotelsRequest) (*SetReserveHotelsResponse, error)

	// SimulCancelReserve was auto-generated from WSDL.
	SimulCancelReserve(input_arr *SimulCancelReserveRequest) (*SimulCancelReserveResponse, error)

	// StaticCreateReserve was auto-generated from WSDL.
	StaticCreateReserve(input_arr *StaticCreateReserveRequest) (*StaticCreateReserveResponse, error)

	// StaticGetPaxsStruct was auto-generated from WSDL.
	StaticGetPaxsStruct(input_arr *StaticGetPaxsStructRequest) (*StaticGetPaxsStructResponse, error)

	// StaticGetProductParameters was auto-generated from WSDL.
	StaticGetProductParameters(input_arr *StaticGetProductParametersRequest) (*StaticGetProductParametersResponse, error)

	// StaticGetReserveFlights was auto-generated from WSDL.
	StaticGetReserveFlights(input_arr *StaticGetReserveFlightsRequest) (*StaticGetReserveFlightsResponse, error)

	// StaticGetReserveHotels was auto-generated from WSDL.
	StaticGetReserveHotels(input_arr *StaticGetReserveHotelsRequest) (*StaticGetReserveHotelsResponse, error)

	// StaticGetReserveRooming was auto-generated from WSDL.
	StaticGetReserveRooming(input_arr *StaticGetReserveRoomingRequest) (*StaticGetReserveRoomingResponse, error)

	// StaticGetRoomsStruct was auto-generated from WSDL.
	StaticGetRoomsStruct(input_arr *StaticGetRoomsStructRequest) (*StaticGetRoomsStructResponse, error)

	// StaticGetSimulation was auto-generated from WSDL.
	StaticGetSimulation(input_arr *StaticGetSimulationRequest) (*StaticGetSimulationResponse, error)

	// StaticSetReserveExtraNights was auto-generated from WSDL.
	StaticSetReserveExtraNights(input_arr *StaticSetReserveExtraNightsRequest) (*StaticSetReserveExtraNightsResponse, error)

	// StaticSetReserveFlights was auto-generated from WSDL.
	StaticSetReserveFlights(input_arr *StaticSetReserveFlightsRequest) (*StaticSetReserveFlightsResponse, error)

	// StaticSetReserveRooming was auto-generated from WSDL.
	StaticSetReserveRooming(input_arr *StaticSetReserveRoomingRequest) (*StaticSetReserveRoomingResponse, error)

	// ValidateVISA was auto-generated from WSDL.
	ValidateVISA(input_arr *ValidateVISARequest) (*ValidateVISAResponse, error)
}

// Accommodation was auto-generated from WSDL.
type Accommodation struct {
	Adults *string `xml:"Adults,omitempty" json:"Adults,omitempty" yaml:"Adults,omitempty"`
	Childs *string `xml:"Childs,omitempty" json:"Childs,omitempty" yaml:"Childs,omitempty"`
	Ages   *string `xml:"Ages,omitempty" json:"Ages,omitempty" yaml:"Ages,omitempty"`
}

// AccommodationArray was auto-generated from WSDL.
type AccommodationArray struct {
	Items []*Accommodation `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Additional was auto-generated from WSDL.
type Additional struct {
	AttrId       *string `xml:"AttrId,omitempty" json:"AttrId,omitempty" yaml:"AttrId,omitempty"`
	AttrName     *string `xml:"AttrName,omitempty" json:"AttrName,omitempty" yaml:"AttrName,omitempty"`
	AttrValue    *string `xml:"AttrValue,omitempty" json:"AttrValue,omitempty" yaml:"AttrValue,omitempty"`
	AttrCode     *string `xml:"AttrCode,omitempty" json:"AttrCode,omitempty" yaml:"AttrCode,omitempty"`
	AttrType     *string `xml:"AttrType,omitempty" json:"AttrType,omitempty" yaml:"AttrType,omitempty"`
	AttrTypeInfo *string `xml:"AttrTypeInfo,omitempty" json:"AttrTypeInfo,omitempty" yaml:"AttrTypeInfo,omitempty"`
	AttrVisible  *string `xml:"AttrVisible,omitempty" json:"AttrVisible,omitempty" yaml:"AttrVisible,omitempty"`
}

// AdditionalArray was auto-generated from WSDL.
type AdditionalArray struct {
	Items []*Additional `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Adt was auto-generated from WSDL.
type Adt struct {
	Quant *string `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
}

// AdtArray was auto-generated from WSDL.
type AdtArray struct {
	Items []*Adt `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Boards was auto-generated from WSDL.
type Boards struct {
	Code *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

// BoardsArray was auto-generated from WSDL.
type BoardsArray struct {
	Items []*Boards `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// BoardsChild was auto-generated from WSDL.
type BoardsChild struct {
	Id        *string `xml:"Id,omitempty" json:"Id,omitempty" yaml:"Id,omitempty"`
	Type      *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Age       *string `xml:"Age,omitempty" json:"Age,omitempty" yaml:"Age,omitempty"`
	Text      *string `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`
	RegText   *string `xml:"RegText,omitempty" json:"RegText,omitempty" yaml:"RegText,omitempty"`
	SellPrice *string `xml:"SellPrice,omitempty" json:"SellPrice,omitempty" yaml:"SellPrice,omitempty"`
	BuyPrice  *string `xml:"BuyPrice,omitempty" json:"BuyPrice,omitempty" yaml:"BuyPrice,omitempty"`
}

// BoardsChildArray was auto-generated from WSDL.
type BoardsChildArray struct {
	Items []*BoardsChild `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// BoardsHtl was auto-generated from WSDL.
type BoardsHtl struct {
	Id          *string           `xml:"Id,omitempty" json:"Id,omitempty" yaml:"Id,omitempty"`
	Type        *string           `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Text        *string           `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`
	RegText     *string           `xml:"RegText,omitempty" json:"RegText,omitempty" yaml:"RegText,omitempty"`
	SellPrice   *string           `xml:"SellPrice,omitempty" json:"SellPrice,omitempty" yaml:"SellPrice,omitempty"`
	BuyPrice    *string           `xml:"BuyPrice,omitempty" json:"BuyPrice,omitempty" yaml:"BuyPrice,omitempty"`
	BoardsChild *BoardsChildArray `xml:"BoardsChild,omitempty" json:"BoardsChild,omitempty" yaml:"BoardsChild,omitempty"`
}

// BoardsHtlArray was auto-generated from WSDL.
type BoardsHtlArray struct {
	Items []*BoardsHtl `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CancelReserveRequest was auto-generated from WSDL.
type CancelReserveRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"reserve,omitempty" json:"reserve,omitempty" yaml:"reserve,omitempty"`
}

// CancelReserveResponse was auto-generated from WSDL.
type CancelReserveResponse struct {
	Status *string      `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Errors *ErrorStruct `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// Chd was auto-generated from WSDL.
type Chd struct {
	Quant       *string      `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	ChdAgeArray *ChdAgeArray `xml:"ChdAgeArray,omitempty" json:"ChdAgeArray,omitempty" yaml:"ChdAgeArray,omitempty"`
}

// ChdAge was auto-generated from WSDL.
type ChdAge struct {
	Age *string `xml:"Age,omitempty" json:"Age,omitempty" yaml:"Age,omitempty"`
}

// ChdAgeArray was auto-generated from WSDL.
type ChdAgeArray struct {
	Items []*ChdAge `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ChdArray was auto-generated from WSDL.
type ChdArray struct {
	Items []*Chd `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ChildBoards was auto-generated from WSDL.
type ChildBoards struct {
	Age       *string `xml:"Age,omitempty" json:"Age,omitempty" yaml:"Age,omitempty"`
	SellValue *string `xml:"SellValue,omitempty" json:"SellValue,omitempty" yaml:"SellValue,omitempty"`
	SellQuant *string `xml:"SellQuant,omitempty" json:"SellQuant,omitempty" yaml:"SellQuant,omitempty"`
}

// ChildBoardsArray was auto-generated from WSDL.
type ChildBoardsArray struct {
	Items []*ChildBoards `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ChildStruct was auto-generated from WSDL.
type ChildStruct struct {
	Min *string `xml:"Min,omitempty" json:"Min,omitempty" yaml:"Min,omitempty"`
	Max *string `xml:"Max,omitempty" json:"Max,omitempty" yaml:"Max,omitempty"`
}

// Clients was auto-generated from WSDL.
type Clients struct {
	Id     *string `xml:"Id,omitempty" json:"Id,omitempty" yaml:"Id,omitempty"`
	ExtRef *string `xml:"ExtRef,omitempty" json:"ExtRef,omitempty" yaml:"ExtRef,omitempty"`
	Name   *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Status *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// ClientsArray was auto-generated from WSDL.
type ClientsArray struct {
	Items []*Clients `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Continent was auto-generated from WSDL.
type Continent struct {
	Code        *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

// ContinentArray was auto-generated from WSDL.
type ContinentArray struct {
	Items []*Continent `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Countries was auto-generated from WSDL.
type Countries struct {
	Code *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

// CountriesArray was auto-generated from WSDL.
type CountriesArray struct {
	Items []*Countries `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Country was auto-generated from WSDL.
type Country struct {
	Code        *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Continent   *string `xml:"Continent,omitempty" json:"Continent,omitempty" yaml:"Continent,omitempty"`
}

// CountryArray was auto-generated from WSDL.
type CountryArray struct {
	Items []*Country `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CredentialsStruct was auto-generated from WSDL.
type CredentialsStruct struct {
	System     *string `xml:"System,omitempty" json:"System,omitempty" yaml:"System,omitempty"`
	Client     *string `xml:"Client,omitempty" json:"Client,omitempty" yaml:"Client,omitempty"`
	Username   *string `xml:"Username,omitempty" json:"Username,omitempty" yaml:"Username,omitempty"`
	Password   *string `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	OprChannel *string `xml:"OprChannel,omitempty" json:"OprChannel,omitempty" yaml:"OprChannel,omitempty"`
	Lang       *string `xml:"Lang,omitempty" json:"Lang,omitempty" yaml:"Lang,omitempty"`
}

// DepDate was auto-generated from WSDL.
type DepDate struct {
	DepDate_ID  *string `xml:"DepDate_ID,omitempty" json:"DepDate_ID,omitempty" yaml:"DepDate_ID,omitempty"`
	DepDate_DSC *string `xml:"DepDate_DSC,omitempty" json:"DepDate_DSC,omitempty" yaml:"DepDate_DSC,omitempty"`
	Allot       *string `xml:"Allot,omitempty" json:"Allot,omitempty" yaml:"Allot,omitempty"`
	Allot_Free  *string `xml:"Allot_Free,omitempty" json:"Allot_Free,omitempty" yaml:"Allot_Free,omitempty"`
	Allot_OBS   *string `xml:"Allot_OBS,omitempty" json:"Allot_OBS,omitempty" yaml:"Allot_OBS,omitempty"`
}

// DepDateArray was auto-generated from WSDL.
type DepDateArray struct {
	Items []*DepDate `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DepLocal was auto-generated from WSDL.
type DepLocal struct {
	DepLocal_ID      *string `xml:"DepLocal_ID,omitempty" json:"DepLocal_ID,omitempty" yaml:"DepLocal_ID,omitempty"`
	DepLocal_DSC     *string `xml:"DepLocal_DSC,omitempty" json:"DepLocal_DSC,omitempty" yaml:"DepLocal_DSC,omitempty"`
	DepLocal_Address *string `xml:"DepLocal_Address,omitempty" json:"DepLocal_Address,omitempty" yaml:"DepLocal_Address,omitempty"`
}

// DepLocalArray was auto-generated from WSDL.
type DepLocalArray struct {
	Items []*DepLocal `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Deplocal was auto-generated from WSDL.
type Deplocal struct {
	DepLocal_ID      *string `xml:"DepLocal_ID,omitempty" json:"DepLocal_ID,omitempty" yaml:"DepLocal_ID,omitempty"`
	DepLocal_DSC     *string `xml:"DepLocal_DSC,omitempty" json:"DepLocal_DSC,omitempty" yaml:"DepLocal_DSC,omitempty"`
	DepLocal_Address *string `xml:"DepLocal_Address,omitempty" json:"DepLocal_Address,omitempty" yaml:"DepLocal_Address,omitempty"`
}

// DeplocalArray was auto-generated from WSDL.
type DeplocalArray struct {
	Items []*Deplocal `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynBaseLocal was auto-generated from WSDL.
type DynBaseLocal struct {
	Code            *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name            *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	MinNights       *string `xml:"MinNights,omitempty" json:"MinNights,omitempty" yaml:"MinNights,omitempty"`
	MaxNights       *string `xml:"MaxNights,omitempty" json:"MaxNights,omitempty" yaml:"MaxNights,omitempty"`
	BeforeAfterBase *string `xml:"BeforeAfterBase,omitempty" json:"BeforeAfterBase,omitempty" yaml:"BeforeAfterBase,omitempty"`
}

// DynBaseLocalArray was auto-generated from WSDL.
type DynBaseLocalArray struct {
	Items []*DynBaseLocal `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynBoards was auto-generated from WSDL.
type DynBoards struct {
	Code        *string           `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	ShortCode   *string           `xml:"ShortCode,omitempty" json:"ShortCode,omitempty" yaml:"ShortCode,omitempty"`
	Name        *string           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	SellValue   *string           `xml:"SellValue,omitempty" json:"SellValue,omitempty" yaml:"SellValue,omitempty"`
	ChildBoards *ChildBoardsArray `xml:"ChildBoards,omitempty" json:"ChildBoards,omitempty" yaml:"ChildBoards,omitempty"`
}

// DynBoardsArray was auto-generated from WSDL.
type DynBoardsArray struct {
	Items []*DynBoards `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynBookingServices was auto-generated from WSDL.
type DynBookingServices struct {
	Type        *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Status      *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// DynBookingServicesArray was auto-generated from WSDL.
type DynBookingServicesArray struct {
	Items []*DynBookingServices `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynCancelPenalties was auto-generated from WSDL.
type DynCancelPenalties struct {
	Deadline *DynDeadlineArray `xml:"Deadline,omitempty" json:"Deadline,omitempty" yaml:"Deadline,omitempty"`
	Charge   *DynChargeArray   `xml:"Charge,omitempty" json:"Charge,omitempty" yaml:"Charge,omitempty"`
}

// DynCancelPenaltiesArray was auto-generated from WSDL.
type DynCancelPenaltiesArray struct {
	Items []*DynCancelPenalties `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynCharge was auto-generated from WSDL.
type DynCharge struct {
	Amount   *string `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
	Currency *string `xml:"Currency,omitempty" json:"Currency,omitempty" yaml:"Currency,omitempty"`
}

// DynChargeArray was auto-generated from WSDL.
type DynChargeArray struct {
	Items []*DynCharge `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynChildAgesPktStatic was auto-generated from WSDL.
type DynChildAgesPktStatic struct {
	Min *string `xml:"Min,omitempty" json:"Min,omitempty" yaml:"Min,omitempty"`
	Max *string `xml:"Max,omitempty" json:"Max,omitempty" yaml:"Max,omitempty"`
}

// DynChildAgesPktStaticArray was auto-generated from WSDL.
type DynChildAgesPktStaticArray struct {
	Items []*DynChildAgesPktStatic `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynChildPrice was auto-generated from WSDL.
type DynChildPrice struct {
	ChildAge    *string `xml:"ChildAge,omitempty" json:"ChildAge,omitempty" yaml:"ChildAge,omitempty"`
	ChildMinAge *string `xml:"ChildMinAge,omitempty" json:"ChildMinAge,omitempty" yaml:"ChildMinAge,omitempty"`
	ChildMaxAge *string `xml:"ChildMaxAge,omitempty" json:"ChildMaxAge,omitempty" yaml:"ChildMaxAge,omitempty"`
	PriceType   *string `xml:"PriceType,omitempty" json:"PriceType,omitempty" yaml:"PriceType,omitempty"`
	Price       *string `xml:"Price,omitempty" json:"Price,omitempty" yaml:"Price,omitempty"`
}

// DynChildPriceArray was auto-generated from WSDL.
type DynChildPriceArray struct {
	Items []*DynChildPrice `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynCreateReserveRequest was auto-generated from WSDL.
type DynCreateReserveRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	PaxStruct   *DynPaxArray       `xml:"PaxStruct,omitempty" json:"PaxStruct,omitempty" yaml:"PaxStruct,omitempty"`
	ClientRef   *string            `xml:"ClientRef,omitempty" json:"ClientRef,omitempty" yaml:"ClientRef,omitempty"`
	UserRef     *string            `xml:"UserRef,omitempty" json:"UserRef,omitempty" yaml:"UserRef,omitempty"`
	Store       *string            `xml:"Store,omitempty" json:"Store,omitempty" yaml:"Store,omitempty"`
	ClientAgent *string            `xml:"ClientAgent,omitempty" json:"ClientAgent,omitempty" yaml:"ClientAgent,omitempty"`
	ClientObs   *string            `xml:"ClientObs,omitempty" json:"ClientObs,omitempty" yaml:"ClientObs,omitempty"`
	Type        *string            `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	CreateUser  *string            `xml:"CreateUser,omitempty" json:"CreateUser,omitempty" yaml:"CreateUser,omitempty"`
}

// DynCreateReserveResponse was auto-generated from WSDL.
type DynCreateReserveResponse struct {
	Reserve      *string      `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	SystemStatus *string      `xml:"SystemStatus,omitempty" json:"SystemStatus,omitempty" yaml:"SystemStatus,omitempty"`
	StatusAll    *string      `xml:"StatusAll,omitempty" json:"StatusAll,omitempty" yaml:"StatusAll,omitempty"`
	StatusHTL    *string      `xml:"StatusHTL,omitempty" json:"StatusHTL,omitempty" yaml:"StatusHTL,omitempty"`
	StatusAVI    *string      `xml:"StatusAVI,omitempty" json:"StatusAVI,omitempty" yaml:"StatusAVI,omitempty"`
	AbreuRef     *string      `xml:"AbreuRef,omitempty" json:"AbreuRef,omitempty" yaml:"AbreuRef,omitempty"`
	Errors       *ErrorStruct `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// DynCriteriaList was auto-generated from WSDL.
type DynCriteriaList struct {
	WorkId               *string                       `xml:"WorkId,omitempty" json:"WorkId,omitempty" yaml:"WorkId,omitempty"`
	ProductCode          *string                       `xml:"ProductCode,omitempty" json:"ProductCode,omitempty" yaml:"ProductCode,omitempty"`
	DateList             *string                       `xml:"DateList,omitempty" json:"DateList,omitempty" yaml:"DateList,omitempty"`
	Nights               *string                       `xml:"Nights,omitempty" json:"Nights,omitempty" yaml:"Nights,omitempty"`
	DepartureLocalCode   *string                       `xml:"DepartureLocalCode,omitempty" json:"DepartureLocalCode,omitempty" yaml:"DepartureLocalCode,omitempty"`
	Status               *string                       `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Urgent               *string                       `xml:"Urgent,omitempty" json:"Urgent,omitempty" yaml:"Urgent,omitempty"`
	UnprocessedDatesList *DynUnprocessedDatesListArray `xml:"UnprocessedDatesList,omitempty" json:"UnprocessedDatesList,omitempty" yaml:"UnprocessedDatesList,omitempty"`
}

// DynCriteriaListArray was auto-generated from WSDL.
type DynCriteriaListArray struct {
	Items []*DynCriteriaList `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynCriteriaListLog was auto-generated from WSDL.
type DynCriteriaListLog struct {
	CriteriaWorkId     *string `xml:"CriteriaWorkId,omitempty" json:"CriteriaWorkId,omitempty" yaml:"CriteriaWorkId,omitempty"`
	ProductCode        *string `xml:"ProductCode,omitempty" json:"ProductCode,omitempty" yaml:"ProductCode,omitempty"`
	DepartureLocalCode *string `xml:"DepartureLocalCode,omitempty" json:"DepartureLocalCode,omitempty" yaml:"DepartureLocalCode,omitempty"`
	DepDate            *string `xml:"DepDate,omitempty" json:"DepDate,omitempty" yaml:"DepDate,omitempty"`
	UpdateTime         *string `xml:"UpdateTime,omitempty" json:"UpdateTime,omitempty" yaml:"UpdateTime,omitempty"`
	RunTime            *string `xml:"RunTime,omitempty" json:"RunTime,omitempty" yaml:"RunTime,omitempty"`
	State              *string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
}

// DynCriteriaListLogArray was auto-generated from WSDL.
type DynCriteriaListLogArray struct {
	Items []*DynCriteriaListLog `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynDates was auto-generated from WSDL.
type DynDates struct {
	Order *string `xml:"Order,omitempty" json:"Order,omitempty" yaml:"Order,omitempty"`
	Date  *string `xml:"Date,omitempty" json:"Date,omitempty" yaml:"Date,omitempty"`
}

// DynDatesArray was auto-generated from WSDL.
type DynDatesArray struct {
	Items []*DynDates `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynDeadline was auto-generated from WSDL.
type DynDeadline struct {
	TimeUnit *string `xml:"TimeUnit,omitempty" json:"TimeUnit,omitempty" yaml:"TimeUnit,omitempty"`
	Units    *string `xml:"Units,omitempty" json:"Units,omitempty" yaml:"Units,omitempty"`
}

// DynDeadlineArray was auto-generated from WSDL.
type DynDeadlineArray struct {
	Items []*DynDeadline `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynDepLocals was auto-generated from WSDL.
type DynDepLocals struct {
	Code        *string        `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description *string        `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Stops       *DynStopsArray `xml:"Stops,omitempty" json:"Stops,omitempty" yaml:"Stops,omitempty"`
}

// DynDepLocalsArray was auto-generated from WSDL.
type DynDepLocalsArray struct {
	Items []*DynDepLocals `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynDepartureDate was auto-generated from WSDL.
type DynDepartureDate struct {
	Date                 *string                `xml:"Date,omitempty" json:"Date,omitempty" yaml:"Date,omitempty"`
	ExtraNightsEnableIn  *string                `xml:"ExtraNightsEnableIn,omitempty" json:"ExtraNightsEnableIn,omitempty" yaml:"ExtraNightsEnableIn,omitempty"`
	ExtraNightsEnableOut *string                `xml:"ExtraNightsEnableOut,omitempty" json:"ExtraNightsEnableOut,omitempty" yaml:"ExtraNightsEnableOut,omitempty"`
	DepartureLocalsParam *string                `xml:"DepartureLocalsParam,omitempty" json:"DepartureLocalsParam,omitempty" yaml:"DepartureLocalsParam,omitempty"`
	DepartureLocations   *DynDepartureDateArray `xml:"DepartureLocations,omitempty" json:"DepartureLocations,omitempty" yaml:"DepartureLocations,omitempty"`
}

// DynDepartureDateArray was auto-generated from WSDL.
type DynDepartureDateArray struct {
	Items []*DynDepartureDate `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynDepartureLocal was auto-generated from WSDL.
type DynDepartureLocal struct {
	Code *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

// DynDepartureLocalArray was auto-generated from WSDL.
type DynDepartureLocalArray struct {
	Items []*DynDepartureLocal `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynExtraInfo was auto-generated from WSDL.
type DynExtraInfo struct {
	Id             *string       `xml:"Id,omitempty" json:"Id,omitempty" yaml:"Id,omitempty"`
	Name           *string       `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	InputMaxLenght *string       `xml:"InputMaxLenght,omitempty" json:"InputMaxLenght,omitempty" yaml:"InputMaxLenght,omitempty"`
	Type           *string       `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	MapType        *string       `xml:"MapType,omitempty" json:"MapType,omitempty" yaml:"MapType,omitempty"`
	Show           *string       `xml:"Show,omitempty" json:"Show,omitempty" yaml:"Show,omitempty"`
	Required       *string       `xml:"Required,omitempty" json:"Required,omitempty" yaml:"Required,omitempty"`
	OptList        *DynListArray `xml:"OptList,omitempty" json:"OptList,omitempty" yaml:"OptList,omitempty"`
}

// DynExtraInfoArray was auto-generated from WSDL.
type DynExtraInfoArray struct {
	Items []*DynExtraInfo `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynExtraLocal was auto-generated from WSDL.
type DynExtraLocal struct {
	Code   *string            `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name   *string            `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Locals *DynBaseLocalArray `xml:"Locals,omitempty" json:"Locals,omitempty" yaml:"Locals,omitempty"`
}

// DynExtraLocalArray was auto-generated from WSDL.
type DynExtraLocalArray struct {
	Items []*DynExtraLocal `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynExtraLocalReq was auto-generated from WSDL.
type DynExtraLocalReq struct {
	Code   *string     `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Locals *LocalArray `xml:"Locals,omitempty" json:"Locals,omitempty" yaml:"Locals,omitempty"`
}

// DynExtraLocalReqArray was auto-generated from WSDL.
type DynExtraLocalReqArray struct {
	Items []*DynExtraLocalReq `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynFlight was auto-generated from WSDL.
type DynFlight struct {
	FlightCode     *string `xml:"FlightCode,omitempty" json:"FlightCode,omitempty" yaml:"FlightCode,omitempty"`
	From           *string `xml:"From,omitempty" json:"From,omitempty" yaml:"From,omitempty"`
	FromDesc       *string `xml:"FromDesc,omitempty" json:"FromDesc,omitempty" yaml:"FromDesc,omitempty"`
	To             *string `xml:"To,omitempty" json:"To,omitempty" yaml:"To,omitempty"`
	ToDesc         *string `xml:"ToDesc,omitempty" json:"ToDesc,omitempty" yaml:"ToDesc,omitempty"`
	DepDate        *string `xml:"DepDate,omitempty" json:"DepDate,omitempty" yaml:"DepDate,omitempty"`
	DepTime        *string `xml:"DepTime,omitempty" json:"DepTime,omitempty" yaml:"DepTime,omitempty"`
	ArrDate        *string `xml:"ArrDate,omitempty" json:"ArrDate,omitempty" yaml:"ArrDate,omitempty"`
	ArrTime        *string `xml:"ArrTime,omitempty" json:"ArrTime,omitempty" yaml:"ArrTime,omitempty"`
	AirCompCode    *string `xml:"AirCompCode,omitempty" json:"AirCompCode,omitempty" yaml:"AirCompCode,omitempty"`
	AirCompLogoUrl *string `xml:"AirCompLogoUrl,omitempty" json:"AirCompLogoUrl,omitempty" yaml:"AirCompLogoUrl,omitempty"`
	AirVndCompCode *string `xml:"AirVndCompCode,omitempty" json:"AirVndCompCode,omitempty" yaml:"AirVndCompCode,omitempty"`
	TransportType  *string `xml:"TransportType,omitempty" json:"TransportType,omitempty" yaml:"TransportType,omitempty"`
	CodeShare      *string `xml:"CodeShare,omitempty" json:"CodeShare,omitempty" yaml:"CodeShare,omitempty"`
	Number         *string `xml:"Number,omitempty" json:"Number,omitempty" yaml:"Number,omitempty"`
	Class          *string `xml:"Class,omitempty" json:"Class,omitempty" yaml:"Class,omitempty"`
	Cabin          *string `xml:"Cabin,omitempty" json:"Cabin,omitempty" yaml:"Cabin,omitempty"`
	Bag            *string `xml:"Bag,omitempty" json:"Bag,omitempty" yaml:"Bag,omitempty"`
	Status         *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Selected       *string `xml:"Selected,omitempty" json:"Selected,omitempty" yaml:"Selected,omitempty"`
}

// DynFlightArray was auto-generated from WSDL.
type DynFlightArray struct {
	Items []*DynFlight `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynFlightGroup was auto-generated from WSDL.
type DynFlightGroup struct {
	FlightGroupCode *string         `xml:"FlightGroupCode,omitempty" json:"FlightGroupCode,omitempty" yaml:"FlightGroupCode,omitempty"`
	NumStopOvers    *string         `xml:"NumStopOvers,omitempty" json:"NumStopOvers,omitempty" yaml:"NumStopOvers,omitempty"`
	Flights         *DynFlightArray `xml:"Flights,omitempty" json:"Flights,omitempty" yaml:"Flights,omitempty"`
}

// DynFlightGroupArray was auto-generated from WSDL.
type DynFlightGroupArray struct {
	Items []*DynFlightGroup `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynFlightMainGroup was auto-generated from WSDL.
type DynFlightMainGroup struct {
	Description               *string               `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	FlightOptions             *DynFlightOptionArray `xml:"FlightOptions,omitempty" json:"FlightOptions,omitempty" yaml:"FlightOptions,omitempty"`
	FlightOptionsSuperBB      *DynFlightOptionArray `xml:"FlightOptionsSuperBB,omitempty" json:"FlightOptionsSuperBB,omitempty" yaml:"FlightOptionsSuperBB,omitempty"`
	FlightOptionsFlightEngine *DynFlightOptionArray `xml:"FlightOptionsFlightEngine,omitempty" json:"FlightOptionsFlightEngine,omitempty" yaml:"FlightOptionsFlightEngine,omitempty"`
}

// DynFlightMainGroupArray was auto-generated from WSDL.
type DynFlightMainGroupArray struct {
	Items []*DynFlightMainGroup `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynFlightOption was auto-generated from WSDL.
type DynFlightOption struct {
	OptionCode         *string                 `xml:"OptionCode,omitempty" json:"OptionCode,omitempty" yaml:"OptionCode,omitempty"`
	FlightType         *string                 `xml:"FlightType,omitempty" json:"FlightType,omitempty" yaml:"FlightType,omitempty"`
	AirAvailVersion    *string                 `xml:"AirAvailVersion,omitempty" json:"AirAvailVersion,omitempty" yaml:"AirAvailVersion,omitempty"`
	RateTaxVal         *string                 `xml:"RateTaxVal,omitempty" json:"RateTaxVal,omitempty" yaml:"RateTaxVal,omitempty"`
	SuplementsTotalVal *string                 `xml:"SuplementsTotalVal,omitempty" json:"SuplementsTotalVal,omitempty" yaml:"SuplementsTotalVal,omitempty"`
	CheapestDifPax     *string                 `xml:"CheapestDifPax,omitempty" json:"CheapestDifPax,omitempty" yaml:"CheapestDifPax,omitempty"`
	Tax                *string                 `xml:"Tax,omitempty" json:"Tax,omitempty" yaml:"Tax,omitempty"`
	Lasttkdt           *string                 `xml:"Lasttkdt,omitempty" json:"Lasttkdt,omitempty" yaml:"Lasttkdt,omitempty"`
	FlightSegments     *DynFlightSegmentsArray `xml:"FlightSegments,omitempty" json:"FlightSegments,omitempty" yaml:"FlightSegments,omitempty"`
}

// DynFlightOptionArray was auto-generated from WSDL.
type DynFlightOptionArray struct {
	Items []*DynFlightOption `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynFlightSegmentSel was auto-generated from WSDL.
type DynFlightSegmentSel struct {
	SegmentCode     *string `xml:"SegmentCode,omitempty" json:"SegmentCode,omitempty" yaml:"SegmentCode,omitempty"`
	FlightGroupCode *string `xml:"FlightGroupCode,omitempty" json:"FlightGroupCode,omitempty" yaml:"FlightGroupCode,omitempty"`
}

// DynFlightSegmentSelArray was auto-generated from WSDL.
type DynFlightSegmentSelArray struct {
	Items []*DynFlightSegmentSel `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynFlightSegments was auto-generated from WSDL.
type DynFlightSegments struct {
	ServiceDesc  *string              `xml:"ServiceDesc,omitempty" json:"ServiceDesc,omitempty" yaml:"ServiceDesc,omitempty"`
	SegmentCode  *string              `xml:"SegmentCode,omitempty" json:"SegmentCode,omitempty" yaml:"SegmentCode,omitempty"`
	FromIATA     *string              `xml:"FromIATA,omitempty" json:"FromIATA,omitempty" yaml:"FromIATA,omitempty"`
	FromIATADesc *string              `xml:"FromIATADesc,omitempty" json:"FromIATADesc,omitempty" yaml:"FromIATADesc,omitempty"`
	ToIATA       *string              `xml:"ToIATA,omitempty" json:"ToIATA,omitempty" yaml:"ToIATA,omitempty"`
	ToIATADesc   *string              `xml:"ToIATADesc,omitempty" json:"ToIATADesc,omitempty" yaml:"ToIATADesc,omitempty"`
	Label        *string              `xml:"Label,omitempty" json:"Label,omitempty" yaml:"Label,omitempty"`
	Date         *string              `xml:"Date,omitempty" json:"Date,omitempty" yaml:"Date,omitempty"`
	Flights      *DynFlightGroupArray `xml:"Flights,omitempty" json:"Flights,omitempty" yaml:"Flights,omitempty"`
}

// DynFlightSegmentsArray was auto-generated from WSDL.
type DynFlightSegmentsArray struct {
	Items []*DynFlightSegments `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynFlightsSelected was auto-generated from WSDL.
type DynFlightsSelected struct {
	OptionCode   *string                   `xml:"OptionCode,omitempty" json:"OptionCode,omitempty" yaml:"OptionCode,omitempty"`
	SegmentLists *DynFlightSegmentSelArray `xml:"SegmentLists,omitempty" json:"SegmentLists,omitempty" yaml:"SegmentLists,omitempty"`
}

// DynFlightsSelectedArray was auto-generated from WSDL.
type DynFlightsSelectedArray struct {
	Items []*DynFlightsSelected `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynGetOptionalsSelectedRequest was auto-generated from WSDL.
type DynGetOptionalsSelectedRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
}

// DynGetOptionalsSelectedResponse was auto-generated from WSDL.
type DynGetOptionalsSelectedResponse struct {
	SelOptionals *DynSelOptionalsArray `xml:"SelOptionals,omitempty" json:"SelOptionals,omitempty" yaml:"SelOptionals,omitempty"`
	Errors       *ErrorStruct          `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// DynGetSimulationRequest was auto-generated from WSDL.
type DynGetSimulationRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	Prebooking  *string            `xml:"Prebooking,omitempty" json:"Prebooking,omitempty" yaml:"Prebooking,omitempty"`
}

// DynGetSimulationResponse was auto-generated from WSDL.
type DynGetSimulationResponse struct {
	BillingType     *string                 `xml:"BillingType,omitempty" json:"BillingType,omitempty" yaml:"BillingType,omitempty"`
	AvailableStatus *string                 `xml:"AvailableStatus,omitempty" json:"AvailableStatus,omitempty" yaml:"AvailableStatus,omitempty"`
	TotalPrice      *string                 `xml:"TotalPrice,omitempty" json:"TotalPrice,omitempty" yaml:"TotalPrice,omitempty"`
	Currency        *string                 `xml:"Currency,omitempty" json:"Currency,omitempty" yaml:"Currency,omitempty"`
	Conditions      *string                 `xml:"Conditions,omitempty" json:"Conditions,omitempty" yaml:"Conditions,omitempty"`
	Services        *DynResServicesArray    `xml:"Services,omitempty" json:"Services,omitempty" yaml:"Services,omitempty"`
	Calcs           *DynResCalcsArray       `xml:"Calcs,omitempty" json:"Calcs,omitempty" yaml:"Calcs,omitempty"`
	Remarks         *DynResRemarksArray     `xml:"Remarks,omitempty" json:"Remarks,omitempty" yaml:"Remarks,omitempty"`
	Pax             *DynResPaxsArray        `xml:"Pax,omitempty" json:"Pax,omitempty" yaml:"Pax,omitempty"`
	HotelWarnings   *DynHotelWarningsArray  `xml:"HotelWarnings,omitempty" json:"HotelWarnings,omitempty" yaml:"HotelWarnings,omitempty"`
	RoomsPenalties  *DynRoomsPenaltiesArray `xml:"RoomsPenalties,omitempty" json:"RoomsPenalties,omitempty" yaml:"RoomsPenalties,omitempty"`
	Policies        *DynPoliciesArray       `xml:"Policies,omitempty" json:"Policies,omitempty" yaml:"Policies,omitempty"`
	ResumeCalcs     *DynResumeCalcsArray    `xml:"ResumeCalcs,omitempty" json:"ResumeCalcs,omitempty" yaml:"ResumeCalcs,omitempty"`
	Errors          *ErrorStruct            `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// DynHotelOption was auto-generated from WSDL.
type DynHotelOption struct {
	Code                      *string                         `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	MasterCode                *string                         `xml:"MasterCode,omitempty" json:"MasterCode,omitempty" yaml:"MasterCode,omitempty"`
	HotelRPH                  *string                         `xml:"HotelRPH,omitempty" json:"HotelRPH,omitempty" yaml:"HotelRPH,omitempty"`
	Name                      *string                         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	RatingType                *string                         `xml:"RatingType,omitempty" json:"RatingType,omitempty" yaml:"RatingType,omitempty"`
	CheckIn                   *string                         `xml:"CheckIn,omitempty" json:"CheckIn,omitempty" yaml:"CheckIn,omitempty"`
	CheckOut                  *string                         `xml:"CheckOut,omitempty" json:"CheckOut,omitempty" yaml:"CheckOut,omitempty"`
	Rating                    *string                         `xml:"Rating,omitempty" json:"Rating,omitempty" yaml:"Rating,omitempty"`
	RNT                       *string                         `xml:"RNT,omitempty" json:"RNT,omitempty" yaml:"RNT,omitempty"`
	Image                     *string                         `xml:"Image,omitempty" json:"Image,omitempty" yaml:"Image,omitempty"`
	Address                   *string                         `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	PostalCode                *string                         `xml:"PostalCode,omitempty" json:"PostalCode,omitempty" yaml:"PostalCode,omitempty"`
	Country                   *string                         `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
	Local                     *string                         `xml:"Local,omitempty" json:"Local,omitempty" yaml:"Local,omitempty"`
	Zone                      *string                         `xml:"Zone,omitempty" json:"Zone,omitempty" yaml:"Zone,omitempty"`
	ZonesDesc                 *string                         `xml:"ZonesDesc,omitempty" json:"ZonesDesc,omitempty" yaml:"ZonesDesc,omitempty"`
	GpsLatitude               *string                         `xml:"GpsLatitude,omitempty" json:"GpsLatitude,omitempty" yaml:"GpsLatitude,omitempty"`
	GpsLongitude              *string                         `xml:"GpsLongitude,omitempty" json:"GpsLongitude,omitempty" yaml:"GpsLongitude,omitempty"`
	ExternalRef               *string                         `xml:"ExternalRef,omitempty" json:"ExternalRef,omitempty" yaml:"ExternalRef,omitempty"`
	HBERef                    *string                         `xml:"HBERef,omitempty" json:"HBERef,omitempty" yaml:"HBERef,omitempty"`
	JuniperRef                *string                         `xml:"JuniperRef,omitempty" json:"JuniperRef,omitempty" yaml:"JuniperRef,omitempty"`
	GiataRef                  *string                         `xml:"GiataRef,omitempty" json:"GiataRef,omitempty" yaml:"GiataRef,omitempty"`
	ExpediaRef                *string                         `xml:"ExpediaRef,omitempty" json:"ExpediaRef,omitempty" yaml:"ExpediaRef,omitempty"`
	HotelBedsRef              *string                         `xml:"HotelBedsRef,omitempty" json:"HotelBedsRef,omitempty" yaml:"HotelBedsRef,omitempty"`
	PriceFrom                 *string                         `xml:"PriceFrom,omitempty" json:"PriceFrom,omitempty" yaml:"PriceFrom,omitempty"`
	PriceFromGrouped          *string                         `xml:"PriceFromGrouped,omitempty" json:"PriceFromGrouped,omitempty" yaml:"PriceFromGrouped,omitempty"`
	CatalogCode               *string                         `xml:"CatalogCode,omitempty" json:"CatalogCode,omitempty" yaml:"CatalogCode,omitempty"`
	ProgramCode               *string                         `xml:"ProgramCode,omitempty" json:"ProgramCode,omitempty" yaml:"ProgramCode,omitempty"`
	SelectedStartDate         *string                         `xml:"SelectedStartDate,omitempty" json:"SelectedStartDate,omitempty" yaml:"SelectedStartDate,omitempty"`
	SelectedEndDate           *string                         `xml:"SelectedEndDate,omitempty" json:"SelectedEndDate,omitempty" yaml:"SelectedEndDate,omitempty"`
	ImageUrl                  *string                         `xml:"ImageUrl,omitempty" json:"ImageUrl,omitempty" yaml:"ImageUrl,omitempty"`
	Store                     *string                         `xml:"Store,omitempty" json:"Store,omitempty" yaml:"Store,omitempty"`
	Integration               *string                         `xml:"Integration,omitempty" json:"Integration,omitempty" yaml:"Integration,omitempty"`
	RoomsOccupancy            *DynRoomsOccupancyArray         `xml:"RoomsOccupancy,omitempty" json:"RoomsOccupancy,omitempty" yaml:"RoomsOccupancy,omitempty"`
	Boards                    *DynBoardsArray                 `xml:"Boards,omitempty" json:"Boards,omitempty" yaml:"Boards,omitempty"`
	Restriction               *DynRestrictionArray            `xml:"Restriction,omitempty" json:"Restriction,omitempty" yaml:"Restriction,omitempty"`
	DepLocals                 *DynDepLocalsArray              `xml:"DepLocals,omitempty" json:"DepLocals,omitempty" yaml:"DepLocals,omitempty"`
	DynSelectedRooms          *DynSelectedRoomsOfiArray       `xml:"DynSelectedRooms,omitempty" json:"DynSelectedRooms,omitempty" yaml:"DynSelectedRooms,omitempty"`
	DynSelectedPktStaticRooms *DynSelectedRoomsPktStaticArray `xml:"DynSelectedPktStaticRooms,omitempty" json:"DynSelectedPktStaticRooms,omitempty" yaml:"DynSelectedPktStaticRooms,omitempty"`
	DynPktStaticOptions       *DynPktStaticOptionsArray       `xml:"DynPktStaticOptions,omitempty" json:"DynPktStaticOptions,omitempty" yaml:"DynPktStaticOptions,omitempty"`
	Suplements                *DynSuplementsArray             `xml:"Suplements,omitempty" json:"Suplements,omitempty" yaml:"Suplements,omitempty"`
	Scoring                   *string                         `xml:"Scoring,omitempty" json:"Scoring,omitempty" yaml:"Scoring,omitempty"`
	Token                     *string                         `xml:"Token,omitempty" json:"Token,omitempty" yaml:"Token,omitempty"`
	HasMore                   *bool                           `xml:"HasMore,omitempty" json:"HasMore,omitempty" yaml:"HasMore,omitempty"`
}

// DynHotelOptionArray was auto-generated from WSDL.
type DynHotelOptionArray struct {
	Items []*DynHotelOption `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynHotelWarnings was auto-generated from WSDL.
type DynHotelWarnings struct {
	Message *string `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
}

// DynHotelWarningsArray was auto-generated from WSDL.
type DynHotelWarningsArray struct {
	Items []*DynHotelWarnings `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynHotelsSelected was auto-generated from WSDL.
type DynHotelsSelected struct {
	ItineraryCode      *string              `xml:"ItineraryCode,omitempty" json:"ItineraryCode,omitempty" yaml:"ItineraryCode,omitempty"`
	Interface          *string              `xml:"Interface,omitempty" json:"Interface,omitempty" yaml:"Interface,omitempty"`
	Connector          *string              `xml:"Connector,omitempty" json:"Connector,omitempty" yaml:"Connector,omitempty"`
	HtlSession         *string              `xml:"HtlSession,omitempty" json:"HtlSession,omitempty" yaml:"HtlSession,omitempty"`
	HotelRPH           *string              `xml:"HotelRPH,omitempty" json:"HotelRPH,omitempty" yaml:"HotelRPH,omitempty"`
	HotelSelected      *string              `xml:"HotelSelected,omitempty" json:"HotelSelected,omitempty" yaml:"HotelSelected,omitempty"`
	HotelSelectedPrice *string              `xml:"HotelSelectedPrice,omitempty" json:"HotelSelectedPrice,omitempty" yaml:"HotelSelectedPrice,omitempty"`
	RoomsSelected      *DynRoomsSelectArray `xml:"RoomsSelected,omitempty" json:"RoomsSelected,omitempty" yaml:"RoomsSelected,omitempty"`
	BoardSelected      *string              `xml:"BoardSelected,omitempty" json:"BoardSelected,omitempty" yaml:"BoardSelected,omitempty"`
	SuplementSelected  *string              `xml:"SuplementSelected,omitempty" json:"SuplementSelected,omitempty" yaml:"SuplementSelected,omitempty"`
}

// DynHotelsSelectedArray was auto-generated from WSDL.
type DynHotelsSelectedArray struct {
	Items []*DynHotelsSelected `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynInfantAgesPktStatic was auto-generated from WSDL.
type DynInfantAgesPktStatic struct {
	Min *string `xml:"Min,omitempty" json:"Min,omitempty" yaml:"Min,omitempty"`
	Max *string `xml:"Max,omitempty" json:"Max,omitempty" yaml:"Max,omitempty"`
}

// DynInfantAgesPktStaticArray was auto-generated from WSDL.
type DynInfantAgesPktStaticArray struct {
	Items []*DynInfantAgesPktStatic `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynInitOptions was auto-generated from WSDL.
type DynInitOptions struct {
	WorkId             *string `xml:"WorkId,omitempty" json:"WorkId,omitempty" yaml:"WorkId,omitempty"`
	DateList           *string `xml:"DateList,omitempty" json:"DateList,omitempty" yaml:"DateList,omitempty"`
	Nights             *string `xml:"Nights,omitempty" json:"Nights,omitempty" yaml:"Nights,omitempty"`
	DepartureLocalCode *string `xml:"DepartureLocalCode,omitempty" json:"DepartureLocalCode,omitempty" yaml:"DepartureLocalCode,omitempty"`
	Status             *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Urgent             *string `xml:"Urgent,omitempty" json:"Urgent,omitempty" yaml:"Urgent,omitempty"`
}

// DynInsurance was auto-generated from WSDL.
type DynInsurance struct {
	Included *Included      `xml:"Included,omitempty" json:"Included,omitempty" yaml:"Included,omitempty"`
	Upgrades *UpgradesArray `xml:"Upgrades,omitempty" json:"Upgrades,omitempty" yaml:"Upgrades,omitempty"`
}

// DynItinerary was auto-generated from WSDL.
type DynItinerary struct {
	Code                       *string                          `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Skeleton_id                *string                          `xml:"skeleton_id,omitempty" json:"skeleton_id,omitempty" yaml:"skeleton_id,omitempty"`
	Name                       *string                          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Interface                  *string                          `xml:"Interface,omitempty" json:"Interface,omitempty" yaml:"Interface,omitempty"`
	Connector                  *string                          `xml:"Connector,omitempty" json:"Connector,omitempty" yaml:"Connector,omitempty"`
	HtlSession                 *string                          `xml:"HtlSession,omitempty" json:"HtlSession,omitempty" yaml:"HtlSession,omitempty"`
	ResStatus                  *string                          `xml:"ResStatus,omitempty" json:"ResStatus,omitempty" yaml:"ResStatus,omitempty"`
	ResStatusDesc              *string                          `xml:"ResStatusDesc,omitempty" json:"ResStatusDesc,omitempty" yaml:"ResStatusDesc,omitempty"`
	ProdType                   *string                          `xml:"ProdType,omitempty" json:"ProdType,omitempty" yaml:"ProdType,omitempty"`
	DynChildAgesPktStatic      *DynChildAgesPktStaticArray      `xml:"DynChildAgesPktStatic,omitempty" json:"DynChildAgesPktStatic,omitempty" yaml:"DynChildAgesPktStatic,omitempty"`
	DynInfantAgesPktStatic     *DynInfantAgesPktStaticArray     `xml:"DynInfantAgesPktStatic,omitempty" json:"DynInfantAgesPktStatic,omitempty" yaml:"DynInfantAgesPktStatic,omitempty"`
	DynSeniorPktStatic         *DynSeniorPktStaticArray         `xml:"DynSeniorPktStatic,omitempty" json:"DynSeniorPktStatic,omitempty" yaml:"DynSeniorPktStatic,omitempty"`
	DynReductionNightPktStatic *DynReductionNightPktStaticArray `xml:"DynReductionNightPktStatic,omitempty" json:"DynReductionNightPktStatic,omitempty" yaml:"DynReductionNightPktStatic,omitempty"`
	HotelOption                *DynHotelOptionArray             `xml:"HotelOption,omitempty" json:"HotelOption,omitempty" yaml:"HotelOption,omitempty"`
}

// DynItineraryArray was auto-generated from WSDL.
type DynItineraryArray struct {
	Items []*DynItinerary `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynList was auto-generated from WSDL.
type DynList struct {
	Code        *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

// DynListArray was auto-generated from WSDL.
type DynListArray struct {
	Items []*DynList `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynOfiSelected was auto-generated from WSDL.
type DynOfiSelected struct {
}

// DynOfiSelectedArray was auto-generated from WSDL.
type DynOfiSelectedArray struct {
	Items []*DynOfiSelected `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynOptionals was auto-generated from WSDL.
type DynOptionals struct {
	Code             *string             `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	ItemId           *string             `xml:"ItemId,omitempty" json:"ItemId,omitempty" yaml:"ItemId,omitempty"`
	ItemDescription  *string             `xml:"itemDescription,omitempty" json:"itemDescription,omitempty" yaml:"itemDescription,omitempty"`
	Type             *string             `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Name             *string             `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	ShortDescription *string             `xml:"ShortDescription,omitempty" json:"ShortDescription,omitempty" yaml:"ShortDescription,omitempty"`
	LongDescription  *string             `xml:"LongDescription,omitempty" json:"LongDescription,omitempty" yaml:"LongDescription,omitempty"`
	FromType         *string             `xml:"FromType,omitempty" json:"FromType,omitempty" yaml:"FromType,omitempty"`
	ToType           *string             `xml:"ToType,omitempty" json:"ToType,omitempty" yaml:"ToType,omitempty"`
	FromDetails      *string             `xml:"FromDetails,omitempty" json:"FromDetails,omitempty" yaml:"FromDetails,omitempty"`
	ToDetails        *string             `xml:"ToDetails,omitempty" json:"ToDetails,omitempty" yaml:"ToDetails,omitempty"`
	Price            *string             `xml:"Price,omitempty" json:"Price,omitempty" yaml:"Price,omitempty"`
	PriceType        *string             `xml:"PriceType,omitempty" json:"PriceType,omitempty" yaml:"PriceType,omitempty"`
	SelectAllPax     *string             `xml:"SelectAllPax,omitempty" json:"SelectAllPax,omitempty" yaml:"SelectAllPax,omitempty"`
	ContractBuyId    *string             `xml:"ContractBuyId,omitempty" json:"ContractBuyId,omitempty" yaml:"ContractBuyId,omitempty"`
	ContractSellId   *string             `xml:"ContractSellId,omitempty" json:"ContractSellId,omitempty" yaml:"ContractSellId,omitempty"`
	Localization     *string             `xml:"Localization,omitempty" json:"Localization,omitempty" yaml:"Localization,omitempty"`
	Zone             *string             `xml:"Zone,omitempty" json:"Zone,omitempty" yaml:"Zone,omitempty"`
	ImageURI         *string             `xml:"ImageURI,omitempty" json:"ImageURI,omitempty" yaml:"ImageURI,omitempty"`
	PreSelDate       *string             `xml:"PreSelDate,omitempty" json:"PreSelDate,omitempty" yaml:"PreSelDate,omitempty"`
	PacoteDias       *string             `xml:"PacoteDias,omitempty" json:"PacoteDias,omitempty" yaml:"PacoteDias,omitempty"`
	Mandatory        *string             `xml:"Mandatory,omitempty" json:"Mandatory,omitempty" yaml:"Mandatory,omitempty"`
	UniqueSelect     *string             `xml:"UniqueSelect,omitempty" json:"UniqueSelect,omitempty" yaml:"UniqueSelect,omitempty"`
	PickupTime       *DynPickupTimeArray `xml:"PickupTime,omitempty" json:"PickupTime,omitempty" yaml:"PickupTime,omitempty"`
	Dates            *DynDatesArray      `xml:"Dates,omitempty" json:"Dates,omitempty" yaml:"Dates,omitempty"`
	PriceChilds      *DynChildPriceArray `xml:"PriceChilds,omitempty" json:"PriceChilds,omitempty" yaml:"PriceChilds,omitempty"`
	Suplements       *DynSuplementsArray `xml:"Suplements,omitempty" json:"Suplements,omitempty" yaml:"Suplements,omitempty"`
}

// DynOptionalsArray was auto-generated from WSDL.
type DynOptionalsArray struct {
	Items []*DynOptionals `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynOptionalsSelected was auto-generated from WSDL.
type DynOptionalsSelected struct {
	Code                *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Date                *string `xml:"Date,omitempty" json:"Date,omitempty" yaml:"Date,omitempty"`
	Adults              *string `xml:"Adults,omitempty" json:"Adults,omitempty" yaml:"Adults,omitempty"`
	ChildAges           *string `xml:"ChildAges,omitempty" json:"ChildAges,omitempty" yaml:"ChildAges,omitempty"`
	PickUp              *string `xml:"PickUp,omitempty" json:"PickUp,omitempty" yaml:"PickUp,omitempty"`
	PickUpTime          *string `xml:"PickUpTime,omitempty" json:"PickUpTime,omitempty" yaml:"PickUpTime,omitempty"`
	Drop                *string `xml:"Drop,omitempty" json:"Drop,omitempty" yaml:"Drop,omitempty"`
	OriginFlight        *string `xml:"OriginFlight,omitempty" json:"OriginFlight,omitempty" yaml:"OriginFlight,omitempty"`
	OriginAirport       *string `xml:"OriginAirport,omitempty" json:"OriginAirport,omitempty" yaml:"OriginAirport,omitempty"`
	FlightDepartureTime *string `xml:"FlightDepartureTime,omitempty" json:"FlightDepartureTime,omitempty" yaml:"FlightDepartureTime,omitempty"`
	ItemId              *string `xml:"ItemId,omitempty" json:"ItemId,omitempty" yaml:"ItemId,omitempty"`
	DropObs             *string `xml:"DropObs,omitempty" json:"DropObs,omitempty" yaml:"DropObs,omitempty"`
	OptionId            *string `xml:"OptionId,omitempty" json:"OptionId,omitempty" yaml:"OptionId,omitempty"`
	ProductName         *string `xml:"ProductName,omitempty" json:"ProductName,omitempty" yaml:"ProductName,omitempty"`
	Localization        *string `xml:"Localization,omitempty" json:"Localization,omitempty" yaml:"Localization,omitempty"`
	Zone                *string `xml:"Zone,omitempty" json:"Zone,omitempty" yaml:"Zone,omitempty"`
	ContractBuyId       *string `xml:"ContractBuyId,omitempty" json:"ContractBuyId,omitempty" yaml:"ContractBuyId,omitempty"`
	ContractSellId      *string `xml:"ContractSellId,omitempty" json:"ContractSellId,omitempty" yaml:"ContractSellId,omitempty"`
	PacoteDias          *string `xml:"PacoteDias,omitempty" json:"PacoteDias,omitempty" yaml:"PacoteDias,omitempty"`
	PickType            *string `xml:"PickType,omitempty" json:"PickType,omitempty" yaml:"PickType,omitempty"`
	DropType            *string `xml:"DropType,omitempty" json:"DropType,omitempty" yaml:"DropType,omitempty"`
	Type                *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	FromCode            *string `xml:"FromCode,omitempty" json:"FromCode,omitempty" yaml:"FromCode,omitempty"`
	ToCode              *string `xml:"ToCode,omitempty" json:"ToCode,omitempty" yaml:"ToCode,omitempty"`
	DateDiffTxt         *string `xml:"DateDiffTxt,omitempty" json:"DateDiffTxt,omitempty" yaml:"DateDiffTxt,omitempty"`
	DurationDays        *string `xml:"DurationDays,omitempty" json:"DurationDays,omitempty" yaml:"DurationDays,omitempty"`
}

// DynOptionalsSelectedArray was auto-generated from WSDL.
type DynOptionalsSelectedArray struct {
	Items []*DynOptionalsSelected `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynOtherMandatoryServices was auto-generated from WSDL.
type DynOtherMandatoryServices struct {
	Type        *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

// DynOtherMandatoryServicesArray was auto-generated from WSDL.
type DynOtherMandatoryServicesArray struct {
	Items []*DynOtherMandatoryServices `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynPax was auto-generated from WSDL.
type DynPax struct {
	Type        *string              `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Title       *string              `xml:"Title,omitempty" json:"Title,omitempty" yaml:"Title,omitempty"`
	Age         *string              `xml:"Age,omitempty" json:"Age,omitempty" yaml:"Age,omitempty"`
	Name        *string              `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Surname     *string              `xml:"Surname,omitempty" json:"Surname,omitempty" yaml:"Surname,omitempty"`
	Room        *string              `xml:"Room,omitempty" json:"Room,omitempty" yaml:"Room,omitempty"`
	RoomNum     *string              `xml:"RoomNum,omitempty" json:"RoomNum,omitempty" yaml:"RoomNum,omitempty"`
	DateBirth   *string              `xml:"DateBirth,omitempty" json:"DateBirth,omitempty" yaml:"DateBirth,omitempty"`
	Holder      *string              `xml:"Holder,omitempty" json:"Holder,omitempty" yaml:"Holder,omitempty"`
	XtraPaxInfo *DynXtraPaxInfoArray `xml:"XtraPaxInfo,omitempty" json:"XtraPaxInfo,omitempty" yaml:"XtraPaxInfo,omitempty"`
}

// DynPaxArray was auto-generated from WSDL.
type DynPaxArray struct {
	Items []*DynPax `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynPenaltiesDetail was auto-generated from WSDL.
type DynPenaltiesDetail struct {
	Info *DynPenaltiesInfoArray `xml:"Info,omitempty" json:"Info,omitempty" yaml:"Info,omitempty"`
}

// DynPenaltiesDetailArray was auto-generated from WSDL.
type DynPenaltiesDetailArray struct {
	Items []*DynPenaltiesDetail `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynPenaltiesInfo was auto-generated from WSDL.
type DynPenaltiesInfo struct {
	CancellationCostsToday *string `xml:"CancellationCostsToday,omitempty" json:"CancellationCostsToday,omitempty" yaml:"CancellationCostsToday,omitempty"`
	NonRefundable          *string `xml:"NonRefundable,omitempty" json:"NonRefundable,omitempty" yaml:"NonRefundable,omitempty"`
}

// DynPenaltiesInfoArray was auto-generated from WSDL.
type DynPenaltiesInfoArray struct {
	Items []*DynPenaltiesInfo `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynPickupTime was auto-generated from WSDL.
type DynPickupTime struct {
	Time *string `xml:"Time,omitempty" json:"Time,omitempty" yaml:"Time,omitempty"`
}

// DynPickupTimeArray was auto-generated from WSDL.
type DynPickupTimeArray struct {
	Items []*DynPickupTime `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynPktStaticOptions was auto-generated from WSDL.
type DynPktStaticOptions struct {
	Option_ID   *string        `xml:"Option_ID,omitempty" json:"Option_ID,omitempty" yaml:"Option_ID,omitempty"`
	Option_Name *string        `xml:"Option_Name,omitempty" json:"Option_Name,omitempty" yaml:"Option_Name,omitempty"`
	Date_ini    *string        `xml:"Date_ini,omitempty" json:"Date_ini,omitempty" yaml:"Date_ini,omitempty"`
	Date_end    *string        `xml:"Date_end,omitempty" json:"Date_end,omitempty" yaml:"Date_end,omitempty"`
	Regime      *string        `xml:"Regime,omitempty" json:"Regime,omitempty" yaml:"Regime,omitempty"`
	RegimeCode  *string        `xml:"RegimeCode,omitempty" json:"RegimeCode,omitempty" yaml:"RegimeCode,omitempty"`
	DepDate     *DepDateArray  `xml:"DepDate,omitempty" json:"DepDate,omitempty" yaml:"DepDate,omitempty"`
	DepLocal    *DepLocalArray `xml:"DepLocal,omitempty" json:"DepLocal,omitempty" yaml:"DepLocal,omitempty"`
}

// DynPktStaticOptionsArray was auto-generated from WSDL.
type DynPktStaticOptionsArray struct {
	Items []*DynPktStaticOptions `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynPolicies was auto-generated from WSDL.
type DynPolicies struct {
	Type       *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	OriginType *string `xml:"OriginType,omitempty" json:"OriginType,omitempty" yaml:"OriginType,omitempty"`
	Service    *string `xml:"Service,omitempty" json:"Service,omitempty" yaml:"Service,omitempty"`
	DateFrom   *string `xml:"DateFrom,omitempty" json:"DateFrom,omitempty" yaml:"DateFrom,omitempty"`
	DateTo     *string `xml:"DateTo,omitempty" json:"DateTo,omitempty" yaml:"DateTo,omitempty"`
	ValueType  *string `xml:"ValueType,omitempty" json:"ValueType,omitempty" yaml:"ValueType,omitempty"`
	Value      *string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// DynPoliciesArray was auto-generated from WSDL.
type DynPoliciesArray struct {
	Items []*DynPolicies `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynPreBookingRequest was auto-generated from WSDL.
type DynPreBookingRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
}

// DynPreBookingResponse was auto-generated from WSDL.
type DynPreBookingResponse struct {
	Status      *string                  `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	SessionHash *string                  `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	Services    *DynBookingServicesArray `xml:"Services,omitempty" json:"Services,omitempty" yaml:"Services,omitempty"`
	Errors      *ErrorStruct             `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// DynPriceAvailResumeRequest was auto-generated from WSDL.
type DynPriceAvailResumeRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
}

// DynPriceAvailResumeResponse was auto-generated from WSDL.
type DynPriceAvailResumeResponse struct {
	Status            *string            `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	TotalGroupedPrice *string            `xml:"TotalGroupedPrice,omitempty" json:"TotalGroupedPrice,omitempty" yaml:"TotalGroupedPrice,omitempty"`
	Currency          *string            `xml:"Currency,omitempty" json:"Currency,omitempty" yaml:"Currency,omitempty"`
	MessageList       *MessagesListArray `xml:"MessageList,omitempty" json:"MessageList,omitempty" yaml:"MessageList,omitempty"`
}

// DynProdDetail was auto-generated from WSDL.
type DynProdDetail struct {
	LocalCode              *string `xml:"LocalCode,omitempty" json:"LocalCode,omitempty" yaml:"LocalCode,omitempty"`
	LocalName              *string `xml:"LocalName,omitempty" json:"LocalName,omitempty" yaml:"LocalName,omitempty"`
	Date                   *string `xml:"Date,omitempty" json:"Date,omitempty" yaml:"Date,omitempty"`
	PriceFrom              *string `xml:"PriceFrom,omitempty" json:"PriceFrom,omitempty" yaml:"PriceFrom,omitempty"`
	Nights                 *string `xml:"Nights,omitempty" json:"Nights,omitempty" yaml:"Nights,omitempty"`
	NonCommissionableTaxes *string `xml:"NonCommissionableTaxes,omitempty" json:"NonCommissionableTaxes,omitempty" yaml:"NonCommissionableTaxes,omitempty"`
	Currency               *string `xml:"Currency,omitempty" json:"Currency,omitempty" yaml:"Currency,omitempty"`
	LastUpdate             *string `xml:"LastUpdate,omitempty" json:"LastUpdate,omitempty" yaml:"LastUpdate,omitempty"`
}

// DynProdDetailArray was auto-generated from WSDL.
type DynProdDetailArray struct {
	Items []*DynProdDetail `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynProductAvailableServicesRequest was auto-generated from WSDL.
type DynProductAvailableServicesRequest struct {
	Credentials                 *CredentialsStruct     `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	UserRef                     *string                `xml:"UserRef,omitempty" json:"UserRef,omitempty" yaml:"UserRef,omitempty"`
	SessionHash                 *string                `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	ProductCode                 *string                `xml:"ProductCode,omitempty" json:"ProductCode,omitempty" yaml:"ProductCode,omitempty"`
	HtlPriceType                *string                `xml:"HtlPriceType,omitempty" json:"HtlPriceType,omitempty" yaml:"HtlPriceType,omitempty"`
	UseCache                    *string                `xml:"UseCache,omitempty" json:"UseCache,omitempty" yaml:"UseCache,omitempty"`
	HtlInternalFullAvailability *string                `xml:"HtlInternalFullAvailability,omitempty" json:"HtlInternalFullAvailability,omitempty" yaml:"HtlInternalFullAvailability,omitempty"`
	DepartureDate               *string                `xml:"DepartureDate,omitempty" json:"DepartureDate,omitempty" yaml:"DepartureDate,omitempty"`
	DepartureLocal              *string                `xml:"DepartureLocal,omitempty" json:"DepartureLocal,omitempty" yaml:"DepartureLocal,omitempty"`
	Store                       *string                `xml:"Store,omitempty" json:"Store,omitempty" yaml:"Store,omitempty"`
	RoomTypes                   *RoomTypeOptionArray   `xml:"RoomTypes,omitempty" json:"RoomTypes,omitempty" yaml:"RoomTypes,omitempty"`
	BaseLocals                  *LocalArray            `xml:"BaseLocals,omitempty" json:"BaseLocals,omitempty" yaml:"BaseLocals,omitempty"`
	ExtraNights                 *string                `xml:"ExtraNights,omitempty" json:"ExtraNights,omitempty" yaml:"ExtraNights,omitempty"`
	ExtraNightsDates            *ExtraNightsOption     `xml:"ExtraNightsDates,omitempty" json:"ExtraNightsDates,omitempty" yaml:"ExtraNightsDates,omitempty"`
	ExtraLocals                 *DynExtraLocalReqArray `xml:"ExtraLocals,omitempty" json:"ExtraLocals,omitempty" yaml:"ExtraLocals,omitempty"`
	HotelType                   *string                `xml:"HotelType,omitempty" json:"HotelType,omitempty" yaml:"HotelType,omitempty"`
	CategoryHotel               *string                `xml:"CategoryHotel,omitempty" json:"CategoryHotel,omitempty" yaml:"CategoryHotel,omitempty"`
	CodeExternalHotel           *string                `xml:"CodeExternalHotel,omitempty" json:"CodeExternalHotel,omitempty" yaml:"CodeExternalHotel,omitempty"`
	FilterW2M                   *string                `xml:"FilterW2M,omitempty" json:"FilterW2M,omitempty" yaml:"FilterW2M,omitempty"`
	CodeMotorFlights            *string                `xml:"CodeMotorFlights,omitempty" json:"CodeMotorFlights,omitempty" yaml:"CodeMotorFlights,omitempty"`
	ReserveToAdd                *string                `xml:"ReserveToAdd,omitempty" json:"ReserveToAdd,omitempty" yaml:"ReserveToAdd,omitempty"`
	ConnectorW2M                *string                `xml:"ConnectorW2M,omitempty" json:"ConnectorW2M,omitempty" yaml:"ConnectorW2M,omitempty"`
	Currency                    *string                `xml:"Currency,omitempty" json:"Currency,omitempty" yaml:"Currency,omitempty"`
}

// DynProductAvailableServicesResponse was auto-generated from
// WSDL.
type DynProductAvailableServicesResponse struct {
	SessionHash             *string                         `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	Code                    *string                         `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	CodeDefined             *string                         `xml:"CodeDefined,omitempty" json:"CodeDefined,omitempty" yaml:"CodeDefined,omitempty"`
	Name                    *string                         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	DynInsurance            *DynInsurance                   `xml:"DynInsurance,omitempty" json:"DynInsurance,omitempty" yaml:"DynInsurance,omitempty"`
	AirMandatoryServicesVal *string                         `xml:"AirMandatoryServicesVal,omitempty" json:"AirMandatoryServicesVal,omitempty" yaml:"AirMandatoryServicesVal,omitempty"`
	FlightSelected          *string                         `xml:"FlightSelected,omitempty" json:"FlightSelected,omitempty" yaml:"FlightSelected,omitempty"`
	FlightMainGroup         *DynFlightMainGroupArray        `xml:"FlightMainGroup,omitempty" json:"FlightMainGroup,omitempty" yaml:"FlightMainGroup,omitempty"`
	Itinerary               *DynItineraryArray              `xml:"Itinerary,omitempty" json:"Itinerary,omitempty" yaml:"Itinerary,omitempty"`
	OtherMandatoryServices  *DynOtherMandatoryServicesArray `xml:"OtherMandatoryServices,omitempty" json:"OtherMandatoryServices,omitempty" yaml:"OtherMandatoryServices,omitempty"`
	Errors                  *ErrorStruct                    `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// DynProductCacheRequest was auto-generated from WSDL.
type DynProductCacheRequest struct {
	Credentials  *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	ProductCode  *string            `xml:"ProductCode,omitempty" json:"ProductCode,omitempty" yaml:"ProductCode,omitempty"`
	Mode         *string            `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
	QueryOptions *DynQueryOptions   `xml:"QueryOptions,omitempty" json:"QueryOptions,omitempty" yaml:"QueryOptions,omitempty"`
	InitOptions  *DynInitOptions    `xml:"InitOptions,omitempty" json:"InitOptions,omitempty" yaml:"InitOptions,omitempty"`
}

// DynProductCacheResponse was auto-generated from WSDL.
type DynProductCacheResponse struct {
	ProductList     *DynProductListArray     `xml:"ProductList,omitempty" json:"ProductList,omitempty" yaml:"ProductList,omitempty"`
	CriteriaList    *DynCriteriaListArray    `xml:"CriteriaList,omitempty" json:"CriteriaList,omitempty" yaml:"CriteriaList,omitempty"`
	CriteriaListLog *DynCriteriaListLogArray `xml:"CriteriaListLog,omitempty" json:"CriteriaListLog,omitempty" yaml:"CriteriaListLog,omitempty"`
	Errors          *ErrorStruct             `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
	ReturnMessage   *string                  `xml:"ReturnMessage,omitempty" json:"ReturnMessage,omitempty" yaml:"ReturnMessage,omitempty"`
}

// DynProductList was auto-generated from WSDL.
type DynProductList struct {
	Code          *string             `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name          *string             `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	ProdCode      *string             `xml:"ProdCode,omitempty" json:"ProdCode,omitempty" yaml:"ProdCode,omitempty"`
	HasFligths    *string             `xml:"HasFligths,omitempty" json:"HasFligths,omitempty" yaml:"HasFligths,omitempty"`
	OperationFrom *string             `xml:"OperationFrom,omitempty" json:"OperationFrom,omitempty" yaml:"OperationFrom,omitempty"`
	OperationTo   *string             `xml:"OperationTo,omitempty" json:"OperationTo,omitempty" yaml:"OperationTo,omitempty"`
	BaseDays      *string             `xml:"BaseDays,omitempty" json:"BaseDays,omitempty" yaml:"BaseDays,omitempty"`
	PriceFrom     *string             `xml:"PriceFrom,omitempty" json:"PriceFrom,omitempty" yaml:"PriceFrom,omitempty"`
	PreCalc       *string             `xml:"PreCalc,omitempty" json:"PreCalc,omitempty" yaml:"PreCalc,omitempty"`
	ProdDetail    *DynProdDetailArray `xml:"ProdDetail,omitempty" json:"ProdDetail,omitempty" yaml:"ProdDetail,omitempty"`
}

// DynProductListArray was auto-generated from WSDL.
type DynProductListArray struct {
	Items []*DynProductList `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynProductOptionalsRequest was auto-generated from WSDL.
type DynProductOptionalsRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
}

// DynProductOptionalsResponse was auto-generated from WSDL.
type DynProductOptionalsResponse struct {
	Optionals   *DynOptionalsArray `xml:"Optionals,omitempty" json:"Optionals,omitempty" yaml:"Optionals,omitempty"`
	Errors      *ErrorStruct       `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	Currency    *string            `xml:"Currency,omitempty" json:"Currency,omitempty" yaml:"Currency,omitempty"`
}

// DynProductParametersRequest was auto-generated from WSDL.
type DynProductParametersRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Productcode *string            `xml:"Productcode,omitempty" json:"Productcode,omitempty" yaml:"Productcode,omitempty"`
}

// DynProductParametersResponse was auto-generated from WSDL.
type DynProductParametersResponse struct {
	Code            *string                 `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name            *string                 `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Localization    *string                 `xml:"Localization,omitempty" json:"Localization,omitempty" yaml:"Localization,omitempty"`
	Country         *string                 `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
	Zone            *string                 `xml:"Zone,omitempty" json:"Zone,omitempty" yaml:"Zone,omitempty"`
	ExtraNightsFrom *string                 `xml:"ExtraNightsFrom,omitempty" json:"ExtraNightsFrom,omitempty" yaml:"ExtraNightsFrom,omitempty"`
	ExtraNightsTo   *string                 `xml:"ExtraNightsTo,omitempty" json:"ExtraNightsTo,omitempty" yaml:"ExtraNightsTo,omitempty"`
	MinPaxReserve   *string                 `xml:"MinPaxReserve,omitempty" json:"MinPaxReserve,omitempty" yaml:"MinPaxReserve,omitempty"`
	MaxPaxReserve   *string                 `xml:"MaxPaxReserve,omitempty" json:"MaxPaxReserve,omitempty" yaml:"MaxPaxReserve,omitempty"`
	DepartureLocals *DynDepartureLocalArray `xml:"DepartureLocals,omitempty" json:"DepartureLocals,omitempty" yaml:"DepartureLocals,omitempty"`
	DepartureDates  *DynDepartureDateArray  `xml:"DepartureDates,omitempty" json:"DepartureDates,omitempty" yaml:"DepartureDates,omitempty"`
	RoomTypes       *DynRoomTypeArray       `xml:"RoomTypes,omitempty" json:"RoomTypes,omitempty" yaml:"RoomTypes,omitempty"`
	BaseLocals      *DynBaseLocalArray      `xml:"BaseLocals,omitempty" json:"BaseLocals,omitempty" yaml:"BaseLocals,omitempty"`
	ExtraLocals     *DynExtraLocalArray     `xml:"ExtraLocals,omitempty" json:"ExtraLocals,omitempty" yaml:"ExtraLocals,omitempty"`
	RequestAges     *DynRequestAges         `xml:"RequestAges,omitempty" json:"RequestAges,omitempty" yaml:"RequestAges,omitempty"`
	Errors          *ErrorStruct            `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// DynQueryOptions was auto-generated from WSDL.
type DynQueryOptions struct {
	DepartureLocalCode *string `xml:"DepartureLocalCode,omitempty" json:"DepartureLocalCode,omitempty" yaml:"DepartureLocalCode,omitempty"`
	Nights             *string `xml:"Nights,omitempty" json:"Nights,omitempty" yaml:"Nights,omitempty"`
}

// DynReductionNightPktStatic was auto-generated from WSDL.
type DynReductionNightPktStatic struct {
	Begin *string `xml:"begin,omitempty" json:"begin,omitempty" yaml:"begin,omitempty"`
	End   *string `xml:"End,omitempty" json:"End,omitempty" yaml:"End,omitempty"`
}

// DynReductionNightPktStaticArray was auto-generated from WSDL.
type DynReductionNightPktStaticArray struct {
	Items []*DynReductionNightPktStatic `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRemoveOptionalsRequest was auto-generated from WSDL.
type DynRemoveOptionalsRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	ItemId      *string            `xml:"ItemId,omitempty" json:"ItemId,omitempty" yaml:"ItemId,omitempty"`
	SelectionId *string            `xml:"SelectionId,omitempty" json:"SelectionId,omitempty" yaml:"SelectionId,omitempty"`
}

// DynRemoveOptionalsResponse was auto-generated from WSDL.
type DynRemoveOptionalsResponse struct {
	Errors *ErrorStruct `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// DynRequestAges was auto-generated from WSDL.
type DynRequestAges struct {
	ReqAgeType  *string `xml:"ReqAgeType,omitempty" json:"ReqAgeType,omitempty" yaml:"ReqAgeType,omitempty"`
	MinAgeAdult *string `xml:"MinAgeAdult,omitempty" json:"MinAgeAdult,omitempty" yaml:"MinAgeAdult,omitempty"`
	MaxAgeAdult *string `xml:"MaxAgeAdult,omitempty" json:"MaxAgeAdult,omitempty" yaml:"MaxAgeAdult,omitempty"`
}

// DynResCalcs was auto-generated from WSDL.
type DynResCalcs struct {
	ServiceCode       *string `xml:"ServiceCode,omitempty" json:"ServiceCode,omitempty" yaml:"ServiceCode,omitempty"`
	Service           *string `xml:"Service,omitempty" json:"Service,omitempty" yaml:"Service,omitempty"`
	Description       *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Quant             *string `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	ComissionPerc     *string `xml:"ComissionPerc,omitempty" json:"ComissionPerc,omitempty" yaml:"ComissionPerc,omitempty"`
	GrossUnitVal      *string `xml:"GrossUnitVal,omitempty" json:"GrossUnitVal,omitempty" yaml:"GrossUnitVal,omitempty"`
	GrossUnitTotalVal *string `xml:"GrossUnitTotalVal,omitempty" json:"GrossUnitTotalVal,omitempty" yaml:"GrossUnitTotalVal,omitempty"`
	GrossTotalVal     *string `xml:"GrossTotalVal,omitempty" json:"GrossTotalVal,omitempty" yaml:"GrossTotalVal,omitempty"`
	ProdId            *string `xml:"ProdId,omitempty" json:"ProdId,omitempty" yaml:"ProdId,omitempty"`
}

// DynResCalcsArray was auto-generated from WSDL.
type DynResCalcsArray struct {
	Items []*DynResCalcs `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynResPaxs was auto-generated from WSDL.
type DynResPaxs struct {
	Room             *string            `xml:"Room,omitempty" json:"Room,omitempty" yaml:"Room,omitempty"`
	Type             *string            `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	RoomNumber       *string            `xml:"RoomNumber,omitempty" json:"RoomNumber,omitempty" yaml:"RoomNumber,omitempty"`
	Age              *string            `xml:"Age,omitempty" json:"Age,omitempty" yaml:"Age,omitempty"`
	RequireBirthDate *string            `xml:"RequireBirthDate,omitempty" json:"RequireBirthDate,omitempty" yaml:"RequireBirthDate,omitempty"`
	ExtraInfo        *DynExtraInfoArray `xml:"ExtraInfo,omitempty" json:"ExtraInfo,omitempty" yaml:"ExtraInfo,omitempty"`
}

// DynResPaxsArray was auto-generated from WSDL.
type DynResPaxsArray struct {
	Items []*DynResPaxs `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynResRemarks was auto-generated from WSDL.
type DynResRemarks struct {
	Type        *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Interface   *string `xml:"Interface,omitempty" json:"Interface,omitempty" yaml:"Interface,omitempty"`
	Text        *string `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`
	RelatedUnit *string `xml:"RelatedUnit,omitempty" json:"RelatedUnit,omitempty" yaml:"RelatedUnit,omitempty"`
}

// DynResRemarksArray was auto-generated from WSDL.
type DynResRemarksArray struct {
	Items []*DynResRemarks `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynResServices was auto-generated from WSDL.
type DynResServices struct {
	Type        *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	SubType     *string `xml:"SubType,omitempty" json:"SubType,omitempty" yaml:"SubType,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Quant       *string `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	DateFrom    *string `xml:"DateFrom,omitempty" json:"DateFrom,omitempty" yaml:"DateFrom,omitempty"`
	DateTo      *string `xml:"DateTo,omitempty" json:"DateTo,omitempty" yaml:"DateTo,omitempty"`
	Status      *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// DynResServicesArray was auto-generated from WSDL.
type DynResServicesArray struct {
	Items []*DynResServices `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRestriction was auto-generated from WSDL.
type DynRestriction struct {
	Type     *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Text     *string `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`
	MinNight *string `xml:"MinNight,omitempty" json:"MinNight,omitempty" yaml:"MinNight,omitempty"`
	DateFrom *string `xml:"DateFrom,omitempty" json:"DateFrom,omitempty" yaml:"DateFrom,omitempty"`
	DateTo   *string `xml:"DateTo,omitempty" json:"DateTo,omitempty" yaml:"DateTo,omitempty"`
}

// DynRestrictionArray was auto-generated from WSDL.
type DynRestrictionArray struct {
	Items []*DynRestriction `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynResumeCalcs was auto-generated from WSDL.
type DynResumeCalcs struct {
	TotalCommisionable         *string `xml:"TotalCommisionable,omitempty" json:"TotalCommisionable,omitempty" yaml:"TotalCommisionable,omitempty"`
	TotalNoCommisionable       *string `xml:"TotalNoCommisionable,omitempty" json:"TotalNoCommisionable,omitempty" yaml:"TotalNoCommisionable,omitempty"`
	TotalCommision             *string `xml:"TotalCommision,omitempty" json:"TotalCommision,omitempty" yaml:"TotalCommision,omitempty"`
	TotalVatCommision          *string `xml:"TotalVatCommision,omitempty" json:"TotalVatCommision,omitempty" yaml:"TotalVatCommision,omitempty"`
	TotalInvoice               *string `xml:"TotalInvoice,omitempty" json:"TotalInvoice,omitempty" yaml:"TotalInvoice,omitempty"`
	TotalPrice                 *string `xml:"TotalPrice,omitempty" json:"TotalPrice,omitempty" yaml:"TotalPrice,omitempty"`
	TotalPriceApplyDiscount    *string `xml:"TotalPriceApplyDiscount,omitempty" json:"TotalPriceApplyDiscount,omitempty" yaml:"TotalPriceApplyDiscount,omitempty"`
	TotalPriceNotApplyDiscount *string `xml:"TotalPriceNotApplyDiscount,omitempty" json:"TotalPriceNotApplyDiscount,omitempty" yaml:"TotalPriceNotApplyDiscount,omitempty"`
	Discount                   *string `xml:"Discount,omitempty" json:"Discount,omitempty" yaml:"Discount,omitempty"`
	Total                      *string `xml:"Total,omitempty" json:"Total,omitempty" yaml:"Total,omitempty"`
}

// DynResumeCalcsArray was auto-generated from WSDL.
type DynResumeCalcsArray struct {
	Items []*DynResumeCalcs `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRoomOptions was auto-generated from WSDL.
type DynRoomOptions struct {
	RoomTypeCode *string `xml:"RoomTypeCode,omitempty" json:"RoomTypeCode,omitempty" yaml:"RoomTypeCode,omitempty"`
	Name         *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	TotalGuest   *string `xml:"TotalGuest,omitempty" json:"TotalGuest,omitempty" yaml:"TotalGuest,omitempty"`
	MaxSeniors   *string `xml:"MaxSeniors,omitempty" json:"MaxSeniors,omitempty" yaml:"MaxSeniors,omitempty"`
}

// DynRoomOptionsArray was auto-generated from WSDL.
type DynRoomOptionsArray struct {
	Items []*DynRoomOptions `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRoomType was auto-generated from WSDL.
type DynRoomType struct {
	Code         *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name         *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	NumAdults    *string `xml:"NumAdults,omitempty" json:"NumAdults,omitempty" yaml:"NumAdults,omitempty"`
	NumChilds    *string `xml:"NumChilds,omitempty" json:"NumChilds,omitempty" yaml:"NumChilds,omitempty"`
	ChildAgeFrom *string `xml:"ChildAgeFrom,omitempty" json:"ChildAgeFrom,omitempty" yaml:"ChildAgeFrom,omitempty"`
	ChildAgeTo   *string `xml:"ChildAgeTo,omitempty" json:"ChildAgeTo,omitempty" yaml:"ChildAgeTo,omitempty"`
}

// DynRoomTypeArray was auto-generated from WSDL.
type DynRoomTypeArray struct {
	Items []*DynRoomType `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRooms was auto-generated from WSDL.
type DynRooms struct {
	Code             *string            `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name             *string            `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	NumAdults        *string            `xml:"NumAdults,omitempty" json:"NumAdults,omitempty" yaml:"NumAdults,omitempty"`
	NumChilds        *string            `xml:"NumChilds,omitempty" json:"NumChilds,omitempty" yaml:"NumChilds,omitempty"`
	ChildsAgeFrom    *string            `xml:"ChildsAgeFrom,omitempty" json:"ChildsAgeFrom,omitempty" yaml:"ChildsAgeFrom,omitempty"`
	GuestTypeMax     *string            `xml:"GuestTypeMax,omitempty" json:"GuestTypeMax,omitempty" yaml:"GuestTypeMax,omitempty"`
	GuestTypeMin     *string            `xml:"GuestTypeMin,omitempty" json:"GuestTypeMin,omitempty" yaml:"GuestTypeMin,omitempty"`
	TotalGuest       *string            `xml:"TotalGuest,omitempty" json:"TotalGuest,omitempty" yaml:"TotalGuest,omitempty"`
	MaxSeniors       *string            `xml:"MaxSeniors,omitempty" json:"MaxSeniors,omitempty" yaml:"MaxSeniors,omitempty"`
	ChildsAgeTo      *string            `xml:"ChildsAgeTo,omitempty" json:"ChildsAgeTo,omitempty" yaml:"ChildsAgeTo,omitempty"`
	AvailStatus      *string            `xml:"AvailStatus,omitempty" json:"AvailStatus,omitempty" yaml:"AvailStatus,omitempty"`
	RPH              *string            `xml:"RPH,omitempty" json:"RPH,omitempty" yaml:"RPH,omitempty"`
	Provider         *string            `xml:"Provider,omitempty" json:"Provider,omitempty" yaml:"Provider,omitempty"`
	RoomNum          *string            `xml:"RoomNum,omitempty" json:"RoomNum,omitempty" yaml:"RoomNum,omitempty"`
	RoomMatch        *string            `xml:"RoomMatch,omitempty" json:"RoomMatch,omitempty" yaml:"RoomMatch,omitempty"`
	UpgradeSupVal    *string            `xml:"UpgradeSupVal,omitempty" json:"UpgradeSupVal,omitempty" yaml:"UpgradeSupVal,omitempty"`
	BoardIncluded    *string            `xml:"BoardIncluded,omitempty" json:"BoardIncluded,omitempty" yaml:"BoardIncluded,omitempty"`
	BoardDescription *string            `xml:"BoardDescription,omitempty" json:"BoardDescription,omitempty" yaml:"BoardDescription,omitempty"`
	BoardCode        *string            `xml:"BoardCode,omitempty" json:"BoardCode,omitempty" yaml:"BoardCode,omitempty"`
	SellValue        *string            `xml:"SellValue,omitempty" json:"SellValue,omitempty" yaml:"SellValue,omitempty"`
	NonRefundable    *string            `xml:"NonRefundable,omitempty" json:"NonRefundable,omitempty" yaml:"NonRefundable,omitempty"`
	RoomsList        *RoomsListArray    `xml:"RoomsList,omitempty" json:"RoomsList,omitempty" yaml:"RoomsList,omitempty"`
	ElementsList     *ElementsListArray `xml:"ElementsList,omitempty" json:"ElementsList,omitempty" yaml:"ElementsList,omitempty"`
}

// DynRoomsArray was auto-generated from WSDL.
type DynRoomsArray struct {
	Items []*DynRooms `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRoomsOccupancy was auto-generated from WSDL.
type DynRoomsOccupancy struct {
	OriginalRoomId *string        `xml:"OriginalRoomId,omitempty" json:"OriginalRoomId,omitempty" yaml:"OriginalRoomId,omitempty"`
	Type           *string        `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Quant          *string        `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	RoomDesc       *string        `xml:"RoomDesc,omitempty" json:"RoomDesc,omitempty" yaml:"RoomDesc,omitempty"`
	RoomGroup      *string        `xml:"RoomGroup,omitempty" json:"RoomGroup,omitempty" yaml:"RoomGroup,omitempty"`
	NumAdults      *string        `xml:"NumAdults,omitempty" json:"NumAdults,omitempty" yaml:"NumAdults,omitempty"`
	NumChilds      *string        `xml:"NumChilds,omitempty" json:"NumChilds,omitempty" yaml:"NumChilds,omitempty"`
	Rooms          *DynRoomsArray `xml:"Rooms,omitempty" json:"Rooms,omitempty" yaml:"Rooms,omitempty"`
}

// DynRoomsOccupancyArray was auto-generated from WSDL.
type DynRoomsOccupancyArray struct {
	Items []*DynRoomsOccupancy `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRoomsOfi was auto-generated from WSDL.
type DynRoomsOfi struct {
	Code  *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Value *string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// DynRoomsOfiArray was auto-generated from WSDL.
type DynRoomsOfiArray struct {
	Items []*DynRoomsOfi `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRoomsPenalties was auto-generated from WSDL.
type DynRoomsPenalties struct {
	Name             *string                  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Code             *string                  `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	PenaltiesDetails *DynPenaltiesDetailArray `xml:"PenaltiesDetails,omitempty" json:"PenaltiesDetails,omitempty" yaml:"PenaltiesDetails,omitempty"`
	CancelPenalties  *DynCancelPenaltiesArray `xml:"CancelPenalties,omitempty" json:"CancelPenalties,omitempty" yaml:"CancelPenalties,omitempty"`
}

// DynRoomsPenaltiesArray was auto-generated from WSDL.
type DynRoomsPenaltiesArray struct {
	Items []*DynRoomsPenalties `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRoomsPktStatic was auto-generated from WSDL.
type DynRoomsPktStatic struct {
	Code  *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Value *string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// DynRoomsPktStaticArray was auto-generated from WSDL.
type DynRoomsPktStaticArray struct {
	Items []*DynRoomsPktStatic `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRoomsSelect was auto-generated from WSDL.
type DynRoomsSelect struct {
	RoomNum   *string `xml:"RoomNum,omitempty" json:"RoomNum,omitempty" yaml:"RoomNum,omitempty"`
	RoomMatch *string `xml:"RoomMatch,omitempty" json:"RoomMatch,omitempty" yaml:"RoomMatch,omitempty"`
	RoomCode  *string `xml:"RoomCode,omitempty" json:"RoomCode,omitempty" yaml:"RoomCode,omitempty"`
}

// DynRoomsSelectArray was auto-generated from WSDL.
type DynRoomsSelectArray struct {
	Items []*DynRoomsSelect `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRoomsSeniorOfi was auto-generated from WSDL.
type DynRoomsSeniorOfi struct {
	Code  *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Value *string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// DynRoomsSeniorOfiArray was auto-generated from WSDL.
type DynRoomsSeniorOfiArray struct {
	Items []*DynRoomsSeniorOfi `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynRoomsSeniorPktStatic was auto-generated from WSDL.
type DynRoomsSeniorPktStatic struct {
	Code  *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Value *string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// DynRoomsSeniorPktStaticArray was auto-generated from WSDL.
type DynRoomsSeniorPktStaticArray struct {
	Items []*DynRoomsSeniorPktStatic `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynSelOptionals was auto-generated from WSDL.
type DynSelOptionals struct {
	Code        *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	ItemSel     *string `xml:"ItemSel,omitempty" json:"ItemSel,omitempty" yaml:"ItemSel,omitempty"`
	IdSel       *string `xml:"IdSel,omitempty" json:"IdSel,omitempty" yaml:"IdSel,omitempty"`
	DaysApplied *string `xml:"DaysApplied,omitempty" json:"DaysApplied,omitempty" yaml:"DaysApplied,omitempty"`
	ProductName *string `xml:"ProductName,omitempty" json:"ProductName,omitempty" yaml:"ProductName,omitempty"`
	DataSel     *string `xml:"DataSel,omitempty" json:"DataSel,omitempty" yaml:"DataSel,omitempty"`
	Adults      *string `xml:"Adults,omitempty" json:"Adults,omitempty" yaml:"Adults,omitempty"`
	Childs      *string `xml:"Childs,omitempty" json:"Childs,omitempty" yaml:"Childs,omitempty"`
	PickDrop    *string `xml:"PickDrop,omitempty" json:"PickDrop,omitempty" yaml:"PickDrop,omitempty"`
	Mandatory   *string `xml:"Mandatory,omitempty" json:"Mandatory,omitempty" yaml:"Mandatory,omitempty"`
}

// DynSelOptionalsArray was auto-generated from WSDL.
type DynSelOptionalsArray struct {
	Items []*DynSelOptionals `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynSelectedRoomsOfi was auto-generated from WSDL.
type DynSelectedRoomsOfi struct {
	RoomCode  *string              `xml:"RoomCode,omitempty" json:"RoomCode,omitempty" yaml:"RoomCode,omitempty"`
	RoomDesc  *string              `xml:"RoomDesc,omitempty" json:"RoomDesc,omitempty" yaml:"RoomDesc,omitempty"`
	RoomType  *string              `xml:"RoomType,omitempty" json:"RoomType,omitempty" yaml:"RoomType,omitempty"`
	NumAdults *string              `xml:"NumAdults,omitempty" json:"NumAdults,omitempty" yaml:"NumAdults,omitempty"`
	NumChild  *string              `xml:"NumChild,omitempty" json:"NumChild,omitempty" yaml:"NumChild,omitempty"`
	ChildAges *string              `xml:"ChildAges,omitempty" json:"ChildAges,omitempty" yaml:"ChildAges,omitempty"`
	DynRooms  *DynRoomOptionsArray `xml:"DynRooms,omitempty" json:"DynRooms,omitempty" yaml:"DynRooms,omitempty"`
}

// DynSelectedRoomsOfiArray was auto-generated from WSDL.
type DynSelectedRoomsOfiArray struct {
	Items []*DynSelectedRoomsOfi `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynSelectedRoomsPktStatic was auto-generated from WSDL.
type DynSelectedRoomsPktStatic struct {
	RoomCode      *string `xml:"RoomCode,omitempty" json:"RoomCode,omitempty" yaml:"RoomCode,omitempty"`
	RoomDesc      *string `xml:"RoomDesc,omitempty" json:"RoomDesc,omitempty" yaml:"RoomDesc,omitempty"`
	ProdType      *string `xml:"ProdType,omitempty" json:"ProdType,omitempty" yaml:"ProdType,omitempty"`
	RoomType      *string `xml:"RoomType,omitempty" json:"RoomType,omitempty" yaml:"RoomType,omitempty"`
	NumAdults     *string `xml:"NumAdults,omitempty" json:"NumAdults,omitempty" yaml:"NumAdults,omitempty"`
	Price         *string `xml:"Price,omitempty" json:"Price,omitempty" yaml:"Price,omitempty"`
	NumChild      *string `xml:"NumChild,omitempty" json:"NumChild,omitempty" yaml:"NumChild,omitempty"`
	AvailStatus   *string `xml:"AvailStatus,omitempty" json:"AvailStatus,omitempty" yaml:"AvailStatus,omitempty"`
	ChildsAgeFrom *string `xml:"ChildsAgeFrom,omitempty" json:"ChildsAgeFrom,omitempty" yaml:"ChildsAgeFrom,omitempty"`
	ChildsAgeto   *string `xml:"ChildsAgeto,omitempty" json:"ChildsAgeto,omitempty" yaml:"ChildsAgeto,omitempty"`
}

// DynSelectedRoomsPktStaticArray was auto-generated from WSDL.
type DynSelectedRoomsPktStaticArray struct {
	Items []*DynSelectedRoomsPktStatic `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynSeniorPktStatic was auto-generated from WSDL.
type DynSeniorPktStatic struct {
	Min *string `xml:"Min,omitempty" json:"Min,omitempty" yaml:"Min,omitempty"`
	Max *string `xml:"Max,omitempty" json:"Max,omitempty" yaml:"Max,omitempty"`
}

// DynSeniorPktStaticArray was auto-generated from WSDL.
type DynSeniorPktStaticArray struct {
	Items []*DynSeniorPktStatic `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynServicesSelectedRequest was auto-generated from WSDL.
type DynServicesSelectedRequest struct {
	Credentials                         *CredentialsStruct            `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash                         *string                       `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	Reserve                             *string                       `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	Local                               *string                       `xml:"Local,omitempty" json:"Local,omitempty" yaml:"Local,omitempty"`
	Regime                              *string                       `xml:"Regime,omitempty" json:"Regime,omitempty" yaml:"Regime,omitempty"`
	Connector                           *string                       `xml:"Connector,omitempty" json:"Connector,omitempty" yaml:"Connector,omitempty"`
	Skeleton_id                         *string                       `xml:"Skeleton_id,omitempty" json:"Skeleton_id,omitempty" yaml:"Skeleton_id,omitempty"`
	Action                              *string                       `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	DepLocal                            *string                       `xml:"DepLocal,omitempty" json:"DepLocal,omitempty" yaml:"DepLocal,omitempty"`
	OptionId                            *string                       `xml:"OptionId,omitempty" json:"OptionId,omitempty" yaml:"OptionId,omitempty"`
	FlightsSelected                     *DynFlightsSelectedArray      `xml:"FlightsSelected,omitempty" json:"FlightsSelected,omitempty" yaml:"FlightsSelected,omitempty"`
	FlightsSelectedSuperBB              *DynFlightsSelectedArray      `xml:"FlightsSelectedSuperBB,omitempty" json:"FlightsSelectedSuperBB,omitempty" yaml:"FlightsSelectedSuperBB,omitempty"`
	FlightsSelectedFlightEngine         *DynFlightsSelectedArray      `xml:"FlightsSelectedFlightEngine,omitempty" json:"FlightsSelectedFlightEngine,omitempty" yaml:"FlightsSelectedFlightEngine,omitempty"`
	FlightsSelectedAirAvailFlightEngine *DynFlightsSelectedArray      `xml:"FlightsSelectedAirAvailFlightEngine,omitempty" json:"FlightsSelectedAirAvailFlightEngine,omitempty" yaml:"FlightsSelectedAirAvailFlightEngine,omitempty"`
	HotelsSelected                      *DynHotelsSelectedArray       `xml:"HotelsSelected,omitempty" json:"HotelsSelected,omitempty" yaml:"HotelsSelected,omitempty"`
	OptionalsSelected                   *DynOptionalsSelectedArray    `xml:"OptionalsSelected,omitempty" json:"OptionalsSelected,omitempty" yaml:"OptionalsSelected,omitempty"`
	RoomsOfi                            *DynRoomsOfiArray             `xml:"RoomsOfi,omitempty" json:"RoomsOfi,omitempty" yaml:"RoomsOfi,omitempty"`
	RoomsSeniorOfi                      *DynRoomsSeniorOfiArray       `xml:"RoomsSeniorOfi,omitempty" json:"RoomsSeniorOfi,omitempty" yaml:"RoomsSeniorOfi,omitempty"`
	RoomsPktStatic                      *DynRoomsPktStaticArray       `xml:"RoomsPktStatic,omitempty" json:"RoomsPktStatic,omitempty" yaml:"RoomsPktStatic,omitempty"`
	RoomsSeniorPktStatic                *DynRoomsSeniorPktStaticArray `xml:"RoomsSeniorPktStatic,omitempty" json:"RoomsSeniorPktStatic,omitempty" yaml:"RoomsSeniorPktStatic,omitempty"`
}

// DynServicesSelectedResponse was auto-generated from WSDL.
type DynServicesSelectedResponse struct {
	Errors *ErrorStruct `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// DynStops was auto-generated from WSDL.
type DynStops struct {
	Code        *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	IdRoute     *string `xml:"IdRoute,omitempty" json:"IdRoute,omitempty" yaml:"IdRoute,omitempty"`
}

// DynStopsArray was auto-generated from WSDL.
type DynStopsArray struct {
	Items []*DynStops `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynSuplements was auto-generated from WSDL.
type DynSuplements struct {
	Index        *string `xml:"Index,omitempty" json:"Index,omitempty" yaml:"Index,omitempty"`
	Code         *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description  *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Value        *string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
	ValueType    *string `xml:"ValueType,omitempty" json:"ValueType,omitempty" yaml:"ValueType,omitempty"`
	BoardCode    *string `xml:"BoardCode,omitempty" json:"BoardCode,omitempty" yaml:"BoardCode,omitempty"`
	RoomCode     *string `xml:"RoomCode,omitempty" json:"RoomCode,omitempty" yaml:"RoomCode,omitempty"`
	BaseRoomCode *string `xml:"BaseRoomCode,omitempty" json:"BaseRoomCode,omitempty" yaml:"BaseRoomCode,omitempty"`
}

// DynSuplementsArray was auto-generated from WSDL.
type DynSuplementsArray struct {
	Items []*DynSuplements `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynUnprocessedDatesList was auto-generated from WSDL.
type DynUnprocessedDatesList struct {
	Date       *string `xml:"Date,omitempty" json:"Date,omitempty" yaml:"Date,omitempty"`
	Origin     *string `xml:"Origin,omitempty" json:"Origin,omitempty" yaml:"Origin,omitempty"`
	LastUpdate *string `xml:"LastUpdate,omitempty" json:"LastUpdate,omitempty" yaml:"LastUpdate,omitempty"`
}

// DynUnprocessedDatesListArray was auto-generated from WSDL.
type DynUnprocessedDatesListArray struct {
	Items []*DynUnprocessedDatesList `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DynXtraPaxInfo was auto-generated from WSDL.
type DynXtraPaxInfo struct {
	XtraPaxInfoId    *string `xml:"XtraPaxInfoId,omitempty" json:"XtraPaxInfoId,omitempty" yaml:"XtraPaxInfoId,omitempty"`
	XtraPaxInfoValue *string `xml:"XtraPaxInfoValue,omitempty" json:"XtraPaxInfoValue,omitempty" yaml:"XtraPaxInfoValue,omitempty"`
}

// DynXtraPaxInfoArray was auto-generated from WSDL.
type DynXtraPaxInfoArray struct {
	Items []*DynXtraPaxInfo `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ElementsList was auto-generated from WSDL.
type ElementsList struct {
	Class         *string `xml:"Class,omitempty" json:"Class,omitempty" yaml:"Class,omitempty"`
	Type          *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Price         *string `xml:"Price,omitempty" json:"Price,omitempty" yaml:"Price,omitempty"`
	OnlyResidents *string `xml:"OnlyResidents,omitempty" json:"OnlyResidents,omitempty" yaml:"OnlyResidents,omitempty"`
	Name          *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description   *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

// ElementsListArray was auto-generated from WSDL.
type ElementsListArray struct {
	Items []*ElementsList `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Error was auto-generated from WSDL.
type Error struct {
	Code *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Desc *string `xml:"Desc,omitempty" json:"Desc,omitempty" yaml:"Desc,omitempty"`
}

// ErrorArray was auto-generated from WSDL.
type ErrorArray struct {
	Items []*Error `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ErrorStruct was auto-generated from WSDL.
type ErrorStruct struct {
	HasErrors *string     `xml:"HasErrors,omitempty" json:"HasErrors,omitempty" yaml:"HasErrors,omitempty"`
	ErrorList *ErrorArray `xml:"ErrorList,omitempty" json:"ErrorList,omitempty" yaml:"ErrorList,omitempty"`
}

// Errors was auto-generated from WSDL.
type Errors struct {
	Error *string `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
}

// ErrorsArray was auto-generated from WSDL.
type ErrorsArray struct {
	Items []*Errors `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ErrorsMsg was auto-generated from WSDL.
type ErrorsMsg struct {
	Error *string `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
}

// ErrorsMsgArray was auto-generated from WSDL.
type ErrorsMsgArray struct {
	Items []*ErrorsMsg `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Extra was auto-generated from WSDL.
type Extra struct {
	AttrId    *string `xml:"AttrId,omitempty" json:"AttrId,omitempty" yaml:"AttrId,omitempty"`
	AttrName  *string `xml:"AttrName,omitempty" json:"AttrName,omitempty" yaml:"AttrName,omitempty"`
	AttrValue *string `xml:"AttrValue,omitempty" json:"AttrValue,omitempty" yaml:"AttrValue,omitempty"`
	AttrCode  *string `xml:"AttrCode,omitempty" json:"AttrCode,omitempty" yaml:"AttrCode,omitempty"`
}

// ExtraNightsOption was auto-generated from WSDL.
type ExtraNightsOption struct {
	Date_diff_nights_start *string `xml:"date_diff_nights_start,omitempty" json:"date_diff_nights_start,omitempty" yaml:"date_diff_nights_start,omitempty"`
	Date_diff_nights_end   *string `xml:"date_diff_nights_end,omitempty" json:"date_diff_nights_end,omitempty" yaml:"date_diff_nights_end,omitempty"`
}

// ExtraNightsOptionArray was auto-generated from WSDL.
type ExtraNightsOptionArray struct {
	Items []*ExtraNightsOption `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// FaturaStruct was auto-generated from WSDL.
type FaturaStruct struct {
	Nome       *string `xml:"Nome,omitempty" json:"Nome,omitempty" yaml:"Nome,omitempty"`
	NIF        *string `xml:"NIF,omitempty" json:"NIF,omitempty" yaml:"NIF,omitempty"`
	Contato    *string `xml:"Contato,omitempty" json:"Contato,omitempty" yaml:"Contato,omitempty"`
	Email      *string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`
	Morada     *string `xml:"Morada,omitempty" json:"Morada,omitempty" yaml:"Morada,omitempty"`
	CPost      *string `xml:"CPost,omitempty" json:"CPost,omitempty" yaml:"CPost,omitempty"`
	Localidade *string `xml:"Localidade,omitempty" json:"Localidade,omitempty" yaml:"Localidade,omitempty"`
}

// FlightIn was auto-generated from WSDL.
type FlightIn struct {
	Transfer       *string `xml:"Transfer,omitempty" json:"Transfer,omitempty" yaml:"Transfer,omitempty"`
	TypeTransporte *string `xml:"TypeTransporte,omitempty" json:"TypeTransporte,omitempty" yaml:"TypeTransporte,omitempty"`
	DateArrived    *string `xml:"DateArrived,omitempty" json:"DateArrived,omitempty" yaml:"DateArrived,omitempty"`
	Airport        *string `xml:"Airport,omitempty" json:"Airport,omitempty" yaml:"Airport,omitempty"`
	Voo            *string `xml:"Voo,omitempty" json:"Voo,omitempty" yaml:"Voo,omitempty"`
	NumVoo         *string `xml:"NumVoo,omitempty" json:"NumVoo,omitempty" yaml:"NumVoo,omitempty"`
	ExitTime       *string `xml:"ExitTime,omitempty" json:"ExitTime,omitempty" yaml:"ExitTime,omitempty"`
	ArrivalTime    *string `xml:"ArrivalTime,omitempty" json:"ArrivalTime,omitempty" yaml:"ArrivalTime,omitempty"`
	Obs            *string `xml:"Obs,omitempty" json:"Obs,omitempty" yaml:"Obs,omitempty"`
}

// FlightInArray was auto-generated from WSDL.
type FlightInArray struct {
	Items []*FlightIn `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// FlightOut was auto-generated from WSDL.
type FlightOut struct {
	Transfer       *string `xml:"Transfer,omitempty" json:"Transfer,omitempty" yaml:"Transfer,omitempty"`
	TypeTransporte *string `xml:"TypeTransporte,omitempty" json:"TypeTransporte,omitempty" yaml:"TypeTransporte,omitempty"`
	DateDeparture  *string `xml:"DateDeparture,omitempty" json:"DateDeparture,omitempty" yaml:"DateDeparture,omitempty"`
	Airport        *string `xml:"Airport,omitempty" json:"Airport,omitempty" yaml:"Airport,omitempty"`
	Voo            *string `xml:"Voo,omitempty" json:"Voo,omitempty" yaml:"Voo,omitempty"`
	NumVoo         *string `xml:"NumVoo,omitempty" json:"NumVoo,omitempty" yaml:"NumVoo,omitempty"`
	ExitTime       *string `xml:"ExitTime,omitempty" json:"ExitTime,omitempty" yaml:"ExitTime,omitempty"`
	ArrivalTime    *string `xml:"ArrivalTime,omitempty" json:"ArrivalTime,omitempty" yaml:"ArrivalTime,omitempty"`
	Obs            *string `xml:"Obs,omitempty" json:"Obs,omitempty" yaml:"Obs,omitempty"`
}

// FlightOutArray was auto-generated from WSDL.
type FlightOutArray struct {
	Items []*FlightOut `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// GeneralConditions was auto-generated from WSDL.
type GeneralConditions struct {
	Type    *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Content *string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

// GeneralConditionsArray was auto-generated from WSDL.
type GeneralConditionsArray struct {
	Items []*GeneralConditions `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// GetClientsRequest was auto-generated from WSDL.
type GetClientsRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
}

// GetClientsResponse was auto-generated from WSDL.
type GetClientsResponse struct {
	Clients *ClientsArray `xml:"Clients,omitempty" json:"Clients,omitempty" yaml:"Clients,omitempty"`
}

// GetHotelDetailsRequest was auto-generated from WSDL.
type GetHotelDetailsRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	HotelCode   *string            `xml:"HotelCode,omitempty" json:"HotelCode,omitempty" yaml:"HotelCode,omitempty"`
}

// GetHotelDetailsResponse was auto-generated from WSDL.
type GetHotelDetailsResponse struct {
	SerializedTecData *string          `xml:"SerializedTecData,omitempty" json:"SerializedTecData,omitempty" yaml:"SerializedTecData,omitempty"`
	TextList          *HotelTextArray  `xml:"TextList,omitempty" json:"TextList,omitempty" yaml:"TextList,omitempty"`
	ImageList         *HotelImageArray `xml:"ImageList,omitempty" json:"ImageList,omitempty" yaml:"ImageList,omitempty"`
	Errors            *ErrorStruct     `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// GetPaxAdditionalInfoRequest was auto-generated from WSDL.
type GetPaxAdditionalInfoRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
}

// GetPaxAdditionalInfoResponse was auto-generated from WSDL.
type GetPaxAdditionalInfoResponse struct {
	Reserve *string       `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	PaxInfo *PaxInfoArray `xml:"PaxInfo,omitempty" json:"PaxInfo,omitempty" yaml:"PaxInfo,omitempty"`
}

// GetPaxInfoRequest was auto-generated from WSDL.
type GetPaxInfoRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	ProdId      *string            `xml:"ProdId,omitempty" json:"ProdId,omitempty" yaml:"ProdId,omitempty"`
}

// GetPaxInfoResponse was auto-generated from WSDL.
type GetPaxInfoResponse struct {
	PaxInfo *InfoArray `xml:"PaxInfo,omitempty" json:"PaxInfo,omitempty" yaml:"PaxInfo,omitempty"`
}

// GetStoresRequest was auto-generated from WSDL.
type GetStoresRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
}

// GetStoresResponse was auto-generated from WSDL.
type GetStoresResponse struct {
	Stores *StoresArray `xml:"Stores,omitempty" json:"Stores,omitempty" yaml:"Stores,omitempty"`
}

// HotelCalcs was auto-generated from WSDL.
type HotelCalcs struct {
	ServiceCode       *string `xml:"ServiceCode,omitempty" json:"ServiceCode,omitempty" yaml:"ServiceCode,omitempty"`
	Description       *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Quant             *string `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	ComissionPerc     *string `xml:"ComissionPerc,omitempty" json:"ComissionPerc,omitempty" yaml:"ComissionPerc,omitempty"`
	BuyTotalVal       *string `xml:"BuyTotalVal,omitempty" json:"BuyTotalVal,omitempty" yaml:"BuyTotalVal,omitempty"`
	BuyUnitTotalVal   *string `xml:"BuyUnitTotalVal,omitempty" json:"BuyUnitTotalVal,omitempty" yaml:"BuyUnitTotalVal,omitempty"`
	GrossUnitTotalVal *string `xml:"GrossUnitTotalVal,omitempty" json:"GrossUnitTotalVal,omitempty" yaml:"GrossUnitTotalVal,omitempty"`
	GrossTotalVal     *string `xml:"GrossTotalVal,omitempty" json:"GrossTotalVal,omitempty" yaml:"GrossTotalVal,omitempty"`
}

// HotelCalcsArray was auto-generated from WSDL.
type HotelCalcsArray struct {
	Items []*HotelCalcs `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// HotelImage was auto-generated from WSDL.
type HotelImage struct {
	ImageUrl *string `xml:"ImageUrl,omitempty" json:"ImageUrl,omitempty" yaml:"ImageUrl,omitempty"`
}

// HotelImageArray was auto-generated from WSDL.
type HotelImageArray struct {
	Items []*HotelImage `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// HotelList was auto-generated from WSDL.
type HotelList struct {
	Id                *string               `xml:"Id,omitempty" json:"Id,omitempty" yaml:"Id,omitempty"`
	Name              *string               `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Rating            *string               `xml:"Rating,omitempty" json:"Rating,omitempty" yaml:"Rating,omitempty"`
	Registry          *string               `xml:"Registry,omitempty" json:"Registry,omitempty" yaml:"Registry,omitempty"`
	ImageUrl          *string               `xml:"ImageUrl,omitempty" json:"ImageUrl,omitempty" yaml:"ImageUrl,omitempty"`
	Address           *string               `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	PostalCode        *string               `xml:"PostalCode,omitempty" json:"PostalCode,omitempty" yaml:"PostalCode,omitempty"`
	Country           *string               `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
	CountryName       *string               `xml:"CountryName,omitempty" json:"CountryName,omitempty" yaml:"CountryName,omitempty"`
	Local             *string               `xml:"Local,omitempty" json:"Local,omitempty" yaml:"Local,omitempty"`
	LocalName         *string               `xml:"LocalName,omitempty" json:"LocalName,omitempty" yaml:"LocalName,omitempty"`
	Zone              *string               `xml:"Zone,omitempty" json:"Zone,omitempty" yaml:"Zone,omitempty"`
	ZoneName          *string               `xml:"ZoneName,omitempty" json:"ZoneName,omitempty" yaml:"ZoneName,omitempty"`
	PhoneNumber       *string               `xml:"PhoneNumber,omitempty" json:"PhoneNumber,omitempty" yaml:"PhoneNumber,omitempty"`
	GpsLat            *string               `xml:"GpsLat,omitempty" json:"GpsLat,omitempty" yaml:"GpsLat,omitempty"`
	GpsLon            *string               `xml:"GpsLon,omitempty" json:"GpsLon,omitempty" yaml:"GpsLon,omitempty"`
	PortalObs         *string               `xml:"PortalObs,omitempty" json:"PortalObs,omitempty" yaml:"PortalObs,omitempty"`
	SerializedTecData *string               `xml:"SerializedTecData,omitempty" json:"SerializedTecData,omitempty" yaml:"SerializedTecData,omitempty"`
	Rooms             *RoomsHtlArray        `xml:"Rooms,omitempty" json:"Rooms,omitempty" yaml:"Rooms,omitempty"`
	Boards            *BoardsHtlArray       `xml:"Boards,omitempty" json:"Boards,omitempty" yaml:"Boards,omitempty"`
	Restrictions      *RestrictionsHtlArray `xml:"Restrictions,omitempty" json:"Restrictions,omitempty" yaml:"Restrictions,omitempty"`
	Observations      *ObservationsHtlArray `xml:"Observations,omitempty" json:"Observations,omitempty" yaml:"Observations,omitempty"`
}

// HotelListArray was auto-generated from WSDL.
type HotelListArray struct {
	Items []*HotelList `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// HotelListRequest was auto-generated from WSDL.
type HotelListRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
}

// HotelListResponse was auto-generated from WSDL.
type HotelListResponse struct {
	Errors         *ErrorStruct `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
	SessionHash    *string      `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	HotelBasicList *string      `xml:"HotelBasicList,omitempty" json:"HotelBasicList,omitempty" yaml:"HotelBasicList,omitempty"`
}

// HotelMasterAttributesRequest was auto-generated from WSDL.
type HotelMasterAttributesRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
}

// HotelMasterAttributesResponse was auto-generated from WSDL.
type HotelMasterAttributesResponse struct {
	Errors    *ErrorStruct    `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
	Countries *CountriesArray `xml:"Countries,omitempty" json:"Countries,omitempty" yaml:"Countries,omitempty"`
	Locals    *LocalsArray    `xml:"Locals,omitempty" json:"Locals,omitempty" yaml:"Locals,omitempty"`
	Boards    *BoardsArray    `xml:"Boards,omitempty" json:"Boards,omitempty" yaml:"Boards,omitempty"`
}

// HotelOffers was auto-generated from WSDL.
type HotelOffers struct {
	OfferId     *string `xml:"OfferId,omitempty" json:"OfferId,omitempty" yaml:"OfferId,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

// HotelOffersArray was auto-generated from WSDL.
type HotelOffersArray struct {
	Items []*HotelOffers `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// HotelPax was auto-generated from WSDL.
type HotelPax struct {
	RoomId      *string `xml:"RoomId,omitempty" json:"RoomId,omitempty" yaml:"RoomId,omitempty"`
	RoomCode    *string `xml:"RoomCode,omitempty" json:"RoomCode,omitempty" yaml:"RoomCode,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	PaxType     *string `xml:"PaxType,omitempty" json:"PaxType,omitempty" yaml:"PaxType,omitempty"`
}

// HotelPaxArray was auto-generated from WSDL.
type HotelPaxArray struct {
	Items []*HotelPax `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// HotelPaxReq was auto-generated from WSDL.
type HotelPaxReq struct {
	Title    *string `xml:"Title,omitempty" json:"Title,omitempty" yaml:"Title,omitempty"`
	Surname  *string `xml:"Surname,omitempty" json:"Surname,omitempty" yaml:"Surname,omitempty"`
	Name     *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	RoomCode *string `xml:"RoomCode,omitempty" json:"RoomCode,omitempty" yaml:"RoomCode,omitempty"`
}

// HotelPaxReqArray was auto-generated from WSDL.
type HotelPaxReqArray struct {
	Items []*HotelPaxReq `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// HotelReservationRequest was auto-generated from WSDL.
type HotelReservationRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	HotelPax    *HotelPaxReqArray  `xml:"HotelPax,omitempty" json:"HotelPax,omitempty" yaml:"HotelPax,omitempty"`
}

// HotelReservationResponse was auto-generated from WSDL.
type HotelReservationResponse struct {
	Errors        *ErrorStruct `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
	ReserveNumber *string      `xml:"ReserveNumber,omitempty" json:"ReserveNumber,omitempty" yaml:"ReserveNumber,omitempty"`
}

// HotelSearchRequest was auto-generated from WSDL.
type HotelSearchRequest struct {
	Credentials   *CredentialsStruct  `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Board         *string             `xml:"Board,omitempty" json:"Board,omitempty" yaml:"Board,omitempty"`
	CheckIn       *string             `xml:"CheckIn,omitempty" json:"CheckIn,omitempty" yaml:"CheckIn,omitempty"`
	CheckOut      *string             `xml:"CheckOut,omitempty" json:"CheckOut,omitempty" yaml:"CheckOut,omitempty"`
	Country       *string             `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
	Local         *string             `xml:"Local,omitempty" json:"Local,omitempty" yaml:"Local,omitempty"`
	Zone          *string             `xml:"Zone,omitempty" json:"Zone,omitempty" yaml:"Zone,omitempty"`
	Accommodation *AccommodationArray `xml:"Accommodation,omitempty" json:"Accommodation,omitempty" yaml:"Accommodation,omitempty"`
}

// HotelSearchResponse was auto-generated from WSDL.
type HotelSearchResponse struct {
	Errors      *ErrorStruct    `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
	SessionHash *string         `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	HotelList   *HotelListArray `xml:"HotelList,omitempty" json:"HotelList,omitempty" yaml:"HotelList,omitempty"`
}

// HotelServices was auto-generated from WSDL.
type HotelServices struct {
	Type        *string           `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Description *string           `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Quant       *string           `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	DateFrom    *string           `xml:"DateFrom,omitempty" json:"DateFrom,omitempty" yaml:"DateFrom,omitempty"`
	DateTo      *string           `xml:"DateTo,omitempty" json:"DateTo,omitempty" yaml:"DateTo,omitempty"`
	Status      *string           `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	HotelOffers *HotelOffersArray `xml:"HotelOffers,omitempty" json:"HotelOffers,omitempty" yaml:"HotelOffers,omitempty"`
}

// HotelServicesArray was auto-generated from WSDL.
type HotelServicesArray struct {
	Items []*HotelServices `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// HotelSimulationRequest was auto-generated from WSDL.
type HotelSimulationRequest struct {
	Credentials    *CredentialsStruct   `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash    *string              `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	HotelId        *string              `xml:"HotelId,omitempty" json:"HotelId,omitempty" yaml:"HotelId,omitempty"`
	RoomsSelection *RoomsSelectionArray `xml:"RoomsSelection,omitempty" json:"RoomsSelection,omitempty" yaml:"RoomsSelection,omitempty"`
	BoardId        *string              `xml:"BoardId,omitempty" json:"BoardId,omitempty" yaml:"BoardId,omitempty"`
}

// HotelSimulationResponse was auto-generated from WSDL.
type HotelSimulationResponse struct {
	Errors        *ErrorStruct        `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
	SessionHash   *string             `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	HotelServices *HotelServicesArray `xml:"HotelServices,omitempty" json:"HotelServices,omitempty" yaml:"HotelServices,omitempty"`
	HotelCalcs    *HotelCalcsArray    `xml:"HotelCalcs,omitempty" json:"HotelCalcs,omitempty" yaml:"HotelCalcs,omitempty"`
	HotelPax      *HotelPaxArray      `xml:"HotelPax,omitempty" json:"HotelPax,omitempty" yaml:"HotelPax,omitempty"`
}

// HotelText was auto-generated from WSDL.
type HotelText struct {
	Category *string `xml:"Category,omitempty" json:"Category,omitempty" yaml:"Category,omitempty"`
	Nome     *string `xml:"Nome,omitempty" json:"Nome,omitempty" yaml:"Nome,omitempty"`
	Tipo     *string `xml:"Tipo,omitempty" json:"Tipo,omitempty" yaml:"Tipo,omitempty"`
	Valor    *string `xml:"Valor,omitempty" json:"Valor,omitempty" yaml:"Valor,omitempty"`
}

// HotelTextArray was auto-generated from WSDL.
type HotelTextArray struct {
	Items []*HotelText `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Hotels was auto-generated from WSDL.
type Hotels struct {
	Description           *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Address               *string `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	Postalcode            *string `xml:"Postalcode,omitempty" json:"Postalcode,omitempty" yaml:"Postalcode,omitempty"`
	Contact               *string `xml:"Contact,omitempty" json:"Contact,omitempty" yaml:"Contact,omitempty"`
	Local                 *string `xml:"Local,omitempty" json:"Local,omitempty" yaml:"Local,omitempty"`
	Room                  *string `xml:"Room,omitempty" json:"Room,omitempty" yaml:"Room,omitempty"`
	BaseBoard_Description *string `xml:"BaseBoard_Description,omitempty" json:"BaseBoard_Description,omitempty" yaml:"BaseBoard_Description,omitempty"`
	BaseBoard_Code        *string `xml:"BaseBoard_Code,omitempty" json:"BaseBoard_Code,omitempty" yaml:"BaseBoard_Code,omitempty"`
	DateCheckIn           *string `xml:"DateCheckIn,omitempty" json:"DateCheckIn,omitempty" yaml:"DateCheckIn,omitempty"`
	DateCheckOut          *string `xml:"DateCheckOut,omitempty" json:"DateCheckOut,omitempty" yaml:"DateCheckOut,omitempty"`
}

// HotelsArray was auto-generated from WSDL.
type HotelsArray struct {
	Items []*Hotels `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Included was auto-generated from WSDL.
type Included struct {
	ID          *string `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Code        *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

// InfantStruct was auto-generated from WSDL.
type InfantStruct struct {
	Min *string `xml:"Min,omitempty" json:"Min,omitempty" yaml:"Min,omitempty"`
	Max *string `xml:"Max,omitempty" json:"Max,omitempty" yaml:"Max,omitempty"`
}

// Info was auto-generated from WSDL.
type Info struct {
	PaxInfoId               *string           `xml:"PaxInfoId,omitempty" json:"PaxInfoId,omitempty" yaml:"PaxInfoId,omitempty"`
	PaxInfoPriority         *string           `xml:"PaxInfoPriority,omitempty" json:"PaxInfoPriority,omitempty" yaml:"PaxInfoPriority,omitempty"`
	PaxInfoElementType      *string           `xml:"PaxInfoElementType,omitempty" json:"PaxInfoElementType,omitempty" yaml:"PaxInfoElementType,omitempty"`
	PaxInfoMapType          *string           `xml:"PaxInfoMapType,omitempty" json:"PaxInfoMapType,omitempty" yaml:"PaxInfoMapType,omitempty"`
	PaxInfoDescription      *string           `xml:"PaxInfoDescription,omitempty" json:"PaxInfoDescription,omitempty" yaml:"PaxInfoDescription,omitempty"`
	PaxInfoRequiredAdult    *string           `xml:"PaxInfoRequiredAdult,omitempty" json:"PaxInfoRequiredAdult,omitempty" yaml:"PaxInfoRequiredAdult,omitempty"`
	PaxInfoVisibilityAdult  *string           `xml:"PaxInfoVisibilityAdult,omitempty" json:"PaxInfoVisibilityAdult,omitempty" yaml:"PaxInfoVisibilityAdult,omitempty"`
	PaxInfoRequiredChild    *string           `xml:"PaxInfoRequiredChild,omitempty" json:"PaxInfoRequiredChild,omitempty" yaml:"PaxInfoRequiredChild,omitempty"`
	PaxInfoVisibilityChild  *string           `xml:"PaxInfoVisibilityChild,omitempty" json:"PaxInfoVisibilityChild,omitempty" yaml:"PaxInfoVisibilityChild,omitempty"`
	PaxInfoRequiredInfant   *string           `xml:"PaxInfoRequiredInfant,omitempty" json:"PaxInfoRequiredInfant,omitempty" yaml:"PaxInfoRequiredInfant,omitempty"`
	PaxInfoVisibilityInfant *string           `xml:"PaxInfoVisibilityInfant,omitempty" json:"PaxInfoVisibilityInfant,omitempty" yaml:"PaxInfoVisibilityInfant,omitempty"`
	PaxInfoApply            *string           `xml:"PaxInfoApply,omitempty" json:"PaxInfoApply,omitempty" yaml:"PaxInfoApply,omitempty"`
	PaxInfoItems            *PaxInfoItemArray `xml:"PaxInfoItems,omitempty" json:"PaxInfoItems,omitempty" yaml:"PaxInfoItems,omitempty"`
}

// InfoArray was auto-generated from WSDL.
type InfoArray struct {
	Items []*Info `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Local was auto-generated from WSDL.
type Local struct {
	Code   *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Nights *string `xml:"Nights,omitempty" json:"Nights,omitempty" yaml:"Nights,omitempty"`
}

// LocalArray was auto-generated from WSDL.
type LocalArray struct {
	Items []*Local `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Locals was auto-generated from WSDL.
type Locals struct {
	Code  *string     `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name  *string     `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Zones *ZonesArray `xml:"Zones,omitempty" json:"Zones,omitempty" yaml:"Zones,omitempty"`
}

// LocalsArray was auto-generated from WSDL.
type LocalsArray struct {
	Items []*Locals `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Location was auto-generated from WSDL.
type Location struct {
	Code        *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Country     *string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
}

// LocationArray was auto-generated from WSDL.
type LocationArray struct {
	Items []*Location `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// MessagesList was auto-generated from WSDL.
type MessagesList struct {
}

// MessagesListArray was auto-generated from WSDL.
type MessagesListArray struct {
	Items []*MessagesList `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// MultimediaContentTexts was auto-generated from WSDL.
type MultimediaContentTexts struct {
	Id          *string `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Cat_code    *string `xml:"cat_code,omitempty" json:"cat_code,omitempty" yaml:"cat_code,omitempty"`
	Cat_desc    *string `xml:"cat_desc,omitempty" json:"cat_desc,omitempty" yaml:"cat_desc,omitempty"`
	Description *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
}

// MultimediaContentTextsArray was auto-generated from WSDL.
type MultimediaContentTextsArray struct {
	Items []*MultimediaContentTexts `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ObservationsHtl was auto-generated from WSDL.
type ObservationsHtl struct {
	Status *string                              `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Detail *RestrictionsObservationsDetailArray `xml:"Detail,omitempty" json:"Detail,omitempty" yaml:"Detail,omitempty"`
}

// ObservationsHtlArray was auto-generated from WSDL.
type ObservationsHtlArray struct {
	Items []*ObservationsHtl `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Option was auto-generated from WSDL.
type Option struct {
	Option_ID     *string        `xml:"Option_ID,omitempty" json:"Option_ID,omitempty" yaml:"Option_ID,omitempty"`
	Option_Name   *string        `xml:"Option_Name,omitempty" json:"Option_Name,omitempty" yaml:"Option_Name,omitempty"`
	Date_ini      *string        `xml:"Date_ini,omitempty" json:"Date_ini,omitempty" yaml:"Date_ini,omitempty"`
	Date_end      *string        `xml:"Date_end,omitempty" json:"Date_end,omitempty" yaml:"Date_end,omitempty"`
	Regime        *string        `xml:"Regime,omitempty" json:"Regime,omitempty" yaml:"Regime,omitempty"`
	RegimeCode    *string        `xml:"RegimeCode,omitempty" json:"RegimeCode,omitempty" yaml:"RegimeCode,omitempty"`
	DepDateArray  *DepDateArray  `xml:"DepDateArray,omitempty" json:"DepDateArray,omitempty" yaml:"DepDateArray,omitempty"`
	DepLocalArray *DepLocalArray `xml:"DepLocalArray,omitempty" json:"DepLocalArray,omitempty" yaml:"DepLocalArray,omitempty"`
}

// OptionArray was auto-generated from WSDL.
type OptionArray struct {
	Items []*Option `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// PagmaticDetail was auto-generated from WSDL.
type PagmaticDetail struct {
	Entidade    *string `xml:"Entidade,omitempty" json:"Entidade,omitempty" yaml:"Entidade,omitempty"`
	SubEntidade *string `xml:"SubEntidade,omitempty" json:"SubEntidade,omitempty" yaml:"SubEntidade,omitempty"`
	Valor       *string `xml:"Valor,omitempty" json:"Valor,omitempty" yaml:"Valor,omitempty"`
}

// PagmaticInfo was auto-generated from WSDL.
type PagmaticInfo struct {
	PagmaticID       *string         `xml:"PagmaticID,omitempty" json:"PagmaticID,omitempty" yaml:"PagmaticID,omitempty"`
	PagmaticHash     *string         `xml:"PagmaticHash,omitempty" json:"PagmaticHash,omitempty" yaml:"PagmaticHash,omitempty"`
	PagmaticDeadline *string         `xml:"PagmaticDeadline,omitempty" json:"PagmaticDeadline,omitempty" yaml:"PagmaticDeadline,omitempty"`
	PagmaticDetail   *PagmaticDetail `xml:"PagmaticDetail,omitempty" json:"PagmaticDetail,omitempty" yaml:"PagmaticDetail,omitempty"`
}

// Pax was auto-generated from WSDL.
type Pax struct {
	Code    *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Desc    *string `xml:"Desc,omitempty" json:"Desc,omitempty" yaml:"Desc,omitempty"`
	Quant   *string `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	AgeFrom *string `xml:"AgeFrom,omitempty" json:"AgeFrom,omitempty" yaml:"AgeFrom,omitempty"`
	AgeTo   *string `xml:"AgeTo,omitempty" json:"AgeTo,omitempty" yaml:"AgeTo,omitempty"`
}

// PaxArray was auto-generated from WSDL.
type PaxArray struct {
	Items []*Pax `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// PaxInfo was auto-generated from WSDL.
type PaxInfo struct {
	ResPaxId   *string          `xml:"ResPaxId,omitempty" json:"ResPaxId,omitempty" yaml:"ResPaxId,omitempty"`
	Title      *string          `xml:"Title,omitempty" json:"Title,omitempty" yaml:"Title,omitempty"`
	Name       *string          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Surname    *string          `xml:"Surname,omitempty" json:"Surname,omitempty" yaml:"Surname,omitempty"`
	Additional *AdditionalArray `xml:"Additional,omitempty" json:"Additional,omitempty" yaml:"Additional,omitempty"`
}

// PaxInfoArray was auto-generated from WSDL.
type PaxInfoArray struct {
	Items []*PaxInfo `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// PaxInfoItem was auto-generated from WSDL.
type PaxInfoItem struct {
	ItemCode        *string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`
	ItemDescription *string `xml:"ItemDescription,omitempty" json:"ItemDescription,omitempty" yaml:"ItemDescription,omitempty"`
}

// PaxInfoItemArray was auto-generated from WSDL.
type PaxInfoItemArray struct {
	Items []*PaxInfoItem `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// PaxRes was auto-generated from WSDL.
type PaxRes struct {
	RoomNum  *string `xml:"RoomNum,omitempty" json:"RoomNum,omitempty" yaml:"RoomNum,omitempty"`
	Title    *string `xml:"Title,omitempty" json:"Title,omitempty" yaml:"Title,omitempty"`
	ChildAge *string `xml:"ChildAge,omitempty" json:"ChildAge,omitempty" yaml:"ChildAge,omitempty"`
	Name     *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Surname  *string `xml:"Surname,omitempty" json:"Surname,omitempty" yaml:"Surname,omitempty"`
	Room     *string `xml:"Room,omitempty" json:"Room,omitempty" yaml:"Room,omitempty"`
}

// PaxResArray was auto-generated from WSDL.
type PaxResArray struct {
	Items []*PaxRes `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// PaxResPaxs was auto-generated from WSDL.
type PaxResPaxs struct {
	NumAdults *string `xml:"NumAdults,omitempty" json:"NumAdults,omitempty" yaml:"NumAdults,omitempty"`
	NumChilds *string `xml:"NumChilds,omitempty" json:"NumChilds,omitempty" yaml:"NumChilds,omitempty"`
	ChildAge  *string `xml:"ChildAge,omitempty" json:"ChildAge,omitempty" yaml:"ChildAge,omitempty"`
}

// PaxResPaxsArray was auto-generated from WSDL.
type PaxResPaxsArray struct {
	Items []*PaxResPaxs `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// PaxRoom was auto-generated from WSDL.
type PaxRoom struct {
	Id            *string `xml:"Id,omitempty" json:"Id,omitempty" yaml:"Id,omitempty"`
	Title         *string `xml:"Title,omitempty" json:"Title,omitempty" yaml:"Title,omitempty"`
	Name          *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Surname       *string `xml:"Surname,omitempty" json:"Surname,omitempty" yaml:"Surname,omitempty"`
	DepartureDesc *string `xml:"DepartureDesc,omitempty" json:"DepartureDesc,omitempty" yaml:"DepartureDesc,omitempty"`
	Insurance     *string `xml:"Insurance,omitempty" json:"Insurance,omitempty" yaml:"Insurance,omitempty"`
	InsuranceCode *string `xml:"InsuranceCode,omitempty" json:"InsuranceCode,omitempty" yaml:"InsuranceCode,omitempty"`
	RoomType      *string `xml:"RoomType,omitempty" json:"RoomType,omitempty" yaml:"RoomType,omitempty"`
	Status        *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// PaxRoomArray was auto-generated from WSDL.
type PaxRoomArray struct {
	Items []*PaxRoom `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Paxs was auto-generated from WSDL.
type Paxs struct {
	PaxId     *string    `xml:"PaxId,omitempty" json:"PaxId,omitempty" yaml:"PaxId,omitempty"`
	PaxName   *string    `xml:"PaxName,omitempty" json:"PaxName,omitempty" yaml:"PaxName,omitempty"`
	FlightIn  *FlightIn  `xml:"FlightIn,omitempty" json:"FlightIn,omitempty" yaml:"FlightIn,omitempty"`
	FlightOut *FlightOut `xml:"FlightOut,omitempty" json:"FlightOut,omitempty" yaml:"FlightOut,omitempty"`
}

// PaxsArray was auto-generated from WSDL.
type PaxsArray struct {
	Items []*Paxs `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// PaxsSimulation was auto-generated from WSDL.
type PaxsSimulation struct {
	NumAdults *string `xml:"NumAdults,omitempty" json:"NumAdults,omitempty" yaml:"NumAdults,omitempty"`
	NumChilds *string `xml:"NumChilds,omitempty" json:"NumChilds,omitempty" yaml:"NumChilds,omitempty"`
	ChildAge  *string `xml:"ChildAge,omitempty" json:"ChildAge,omitempty" yaml:"ChildAge,omitempty"`
}

// PaxsStatus was auto-generated from WSDL.
type PaxsStatus struct {
	PaxId  *string `xml:"PaxId,omitempty" json:"PaxId,omitempty" yaml:"PaxId,omitempty"`
	Status *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Error  *string `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
}

// PaxsStatusArray was auto-generated from WSDL.
type PaxsStatusArray struct {
	Items []*PaxsStatus `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// PaymentMethods was auto-generated from WSDL.
type PaymentMethods struct {
	Code *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Desc *string `xml:"Desc,omitempty" json:"Desc,omitempty" yaml:"Desc,omitempty"`
}

// PaymentMethodsArray was auto-generated from WSDL.
type PaymentMethodsArray struct {
	Items []*PaymentMethods `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Photo was auto-generated from WSDL.
type Photo struct {
	ImageUrl *string `xml:"ImageUrl,omitempty" json:"ImageUrl,omitempty" yaml:"ImageUrl,omitempty"`
}

// PhotoArray was auto-generated from WSDL.
type PhotoArray struct {
	Items []*Photo `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// PoliciesDetail was auto-generated from WSDL.
type PoliciesDetail struct {
	ValueType   *string `xml:"ValueType,omitempty" json:"ValueType,omitempty" yaml:"ValueType,omitempty"`
	Value       *string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	DaysFrom    *string `xml:"DaysFrom,omitempty" json:"DaysFrom,omitempty" yaml:"DaysFrom,omitempty"`
	DaysTo      *string `xml:"DaysTo,omitempty" json:"DaysTo,omitempty" yaml:"DaysTo,omitempty"`
}

// PoliciesDetailArray was auto-generated from WSDL.
type PoliciesDetailArray struct {
	Items []*PoliciesDetail `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Price was auto-generated from WSDL.
type Price struct {
	Service       *string `xml:"Service,omitempty" json:"Service,omitempty" yaml:"Service,omitempty"`
	Price_Type    *string `xml:"Price_Type,omitempty" json:"Price_Type,omitempty" yaml:"Price_Type,omitempty"`
	Description   *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Quant         *string `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	Sell_Value    *string `xml:"Sell_Value,omitempty" json:"Sell_Value,omitempty" yaml:"Sell_Value,omitempty"`
	Total_Value   *string `xml:"Total_Value,omitempty" json:"Total_Value,omitempty" yaml:"Total_Value,omitempty"`
	Comission     *string `xml:"Comission,omitempty" json:"Comission,omitempty" yaml:"Comission,omitempty"`
	Optional      *string `xml:"Optional,omitempty" json:"Optional,omitempty" yaml:"Optional,omitempty"`
	RoomCode      *string `xml:"RoomCode,omitempty" json:"RoomCode,omitempty" yaml:"RoomCode,omitempty"`
	TypePax       *string `xml:"TypePax,omitempty" json:"TypePax,omitempty" yaml:"TypePax,omitempty"`
	NumChild      *string `xml:"NumChild,omitempty" json:"NumChild,omitempty" yaml:"NumChild,omitempty"`
	PriceTypeCode *string `xml:"PriceTypeCode,omitempty" json:"PriceTypeCode,omitempty" yaml:"PriceTypeCode,omitempty"`
}

// PriceArray was auto-generated from WSDL.
type PriceArray struct {
	Items []*Price `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// NOT AUTO-GENERATED
type ProductWrapper struct {
	Product    Product  `json:"product"`
	Tags       []string `json:"tags"`
	Enabled    *bool    `json:"enabled"`
	VisitCount int64    `json:"visitCount"`
}

// Product was auto-generated from WSDL.
type Product struct {
	Code                   *string                 `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	ProdCode               *string                 `xml:"ProdCode,omitempty" json:"ProdCode,omitempty" yaml:"ProdCode,omitempty"`
	Name                   *string                 `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	ImageUrl               *string                 `xml:"ImageUrl,omitempty" json:"ImageUrl,omitempty" yaml:"ImageUrl,omitempty"`
	Continent              *string                 `xml:"Continent,omitempty" json:"Continent,omitempty" yaml:"Continent,omitempty"`
	Country                *string                 `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
	Location               *string                 `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	Zone                   *string                 `xml:"Zone,omitempty" json:"Zone,omitempty" yaml:"Zone,omitempty"`
	BaseDays               *string                 `xml:"BaseDays,omitempty" json:"BaseDays,omitempty" yaml:"BaseDays,omitempty"`
	OperationFrom          *string                 `xml:"OperationFrom,omitempty" json:"OperationFrom,omitempty" yaml:"OperationFrom,omitempty"`
	OperationTo            *string                 `xml:"OperationTo,omitempty" json:"OperationTo,omitempty" yaml:"OperationTo,omitempty"`
	ShortDesc              *string                 `xml:"ShortDesc,omitempty" json:"ShortDesc,omitempty" yaml:"ShortDesc,omitempty"`
	PopupInfo              *string                 `xml:"PopupInfo,omitempty" json:"PopupInfo,omitempty" yaml:"PopupInfo,omitempty"`
	PriceFrom              *string                 `xml:"PriceFrom,omitempty" json:"PriceFrom,omitempty" yaml:"PriceFrom,omitempty"`
	ProdType               *string                 `xml:"ProdType,omitempty" json:"ProdType,omitempty" yaml:"ProdType,omitempty"`
	TextContentsArray      *TextContentArray       `xml:"TextContentsArray,omitempty" json:"TextContentsArray,omitempty" yaml:"TextContentsArray,omitempty"`
	PhotoArray             *PhotoArray             `xml:"PhotoArray,omitempty" json:"PhotoArray,omitempty" yaml:"PhotoArray,omitempty"`
	GeneralConditionsArray *GeneralConditionsArray `xml:"GeneralConditionsArray,omitempty" json:"GeneralConditionsArray,omitempty" yaml:"GeneralConditionsArray,omitempty"`
	CanReserve             *string                 `xml:"CanReserve,omitempty" json:"CanReserve,omitempty" yaml:"CanReserve,omitempty"`
	Precalc                *string                 `xml:"Precalc,omitempty" json:"Precalc,omitempty" yaml:"Precalc,omitempty"`
	HasFlights             *string                 `xml:"HasFlights,omitempty" json:"HasFlights,omitempty" yaml:"HasFlights,omitempty"`
	ItineraryQTY           *string                 `xml:"ItineraryQTY,omitempty" json:"ItineraryQTY,omitempty" yaml:"ItineraryQTY,omitempty"`
}

// ProductArray was auto-generated from WSDL.
type ProductArray struct {
	Items []*Product `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ProductArrayByCalendar was auto-generated from WSDL.
type ProductArrayByCalendar struct {
	ProdCode     *string `xml:"ProdCode,omitempty" json:"ProdCode,omitempty" yaml:"ProdCode,omitempty"`
	ProdCodeDesc *string `xml:"ProdCodeDesc,omitempty" json:"ProdCodeDesc,omitempty" yaml:"ProdCodeDesc,omitempty"`
	ProdCodeName *string `xml:"ProdCodeName,omitempty" json:"ProdCodeName,omitempty" yaml:"ProdCodeName,omitempty"`
	OptionID     *string `xml:"OptionID,omitempty" json:"OptionID,omitempty" yaml:"OptionID,omitempty"`
	OptionDesc   *string `xml:"OptionDesc,omitempty" json:"OptionDesc,omitempty" yaml:"OptionDesc,omitempty"`
	DepDateID    *string `xml:"DepDateID,omitempty" json:"DepDateID,omitempty" yaml:"DepDateID,omitempty"`
	Allot        *string `xml:"Allot,omitempty" json:"Allot,omitempty" yaml:"Allot,omitempty"`
	DepDateDesc  *string `xml:"DepDateDesc,omitempty" json:"DepDateDesc,omitempty" yaml:"DepDateDesc,omitempty"`
	DepLocalID   *string `xml:"DepLocalID,omitempty" json:"DepLocalID,omitempty" yaml:"DepLocalID,omitempty"`
	DepLocalDesc *string `xml:"DepLocalDesc,omitempty" json:"DepLocalDesc,omitempty" yaml:"DepLocalDesc,omitempty"`
	AvailStatus  *string `xml:"AvailStatus,omitempty" json:"AvailStatus,omitempty" yaml:"AvailStatus,omitempty"`
}

// ProductArrayByCalendarArray was auto-generated from WSDL.
type ProductArrayByCalendarArray struct {
	Items []*ProductArrayByCalendar `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ProductType was auto-generated from WSDL.
type ProductType struct {
	Code        *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

// ProductTypeArray was auto-generated from WSDL.
type ProductTypeArray struct {
	Items []*ProductType `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ReductionNights was auto-generated from WSDL.
type ReductionNights struct {
	Begin *string `xml:"Begin,omitempty" json:"Begin,omitempty" yaml:"Begin,omitempty"`
	End   *string `xml:"End,omitempty" json:"End,omitempty" yaml:"End,omitempty"`
}

// ResCancelPolicies was auto-generated from WSDL.
type ResCancelPolicies struct {
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	StartDate   *string `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`
	EndDate     *string `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`
	Value       *string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// ResCancelPoliciesArray was auto-generated from WSDL.
type ResCancelPoliciesArray struct {
	Items []*ResCancelPolicies `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ResHead was auto-generated from WSDL.
type ResHead struct {
	ClientCode        *string `xml:"ClientCode,omitempty" json:"ClientCode,omitempty" yaml:"ClientCode,omitempty"`
	ClientName        *string `xml:"ClientName,omitempty" json:"ClientName,omitempty" yaml:"ClientName,omitempty"`
	Product           *string `xml:"Product,omitempty" json:"Product,omitempty" yaml:"Product,omitempty"`
	DepDate           *string `xml:"DepDate,omitempty" json:"DepDate,omitempty" yaml:"DepDate,omitempty"`
	ArrDate           *string `xml:"ArrDate,omitempty" json:"ArrDate,omitempty" yaml:"ArrDate,omitempty"`
	TotalGrossVal     *string `xml:"TotalGrossVal,omitempty" json:"TotalGrossVal,omitempty" yaml:"TotalGrossVal,omitempty"`
	TotalGrossPVP     *string `xml:"TotalGrossPVP,omitempty" json:"TotalGrossPVP,omitempty" yaml:"TotalGrossPVP,omitempty"`
	TotalComissionVal *string `xml:"TotalComissionVal,omitempty" json:"TotalComissionVal,omitempty" yaml:"TotalComissionVal,omitempty"`
	TotalInvoiceVal   *string `xml:"TotalInvoiceVal,omitempty" json:"TotalInvoiceVal,omitempty" yaml:"TotalInvoiceVal,omitempty"`
	TotalNetVal       *string `xml:"TotalNetVal,omitempty" json:"TotalNetVal,omitempty" yaml:"TotalNetVal,omitempty"`
	ClientRef         *string `xml:"ClientRef,omitempty" json:"ClientRef,omitempty" yaml:"ClientRef,omitempty"`
	ClientObs         *string `xml:"ClientObs,omitempty" json:"ClientObs,omitempty" yaml:"ClientObs,omitempty"`
	SyncReserve       *string `xml:"SyncReserve,omitempty" json:"SyncReserve,omitempty" yaml:"SyncReserve,omitempty"`
}

// ResHeadArray was auto-generated from WSDL.
type ResHeadArray struct {
	Items []*ResHead `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ResList was auto-generated from WSDL.
type ResList struct {
	Reserve    *string `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	ClientName *string `xml:"ClientName,omitempty" json:"ClientName,omitempty" yaml:"ClientName,omitempty"`
	CreateDate *string `xml:"CreateDate,omitempty" json:"CreateDate,omitempty" yaml:"CreateDate,omitempty"`
	DepDate    *string `xml:"DepDate,omitempty" json:"DepDate,omitempty" yaml:"DepDate,omitempty"`
	ArrDate    *string `xml:"ArrDate,omitempty" json:"ArrDate,omitempty" yaml:"ArrDate,omitempty"`
	FirstPax   *string `xml:"FirstPax,omitempty" json:"FirstPax,omitempty" yaml:"FirstPax,omitempty"`
	AgentName  *string `xml:"AgentName,omitempty" json:"AgentName,omitempty" yaml:"AgentName,omitempty"`
	AgentRef   *string `xml:"AgentRef,omitempty" json:"AgentRef,omitempty" yaml:"AgentRef,omitempty"`
}

// ResListArray was auto-generated from WSDL.
type ResListArray struct {
	Items []*ResList `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ResObs was auto-generated from WSDL.
type ResObs struct {
	Text *string `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`
	User *string `xml:"User,omitempty" json:"User,omitempty" yaml:"User,omitempty"`
}

// ResObsArray was auto-generated from WSDL.
type ResObsArray struct {
	Items []*ResObs `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ResPax was auto-generated from WSDL.
type ResPax struct {
	Title       *string `xml:"Title,omitempty" json:"Title,omitempty" yaml:"Title,omitempty"`
	DateofBirth *string `xml:"DateofBirth,omitempty" json:"DateofBirth,omitempty" yaml:"DateofBirth,omitempty"`
	Name        *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Surname     *string `xml:"Surname,omitempty" json:"Surname,omitempty" yaml:"Surname,omitempty"`
	Room        *string `xml:"Room,omitempty" json:"Room,omitempty" yaml:"Room,omitempty"`
}

// ResPaxArray was auto-generated from WSDL.
type ResPaxArray struct {
	Items []*ResPax `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ResPaymentConditions was auto-generated from WSDL.
type ResPaymentConditions struct {
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	StartDate   *string `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`
	EndDate     *string `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`
	Value       *string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// ResPaymentConditionsArray was auto-generated from WSDL.
type ResPaymentConditionsArray struct {
	Items []*ResPaymentConditions `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ResService was auto-generated from WSDL.
type ResService struct {
	Type        *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Quant       *string `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	DateFrom    *string `xml:"DateFrom,omitempty" json:"DateFrom,omitempty" yaml:"DateFrom,omitempty"`
	DateTo      *string `xml:"DateTo,omitempty" json:"DateTo,omitempty" yaml:"DateTo,omitempty"`
	Status      *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// ResServiceArray was auto-generated from WSDL.
type ResServiceArray struct {
	Items []*ResService `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ReserveListRequest was auto-generated from WSDL.
type ReserveListRequest struct {
	Credentials    *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	DepDateFrom    *string            `xml:"DepDateFrom,omitempty" json:"DepDateFrom,omitempty" yaml:"DepDateFrom,omitempty"`
	DepDateTo      *string            `xml:"DepDateTo,omitempty" json:"DepDateTo,omitempty" yaml:"DepDateTo,omitempty"`
	CreateDateFrom *string            `xml:"CreateDateFrom,omitempty" json:"CreateDateFrom,omitempty" yaml:"CreateDateFrom,omitempty"`
	CreateDateTo   *string            `xml:"CreateDateTo,omitempty" json:"CreateDateTo,omitempty" yaml:"CreateDateTo,omitempty"`
	Reserve        *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	Status         *string            `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// ReserveListResponse was auto-generated from WSDL.
type ReserveListResponse struct {
	ReserveList *ResListArray `xml:"ReserveList,omitempty" json:"ReserveList,omitempty" yaml:"ReserveList,omitempty"`
	Errors      *ErrorStruct  `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// RestrictionsHtl was auto-generated from WSDL.
type RestrictionsHtl struct {
	Status *string                              `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Detail *RestrictionsObservationsDetailArray `xml:"Detail,omitempty" json:"Detail,omitempty" yaml:"Detail,omitempty"`
}

// RestrictionsHtlArray was auto-generated from WSDL.
type RestrictionsHtlArray struct {
	Items []*RestrictionsHtl `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// RestrictionsObservationsDetail was auto-generated from WSDL.
type RestrictionsObservationsDetail struct {
	Type        *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	MinNights   *string `xml:"MinNights,omitempty" json:"MinNights,omitempty" yaml:"MinNights,omitempty"`
	DateFrom    *string `xml:"DateFrom,omitempty" json:"DateFrom,omitempty" yaml:"DateFrom,omitempty"`
	DateTo      *string `xml:"DateTo,omitempty" json:"DateTo,omitempty" yaml:"DateTo,omitempty"`
}

// RestrictionsObservationsDetailArray was auto-generated from
// WSDL.
type RestrictionsObservationsDetailArray struct {
	Items []*RestrictionsObservationsDetail `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// RetrieveReserveRequest was auto-generated from WSDL.
type RetrieveReserveRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
}

// RetrieveReserveResponse was auto-generated from WSDL.
type RetrieveReserveResponse struct {
	Head              *ResHeadArray              `xml:"Head,omitempty" json:"Head,omitempty" yaml:"Head,omitempty"`
	Paxs              *ResPaxArray               `xml:"Paxs,omitempty" json:"Paxs,omitempty" yaml:"Paxs,omitempty"`
	Services          *ResServiceArray           `xml:"Services,omitempty" json:"Services,omitempty" yaml:"Services,omitempty"`
	Calcs             *DynResCalcsArray          `xml:"Calcs,omitempty" json:"Calcs,omitempty" yaml:"Calcs,omitempty"`
	Obs               *ResObsArray               `xml:"Obs,omitempty" json:"Obs,omitempty" yaml:"Obs,omitempty"`
	CancelPolicies    *ResCancelPoliciesArray    `xml:"CancelPolicies,omitempty" json:"CancelPolicies,omitempty" yaml:"CancelPolicies,omitempty"`
	PaymentConditions *ResPaymentConditionsArray `xml:"PaymentConditions,omitempty" json:"PaymentConditions,omitempty" yaml:"PaymentConditions,omitempty"`
	Errors            *ErrorStruct               `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// Room was auto-generated from WSDL.
type Room struct {
	Room_Code *string `xml:"Room_Code,omitempty" json:"Room_Code,omitempty" yaml:"Room_Code,omitempty"`
	Room_Desc *string `xml:"Room_Desc,omitempty" json:"Room_Desc,omitempty" yaml:"Room_Desc,omitempty"`
	Adults    *string `xml:"Adults,omitempty" json:"Adults,omitempty" yaml:"Adults,omitempty"`
	Childs    *string `xml:"Childs,omitempty" json:"Childs,omitempty" yaml:"Childs,omitempty"`
	Price     *string `xml:"Price,omitempty" json:"Price,omitempty" yaml:"Price,omitempty"`
	Status    *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Age_from  *string `xml:"Age_from,omitempty" json:"Age_from,omitempty" yaml:"Age_from,omitempty"`
	Age_to    *string `xml:"Age_to,omitempty" json:"Age_to,omitempty" yaml:"Age_to,omitempty"`
}

// RoomArray was auto-generated from WSDL.
type RoomArray struct {
	Items []*Room `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// RoomTypeOption was auto-generated from WSDL.
type RoomTypeOption struct {
	RoomNum   *string `xml:"RoomNum,omitempty" json:"RoomNum,omitempty" yaml:"RoomNum,omitempty"`
	Code      *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	AdultAges *string `xml:"AdultAges,omitempty" json:"AdultAges,omitempty" yaml:"AdultAges,omitempty"`
	ChildAges *string `xml:"ChildAges,omitempty" json:"ChildAges,omitempty" yaml:"ChildAges,omitempty"`
}

// RoomTypeOptionArray was auto-generated from WSDL.
type RoomTypeOptionArray struct {
	Items []*RoomTypeOption `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Rooms was auto-generated from WSDL.
type Rooms struct {
	Room_Code *string `xml:"Room_Code,omitempty" json:"Room_Code,omitempty" yaml:"Room_Code,omitempty"`
	Room_Age1 *string `xml:"Room_Age1,omitempty" json:"Room_Age1,omitempty" yaml:"Room_Age1,omitempty"`
	Room_Age2 *string `xml:"Room_Age2,omitempty" json:"Room_Age2,omitempty" yaml:"Room_Age2,omitempty"`
}

// RoomsArray was auto-generated from WSDL.
type RoomsArray struct {
	Items []*Rooms `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// RoomsHtl was auto-generated from WSDL.
type RoomsHtl struct {
	Index         *string `xml:"Index,omitempty" json:"Index,omitempty" yaml:"Index,omitempty"`
	Id            *string `xml:"Id,omitempty" json:"Id,omitempty" yaml:"Id,omitempty"`
	Description   *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	NumAdults     *string `xml:"NumAdults,omitempty" json:"NumAdults,omitempty" yaml:"NumAdults,omitempty"`
	NumChilds     *string `xml:"NumChilds,omitempty" json:"NumChilds,omitempty" yaml:"NumChilds,omitempty"`
	RoomCode      *string `xml:"RoomCode,omitempty" json:"RoomCode,omitempty" yaml:"RoomCode,omitempty"`
	AgeChildsFrom *string `xml:"AgeChildsFrom,omitempty" json:"AgeChildsFrom,omitempty" yaml:"AgeChildsFrom,omitempty"`
	AgeChildsTo   *string `xml:"AgeChildsTo,omitempty" json:"AgeChildsTo,omitempty" yaml:"AgeChildsTo,omitempty"`
	Status        *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	AvailStatus   *string `xml:"AvailStatus,omitempty" json:"AvailStatus,omitempty" yaml:"AvailStatus,omitempty"`
	SellPrice     *string `xml:"SellPrice,omitempty" json:"SellPrice,omitempty" yaml:"SellPrice,omitempty"`
	BuyPrice      *string `xml:"BuyPrice,omitempty" json:"BuyPrice,omitempty" yaml:"BuyPrice,omitempty"`
}

// RoomsHtlArray was auto-generated from WSDL.
type RoomsHtlArray struct {
	Items []*RoomsHtl `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// RoomsList was auto-generated from WSDL.
type RoomsList struct {
	RoomNum *string `xml:"RoomNum,omitempty" json:"RoomNum,omitempty" yaml:"RoomNum,omitempty"`
	Name    *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Board   *string `xml:"Board,omitempty" json:"Board,omitempty" yaml:"Board,omitempty"`
}

// RoomsListArray was auto-generated from WSDL.
type RoomsListArray struct {
	Items []*RoomsList `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// RoomsSelection was auto-generated from WSDL.
type RoomsSelection struct {
	RoomIndex *string `xml:"RoomIndex,omitempty" json:"RoomIndex,omitempty" yaml:"RoomIndex,omitempty"`
	Quant     *string `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	Ages      *string `xml:"Ages,omitempty" json:"Ages,omitempty" yaml:"Ages,omitempty"`
}

// RoomsSelectionArray was auto-generated from WSDL.
type RoomsSelectionArray struct {
	Items []*RoomsSelection `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SearchProductByCalendarRequest was auto-generated from WSDL.
type SearchProductByCalendarRequest struct {
	Credentials  *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Month        *string            `xml:"Month,omitempty" json:"Month,omitempty" yaml:"Month,omitempty"`
	Year         *string            `xml:"Year,omitempty" json:"Year,omitempty" yaml:"Year,omitempty"`
	ProdCode     *string            `xml:"ProdCode,omitempty" json:"ProdCode,omitempty" yaml:"ProdCode,omitempty"`
	ProdCodeDesc *string            `xml:"ProdCodeDesc,omitempty" json:"ProdCodeDesc,omitempty" yaml:"ProdCodeDesc,omitempty"`
	ProdCodeName *string            `xml:"ProdCodeName,omitempty" json:"ProdCodeName,omitempty" yaml:"ProdCodeName,omitempty"`
	DepLocal     *string            `xml:"DepLocal,omitempty" json:"DepLocal,omitempty" yaml:"DepLocal,omitempty"`
}

// SearchProductByCalendarResponse was auto-generated from WSDL.
type SearchProductByCalendarResponse struct {
	TotalProducts          *string                      `xml:"TotalProducts,omitempty" json:"TotalProducts,omitempty" yaml:"TotalProducts,omitempty"`
	SearchStatus           *string                      `xml:"SearchStatus,omitempty" json:"SearchStatus,omitempty" yaml:"SearchStatus,omitempty"`
	TotalProductsPerDay    *TotalProductsPerDayArray    `xml:"TotalProductsPerDay,omitempty" json:"TotalProductsPerDay,omitempty" yaml:"TotalProductsPerDay,omitempty"`
	ProductArrayByCalendar *ProductArrayByCalendarArray `xml:"ProductArrayByCalendar,omitempty" json:"ProductArrayByCalendar,omitempty" yaml:"ProductArrayByCalendar,omitempty"`
}

// SearchProductMasterData was auto-generated from WSDL.
type SearchProductMasterData struct {
	ContinentsArray  *ContinentArray   `xml:"ContinentsArray,omitempty" json:"ContinentsArray,omitempty" yaml:"ContinentsArray,omitempty"`
	CountriesArray   *CountryArray     `xml:"CountriesArray,omitempty" json:"CountriesArray,omitempty" yaml:"CountriesArray,omitempty"`
	LocationsArray   *LocationArray    `xml:"LocationsArray,omitempty" json:"LocationsArray,omitempty" yaml:"LocationsArray,omitempty"`
	ZonesArray       *ZoneArray        `xml:"ZonesArray,omitempty" json:"ZonesArray,omitempty" yaml:"ZonesArray,omitempty"`
	ProductTypeArray *ProductTypeArray `xml:"ProductTypeArray,omitempty" json:"ProductTypeArray,omitempty" yaml:"ProductTypeArray,omitempty"`
}

// SearchProductMasterDataRequest was auto-generated from WSDL.
type SearchProductMasterDataRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
}

// SearchProductMasterDataResponse was auto-generated from WSDL.
type SearchProductMasterDataResponse struct {
	SearchProductMasterDataArray *SearchProductMasterData `xml:"SearchProductMasterDataArray,omitempty" json:"SearchProductMasterDataArray,omitempty" yaml:"SearchProductMasterDataArray,omitempty"`
}

// SearchProductRequest was auto-generated from WSDL.
type SearchProductRequest struct {
	Credentials   *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Name          *string            `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Continent     *string            `xml:"Continent,omitempty" json:"Continent,omitempty" yaml:"Continent,omitempty"`
	Country       *string            `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
	Location      *string            `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	Zone          *string            `xml:"Zone,omitempty" json:"Zone,omitempty" yaml:"Zone,omitempty"`
	DepDate       *string            `xml:"DepDate,omitempty" json:"DepDate,omitempty" yaml:"DepDate,omitempty"`
	DepDateOffset *string            `xml:"DepDateOffset,omitempty" json:"DepDateOffset,omitempty" yaml:"DepDateOffset,omitempty"`
	ProdDateFrom  *string            `xml:"ProdDateFrom,omitempty" json:"ProdDateFrom,omitempty" yaml:"ProdDateFrom,omitempty"`
	Type          *string            `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Pkt_type      *string            `xml:"Pkt_type,omitempty" json:"Pkt_type,omitempty" yaml:"Pkt_type,omitempty"`
	PriceFrom     *string            `xml:"PriceFrom,omitempty" json:"PriceFrom,omitempty" yaml:"PriceFrom,omitempty"`
	PriceTo       *string            `xml:"PriceTo,omitempty" json:"PriceTo,omitempty" yaml:"PriceTo,omitempty"`
	Promo         *string            `xml:"Promo,omitempty" json:"Promo,omitempty" yaml:"Promo,omitempty"`
	Status_b2b    *string            `xml:"Status_b2b,omitempty" json:"Status_b2b,omitempty" yaml:"Status_b2b,omitempty"`
	Status_b2c    *string            `xml:"Status_b2c,omitempty" json:"Status_b2c,omitempty" yaml:"Status_b2c,omitempty"`
	Market        *string            `xml:"Market,omitempty" json:"Market,omitempty" yaml:"Market,omitempty"`
}

// SearchProductResponse was auto-generated from WSDL.
type SearchProductResponse struct {
	TotalProducts *string       `xml:"TotalProducts,omitempty" json:"TotalProducts,omitempty" yaml:"TotalProducts,omitempty"`
	ProductArray  *ProductArray `xml:"ProductArray,omitempty" json:"ProductArray,omitempty" yaml:"ProductArray,omitempty"`
}

// SearchProductTypesRequest was auto-generated from WSDL.
type SearchProductTypesRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
}

// SearchProductTypesResponse was auto-generated from WSDL.
type SearchProductTypesResponse struct {
	ProductTypeArray *ProductTypeArray `xml:"ProductTypeArray,omitempty" json:"ProductTypeArray,omitempty" yaml:"ProductTypeArray,omitempty"`
}

// SeniorStruct was auto-generated from WSDL.
type SeniorStruct struct {
	Min *string `xml:"Min,omitempty" json:"Min,omitempty" yaml:"Min,omitempty"`
	Max *string `xml:"Max,omitempty" json:"Max,omitempty" yaml:"Max,omitempty"`
}

// Service was auto-generated from WSDL.
type Service_ struct {
	Type        *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Desc        *string `xml:"Desc,omitempty" json:"Desc,omitempty" yaml:"Desc,omitempty"`
	Origin      *string `xml:"Origin,omitempty" json:"Origin,omitempty" yaml:"Origin,omitempty"`
	Destination *string `xml:"Destination,omitempty" json:"Destination,omitempty" yaml:"Destination,omitempty"`
	Time1       *string `xml:"Time1,omitempty" json:"Time1,omitempty" yaml:"Time1,omitempty"`
	Time2       *string `xml:"Time2,omitempty" json:"Time2,omitempty" yaml:"Time2,omitempty"`
	Data1       *string `xml:"Data1,omitempty" json:"Data1,omitempty" yaml:"Data1,omitempty"`
	Data2       *string `xml:"Data2,omitempty" json:"Data2,omitempty" yaml:"Data2,omitempty"`
	Quant       *string `xml:"Quant,omitempty" json:"Quant,omitempty" yaml:"Quant,omitempty"`
	Status      *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// ServiceArray was auto-generated from WSDL.
type ServiceArray struct {
	Items []*Service_ `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SetBookingOfflineRequest was auto-generated from WSDL.
type SetBookingOfflineRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	Mode        *string            `xml:"Mode,omitempty" json:"Mode,omitempty" yaml:"Mode,omitempty"`
}

// SetBookingOfflineResponse was auto-generated from WSDL.
type SetBookingOfflineResponse struct {
	Status  *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Message *string `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
}

// SetPaxAdditionalInfoRequest was auto-generated from WSDL.
type SetPaxAdditionalInfoRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	PaxInfo     *PaxInfoArray      `xml:"PaxInfo,omitempty" json:"PaxInfo,omitempty" yaml:"PaxInfo,omitempty"`
}

// SetPaxAdditionalInfoResponse was auto-generated from WSDL.
type SetPaxAdditionalInfoResponse struct {
	Reserve *string `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	Status  *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// SetReserveHotelsRequest was auto-generated from WSDL.
type SetReserveHotelsRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	Hotels      *HotelsArray       `xml:"Hotels,omitempty" json:"Hotels,omitempty" yaml:"Hotels,omitempty"`
}

// SetReserveHotelsResponse was auto-generated from WSDL.
type SetReserveHotelsResponse struct {
	Reserve *string `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	Status  *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Message *string `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
}

// SimulCancelReserveRequest was auto-generated from WSDL.
type SimulCancelReserveRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"reserve,omitempty" json:"reserve,omitempty" yaml:"reserve,omitempty"`
}

// SimulCancelReserveResponse was auto-generated from WSDL.
type SimulCancelReserveResponse struct {
	Status         *string              `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	ErrorsMsg      *ErrorsMsgArray      `xml:"ErrorsMsg,omitempty" json:"ErrorsMsg,omitempty" yaml:"ErrorsMsg,omitempty"`
	WarningsMsg    *WarningsMsgArray    `xml:"WarningsMsg,omitempty" json:"WarningsMsg,omitempty" yaml:"WarningsMsg,omitempty"`
	PoliciesDetail *PoliciesDetailArray `xml:"PoliciesDetail,omitempty" json:"PoliciesDetail,omitempty" yaml:"PoliciesDetail,omitempty"`
	Errors         *ErrorStruct         `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
}

// StaticCreateReserveRequest was auto-generated from WSDL.
type StaticCreateReserveRequest struct {
	Credentials      *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	PaxResArray      *PaxResArray       `xml:"PaxResArray,omitempty" json:"PaxResArray,omitempty" yaml:"PaxResArray,omitempty"`
	PaxResPaxs       *PaxResPaxs        `xml:"PaxResPaxs,omitempty" json:"PaxResPaxs,omitempty" yaml:"PaxResPaxs,omitempty"`
	Seniors          *string            `xml:"Seniors,omitempty" json:"Seniors,omitempty" yaml:"Seniors,omitempty"`
	Clientref        *string            `xml:"Clientref,omitempty" json:"Clientref,omitempty" yaml:"Clientref,omitempty"`
	Clientsobs       *string            `xml:"Clientsobs,omitempty" json:"Clientsobs,omitempty" yaml:"Clientsobs,omitempty"`
	SessionHash      *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	PaymentMethod    *string            `xml:"PaymentMethod,omitempty" json:"PaymentMethod,omitempty" yaml:"PaymentMethod,omitempty"`
	TotalPrice       *string            `xml:"TotalPrice,omitempty" json:"TotalPrice,omitempty" yaml:"TotalPrice,omitempty"`
	Fatura           *string            `xml:"Fatura,omitempty" json:"Fatura,omitempty" yaml:"Fatura,omitempty"`
	FaturaStruct     *FaturaStruct      `xml:"FaturaStruct,omitempty" json:"FaturaStruct,omitempty" yaml:"FaturaStruct,omitempty"`
	PromoCode        *string            `xml:"PromoCode,omitempty" json:"PromoCode,omitempty" yaml:"PromoCode,omitempty"`
	Discount         *string            `xml:"Discount,omitempty" json:"Discount,omitempty" yaml:"Discount,omitempty"`
	Optionals        *string            `xml:"Optionals,omitempty" json:"Optionals,omitempty" yaml:"Optionals,omitempty"`
	ReserveToAdd     *string            `xml:"ReserveToAdd,omitempty" json:"ReserveToAdd,omitempty" yaml:"ReserveToAdd,omitempty"`
	ExtraNightsStart *string            `xml:"ExtraNightsStart,omitempty" json:"ExtraNightsStart,omitempty" yaml:"ExtraNightsStart,omitempty"`
	ExtraNightsEnd   *string            `xml:"ExtraNightsEnd,omitempty" json:"ExtraNightsEnd,omitempty" yaml:"ExtraNightsEnd,omitempty"`
}

// StaticCreateReserveResponse was auto-generated from WSDL.
type StaticCreateReserveResponse struct {
	Reserve      *string       `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	Status       *string       `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	PagmaticInfo *PagmaticInfo `xml:"PagmaticInfo,omitempty" json:"PagmaticInfo,omitempty" yaml:"PagmaticInfo,omitempty"`
	FaturaStruct *FaturaStruct `xml:"FaturaStruct,omitempty" json:"FaturaStruct,omitempty" yaml:"FaturaStruct,omitempty"`
}

// StaticGetPaxsStructRequest was auto-generated from WSDL.
type StaticGetPaxsStructRequest struct {
	Credentials    *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	OptionID       *string            `xml:"OptionID,omitempty" json:"OptionID,omitempty" yaml:"OptionID,omitempty"`
	DepartureDate  *string            `xml:"DepartureDate,omitempty" json:"DepartureDate,omitempty" yaml:"DepartureDate,omitempty"`
	DepartureLocal *string            `xml:"DepartureLocal,omitempty" json:"DepartureLocal,omitempty" yaml:"DepartureLocal,omitempty"`
	PaxArray       *string            `xml:"PaxArray,omitempty" json:"PaxArray,omitempty" yaml:"PaxArray,omitempty"`
	SessionHash    *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
}

// StaticGetPaxsStructResponse was auto-generated from WSDL.
type StaticGetPaxsStructResponse struct {
	Session_ID *string   `xml:"Session_ID,omitempty" json:"Session_ID,omitempty" yaml:"Session_ID,omitempty"`
	PaxArray   *PaxArray `xml:"PaxArray,omitempty" json:"PaxArray,omitempty" yaml:"PaxArray,omitempty"`
}

// StaticGetProductParametersRequest was auto-generated from WSDL.
type StaticGetProductParametersRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	ProductCode *string            `xml:"ProductCode,omitempty" json:"ProductCode,omitempty" yaml:"ProductCode,omitempty"`
	Month       *string            `xml:"Month,omitempty" json:"Month,omitempty" yaml:"Month,omitempty"`
	Year        *string            `xml:"Year,omitempty" json:"Year,omitempty" yaml:"Year,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
}

// StaticGetProductParametersResponse was auto-generated from WSDL.
type StaticGetProductParametersResponse struct {
	Code                        *string                      `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name                        *string                      `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	ProdType                    *string                      `xml:"ProdType,omitempty" json:"ProdType,omitempty" yaml:"ProdType,omitempty"`
	Session_ID                  *string                      `xml:"Session_ID,omitempty" json:"Session_ID,omitempty" yaml:"Session_ID,omitempty"`
	ResStatus                   *string                      `xml:"ResStatus,omitempty" json:"ResStatus,omitempty" yaml:"ResStatus,omitempty"`
	ResStatusDesc               *string                      `xml:"ResStatusDesc,omitempty" json:"ResStatusDesc,omitempty" yaml:"ResStatusDesc,omitempty"`
	ChildAgesArray              *ChildStruct                 `xml:"ChildAgesArray,omitempty" json:"ChildAgesArray,omitempty" yaml:"ChildAgesArray,omitempty"`
	InfantAgesArray             *InfantStruct                `xml:"InfantAgesArray,omitempty" json:"InfantAgesArray,omitempty" yaml:"InfantAgesArray,omitempty"`
	SeniorArray                 *SeniorStruct                `xml:"SeniorArray,omitempty" json:"SeniorArray,omitempty" yaml:"SeniorArray,omitempty"`
	ReductionNightsArray        *ReductionNights             `xml:"ReductionNightsArray,omitempty" json:"ReductionNightsArray,omitempty" yaml:"ReductionNightsArray,omitempty"`
	MultimediaContentTextsArray *MultimediaContentTextsArray `xml:"MultimediaContentTextsArray,omitempty" json:"MultimediaContentTextsArray,omitempty" yaml:"MultimediaContentTextsArray,omitempty"`
	OptionArray                 *OptionArray                 `xml:"OptionArray,omitempty" json:"OptionArray,omitempty" yaml:"OptionArray,omitempty"`
}

// StaticGetReserveFlightsRequest was auto-generated from WSDL.
type StaticGetReserveFlightsRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
}

// StaticGetReserveFlightsResponse was auto-generated from WSDL.
type StaticGetReserveFlightsResponse struct {
	Paxs *PaxsArray `xml:"Paxs,omitempty" json:"Paxs,omitempty" yaml:"Paxs,omitempty"`
}

// StaticGetReserveHotelsRequest was auto-generated from WSDL.
type StaticGetReserveHotelsRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
}

// StaticGetReserveHotelsResponse was auto-generated from WSDL.
type StaticGetReserveHotelsResponse struct {
	Reserve *string      `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	Error   *string      `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
	Hotels  *HotelsArray `xml:"Hotels,omitempty" json:"Hotels,omitempty" yaml:"Hotels,omitempty"`
}

// StaticGetReserveRoomingRequest was auto-generated from WSDL.
type StaticGetReserveRoomingRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
}

// StaticGetReserveRoomingResponse was auto-generated from WSDL.
type StaticGetReserveRoomingResponse struct {
	PaxRoom *PaxRoomArray `xml:"PaxRoom,omitempty" json:"PaxRoom,omitempty" yaml:"PaxRoom,omitempty"`
}

// StaticGetRoomsStructRequest was auto-generated from WSDL.
type StaticGetRoomsStructRequest struct {
	Credentials     *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Option_id       *string            `xml:"Option_id,omitempty" json:"Option_id,omitempty" yaml:"Option_id,omitempty"`
	Departure_Date  *string            `xml:"Departure_Date,omitempty" json:"Departure_Date,omitempty" yaml:"Departure_Date,omitempty"`
	Departure_Local *string            `xml:"Departure_Local,omitempty" json:"Departure_Local,omitempty" yaml:"Departure_Local,omitempty"`
	SessionHash     *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
}

// StaticGetRoomsStructResponse was auto-generated from WSDL.
type StaticGetRoomsStructResponse struct {
	Session_ID *string    `xml:"Session_ID,omitempty" json:"Session_ID,omitempty" yaml:"Session_ID,omitempty"`
	RoomArray  *RoomArray `xml:"RoomArray,omitempty" json:"RoomArray,omitempty" yaml:"RoomArray,omitempty"`
}

// StaticGetSimulationRequest was auto-generated from WSDL.
type StaticGetSimulationRequest struct {
	Credentials         *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	OptionID            *string            `xml:"OptionID,omitempty" json:"OptionID,omitempty" yaml:"OptionID,omitempty"`
	DepartureDate       *string            `xml:"DepartureDate,omitempty" json:"DepartureDate,omitempty" yaml:"DepartureDate,omitempty"`
	DepartureLocal      *string            `xml:"DepartureLocal,omitempty" json:"DepartureLocal,omitempty" yaml:"DepartureLocal,omitempty"`
	SubtractNightsStart *string            `xml:"SubtractNightsStart,omitempty" json:"SubtractNightsStart,omitempty" yaml:"SubtractNightsStart,omitempty"`
	SubtractNightsEnd   *string            `xml:"SubtractNightsEnd,omitempty" json:"SubtractNightsEnd,omitempty" yaml:"SubtractNightsEnd,omitempty"`
	ExtraNightsStart    *string            `xml:"ExtraNightsStart,omitempty" json:"ExtraNightsStart,omitempty" yaml:"ExtraNightsStart,omitempty"`
	ExtraNightsEnd      *string            `xml:"ExtraNightsEnd,omitempty" json:"ExtraNightsEnd,omitempty" yaml:"ExtraNightsEnd,omitempty"`
	Seniors             *string            `xml:"Seniors,omitempty" json:"Seniors,omitempty" yaml:"Seniors,omitempty"`
	Rooms               *RoomsArray        `xml:"Rooms,omitempty" json:"Rooms,omitempty" yaml:"Rooms,omitempty"`
	Paxs                *PaxsSimulation    `xml:"Paxs,omitempty" json:"Paxs,omitempty" yaml:"Paxs,omitempty"`
	SessionHash         *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
	ReserveToAdd        *string            `xml:"ReserveToAdd,omitempty" json:"ReserveToAdd,omitempty" yaml:"ReserveToAdd,omitempty"`
}

// StaticGetSimulationResponse was auto-generated from WSDL.
type StaticGetSimulationResponse struct {
	Session_ID      *string              `xml:"Session_ID,omitempty" json:"Session_ID,omitempty" yaml:"Session_ID,omitempty"`
	Success         *string              `xml:"Success,omitempty" json:"Success,omitempty" yaml:"Success,omitempty"`
	Errors          *ErrorsArray         `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`
	AvailableStatus *string              `xml:"AvailableStatus,omitempty" json:"AvailableStatus,omitempty" yaml:"AvailableStatus,omitempty"`
	TotalPrice      *string              `xml:"TotalPrice,omitempty" json:"TotalPrice,omitempty" yaml:"TotalPrice,omitempty"`
	ServiceArray    *ServiceArray        `xml:"ServiceArray,omitempty" json:"ServiceArray,omitempty" yaml:"ServiceArray,omitempty"`
	PriceArray      *PriceArray          `xml:"PriceArray,omitempty" json:"PriceArray,omitempty" yaml:"PriceArray,omitempty"`
	PaymentMethods  *PaymentMethodsArray `xml:"PaymentMethods,omitempty" json:"PaymentMethods,omitempty" yaml:"PaymentMethods,omitempty"`
	Fatura          *FaturaStruct        `xml:"Fatura,omitempty" json:"Fatura,omitempty" yaml:"Fatura,omitempty"`
}

// StaticSetReserveExtraNightsRequest was auto-generated from WSDL.
type StaticSetReserveExtraNightsRequest struct {
	Credentials    *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve        *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	TypeExtraNight *string            `xml:"TypeExtraNight,omitempty" json:"TypeExtraNight,omitempty" yaml:"TypeExtraNight,omitempty"`
	Status         *string            `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// StaticSetReserveExtraNightsResponse was auto-generated from
// WSDL.
type StaticSetReserveExtraNightsResponse struct {
	Status *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Error  *string `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
}

// StaticSetReserveFlightsRequest was auto-generated from WSDL.
type StaticSetReserveFlightsRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	Reserve     *string            `xml:"Reserve,omitempty" json:"Reserve,omitempty" yaml:"Reserve,omitempty"`
	Paxs        *PaxsArray         `xml:"Paxs,omitempty" json:"Paxs,omitempty" yaml:"Paxs,omitempty"`
}

// StaticSetReserveFlightsResponse was auto-generated from WSDL.
type StaticSetReserveFlightsResponse struct {
	PaxsStatus *PaxsStatusArray `xml:"PaxsStatus,omitempty" json:"PaxsStatus,omitempty" yaml:"PaxsStatus,omitempty"`
}

// StaticSetReserveRoomingRequest was auto-generated from WSDL.
type StaticSetReserveRoomingRequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	PaxRoom     *PaxRoomArray      `xml:"PaxRoom,omitempty" json:"PaxRoom,omitempty" yaml:"PaxRoom,omitempty"`
}

// StaticSetReserveRoomingResponse was auto-generated from WSDL.
type StaticSetReserveRoomingResponse struct {
	Status *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	Error  *string `xml:"Error,omitempty" json:"Error,omitempty" yaml:"Error,omitempty"`
}

// Stores was auto-generated from WSDL.
type Stores struct {
	Id     *string `xml:"Id,omitempty" json:"Id,omitempty" yaml:"Id,omitempty"`
	Name   *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Status *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// StoresArray was auto-generated from WSDL.
type StoresArray struct {
	Items []*Stores `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// TextContent was auto-generated from WSDL.
type TextContent struct {
	Type    *string `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	Content *string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

// TextContentArray was auto-generated from WSDL.
type TextContentArray struct {
	Items []*TextContent `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// TotalProductsPerDay was auto-generated from WSDL.
type TotalProductsPerDay struct {
	Day           *string `xml:"Day,omitempty" json:"Day,omitempty" yaml:"Day,omitempty"`
	TotalProducts *string `xml:"TotalProducts,omitempty" json:"TotalProducts,omitempty" yaml:"TotalProducts,omitempty"`
}

// TotalProductsPerDayArray was auto-generated from WSDL.
type TotalProductsPerDayArray struct {
	Items []*TotalProductsPerDay `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Upgrades was auto-generated from WSDL.
type Upgrades struct {
	ID          *string `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Code        *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Sellvalue   *string `xml:"Sellvalue,omitempty" json:"Sellvalue,omitempty" yaml:"Sellvalue,omitempty"`
}

// UpgradesArray was auto-generated from WSDL.
type UpgradesArray struct {
	Items []*Upgrades `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ValidateVISARequest was auto-generated from WSDL.
type ValidateVISARequest struct {
	Credentials *CredentialsStruct `xml:"Credentials,omitempty" json:"Credentials,omitempty" yaml:"Credentials,omitempty"`
	SessionHash *string            `xml:"SessionHash,omitempty" json:"SessionHash,omitempty" yaml:"SessionHash,omitempty"`
}

// ValidateVISAResponse was auto-generated from WSDL.
type ValidateVISAResponse struct {
	Status *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// WarningsMsg was auto-generated from WSDL.
type WarningsMsg struct {
	Warning *string `xml:"Warning,omitempty" json:"Warning,omitempty" yaml:"Warning,omitempty"`
}

// WarningsMsgArray was auto-generated from WSDL.
type WarningsMsgArray struct {
	Items []*WarningsMsg `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Zone was auto-generated from WSDL.
type Zone struct {
	Code        *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Location    *string `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
}

// ZoneArray was auto-generated from WSDL.
type ZoneArray struct {
	Items []*Zone `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Zones was auto-generated from WSDL.
type Zones struct {
	Code *string `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Name *string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

// ZonesArray was auto-generated from WSDL.
type ZonesArray struct {
	Items []*Zones `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// StringArray was auto-generated from WSDL.
type StringArray struct {
	Items []string `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Operation wrapper for CancelReserve.
// OperationCancelReserveSoapIn was auto-generated from WSDL.
type OperationCancelReserveSoapIn struct {
	Input_arr *CancelReserveRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for CancelReserve.
// OperationCancelReserveSoapOut was auto-generated from WSDL.
type OperationCancelReserveSoapOut struct {
	Return *CancelReserveResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynCreateReserve.
// OperationDynCreateReserveSoapIn was auto-generated from WSDL.
type OperationDynCreateReserveSoapIn struct {
	Input_arr *DynCreateReserveRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynCreateReserve.
// OperationDynCreateReserveSoapOut was auto-generated from WSDL.
type OperationDynCreateReserveSoapOut struct {
	Return *DynCreateReserveResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynGetOptionalsSelected.
// OperationDynGetOptionalsSelectedSoapIn was auto-generated from
// WSDL.
type OperationDynGetOptionalsSelectedSoapIn struct {
	Input_arr *DynGetOptionalsSelectedRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynGetOptionalsSelected.
// OperationDynGetOptionalsSelectedSoapOut was auto-generated from
// WSDL.
type OperationDynGetOptionalsSelectedSoapOut struct {
	Return *DynGetOptionalsSelectedResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynGetProductParameters.
// OperationDynGetProductParametersSoapIn was auto-generated from
// WSDL.
type OperationDynGetProductParametersSoapIn struct {
	Input_arr *DynProductParametersRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynGetProductParameters.
// OperationDynGetProductParametersSoapOut was auto-generated from
// WSDL.
type OperationDynGetProductParametersSoapOut struct {
	Return *DynProductParametersResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynGetSimulation.
// OperationDynGetSimulationSoapIn was auto-generated from WSDL.
type OperationDynGetSimulationSoapIn struct {
	Input_arr *DynGetSimulationRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynGetSimulation.
// OperationDynGetSimulationSoapOut was auto-generated from WSDL.
type OperationDynGetSimulationSoapOut struct {
	Return *DynGetSimulationResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynPreBooking.
// OperationDynPreBookingSoapIn was auto-generated from WSDL.
type OperationDynPreBookingSoapIn struct {
	Input_arr *DynPreBookingRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynPreBooking.
// OperationDynPreBookingSoapOut was auto-generated from WSDL.
type OperationDynPreBookingSoapOut struct {
	Return *DynPreBookingResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynPriceAvailResume.
// OperationDynPriceAvailResumeSoapIn was auto-generated from WSDL.
type OperationDynPriceAvailResumeSoapIn struct {
	Input_arr *DynPriceAvailResumeRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynPriceAvailResume.
// OperationDynPriceAvailResumeSoapOut was auto-generated from
// WSDL.
type OperationDynPriceAvailResumeSoapOut struct {
	Return *DynPriceAvailResumeResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynProductCache.
// OperationDynProductCacheSoapIn was auto-generated from WSDL.
type OperationDynProductCacheSoapIn struct {
	Input_arr *DynProductCacheRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynProductCache.
// OperationDynProductCacheSoapOut was auto-generated from WSDL.
type OperationDynProductCacheSoapOut struct {
	Return *DynProductCacheResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynProductOptionals.
// OperationDynProductOptionalsSoapIn was auto-generated from WSDL.
type OperationDynProductOptionalsSoapIn struct {
	Input_arr *DynProductOptionalsRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynProductOptionals.
// OperationDynProductOptionalsSoapOut was auto-generated from
// WSDL.
type OperationDynProductOptionalsSoapOut struct {
	Return *DynProductOptionalsResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynRemoveOptionals.
// OperationDynRemoveOptionalsSoapIn was auto-generated from WSDL.
type OperationDynRemoveOptionalsSoapIn struct {
	Input_arr *DynRemoveOptionalsRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynRemoveOptionals.
// OperationDynRemoveOptionalsSoapOut was auto-generated from WSDL.
type OperationDynRemoveOptionalsSoapOut struct {
	Return *DynRemoveOptionalsResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynSearchProductAvailableServices.
// OperationDynSearchProductAvailableServicesSoapIn was auto-generated
// from WSDL.
type OperationDynSearchProductAvailableServicesSoapIn struct {
	Input_arr *DynProductAvailableServicesRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynSearchProductAvailableServices.
// OperationDynSearchProductAvailableServicesSoapOut was auto-generated
// from WSDL.
type OperationDynSearchProductAvailableServicesSoapOut struct {
	Return *DynProductAvailableServicesResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for DynSetServicesSelected.
// OperationDynSetServicesSelectedSoapIn was auto-generated from
// WSDL.
type OperationDynSetServicesSelectedSoapIn struct {
	Input_arr *DynServicesSelectedRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for DynSetServicesSelected.
// OperationDynSetServicesSelectedSoapOut was auto-generated from
// WSDL.
type OperationDynSetServicesSelectedSoapOut struct {
	Return *DynServicesSelectedResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for GetClients.
// OperationGetClientsSoapIn was auto-generated from WSDL.
type OperationGetClientsSoapIn struct {
	Input_arr *GetClientsRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for GetClients.
// OperationGetClientsSoapOut was auto-generated from WSDL.
type OperationGetClientsSoapOut struct {
	Return *GetClientsResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for GetHotelDetails.
// OperationGetHotelDetailsSoapIn was auto-generated from WSDL.
type OperationGetHotelDetailsSoapIn struct {
	Input_arr *GetHotelDetailsRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for GetHotelDetails.
// OperationGetHotelDetailsSoapOut was auto-generated from WSDL.
type OperationGetHotelDetailsSoapOut struct {
	Return *GetHotelDetailsResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for GetInfoPax.
// OperationGetInfoPaxSoapIn was auto-generated from WSDL.
type OperationGetInfoPaxSoapIn struct {
	Input_arr *GetPaxInfoRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for GetInfoPax.
// OperationGetInfoPaxSoapOut was auto-generated from WSDL.
type OperationGetInfoPaxSoapOut struct {
	Return *GetPaxInfoResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for GetPaxAdditionalInfo.
// OperationGetPaxAdditionalInfoSoapIn was auto-generated from
// WSDL.
type OperationGetPaxAdditionalInfoSoapIn struct {
	Input_arr *GetPaxAdditionalInfoRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for GetPaxAdditionalInfo.
// OperationGetPaxAdditionalInfoSoapOut was auto-generated from
// WSDL.
type OperationGetPaxAdditionalInfoSoapOut struct {
	Return *GetPaxAdditionalInfoResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for GetStores.
// OperationGetStoresSoapIn was auto-generated from WSDL.
type OperationGetStoresSoapIn struct {
	Input_arr *GetStoresRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for GetStores.
// OperationGetStoresSoapOut was auto-generated from WSDL.
type OperationGetStoresSoapOut struct {
	Return *GetStoresResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for ReserveList.
// OperationReserveListSoapIn was auto-generated from WSDL.
type OperationReserveListSoapIn struct {
	Input_arr *ReserveListRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for ReserveList.
// OperationReserveListSoapOut was auto-generated from WSDL.
type OperationReserveListSoapOut struct {
	Return *ReserveListResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for RetrieveReserve.
// OperationRetrieveReserveSoapIn was auto-generated from WSDL.
type OperationRetrieveReserveSoapIn struct {
	Input_arr *RetrieveReserveRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for RetrieveReserve.
// OperationRetrieveReserveSoapOut was auto-generated from WSDL.
type OperationRetrieveReserveSoapOut struct {
	Return *RetrieveReserveResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SearchProductByCalendar.
// OperationSearchProductByCalendarSoapIn was auto-generated from
// WSDL.
type OperationSearchProductByCalendarSoapIn struct {
	Input_arr *SearchProductByCalendarRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for SearchProductByCalendar.
// OperationSearchProductByCalendarSoapOut was auto-generated from
// WSDL.
type OperationSearchProductByCalendarSoapOut struct {
	Return *SearchProductByCalendarResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SearchProductMasterData.
// OperationSearchProductMasterDataSoapIn was auto-generated from
// WSDL.
type OperationSearchProductMasterDataSoapIn struct {
	Input_arr *SearchProductMasterDataRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for SearchProductMasterData.
// OperationSearchProductMasterDataSoapOut was auto-generated from
// WSDL.
type OperationSearchProductMasterDataSoapOut struct {
	Return *SearchProductMasterDataResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SearchProductTypes.
// OperationSearchProductTypesSoapIn was auto-generated from WSDL.
type OperationSearchProductTypesSoapIn struct {
	Input_arr *SearchProductTypesRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for SearchProductTypes.
// OperationSearchProductTypesSoapOut was auto-generated from WSDL.
type OperationSearchProductTypesSoapOut struct {
	Return *SearchProductTypesResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SearchProducts.
// OperationSearchProductsSoapIn was auto-generated from WSDL.
type OperationSearchProductsSoapIn struct {
	Input_arr *SearchProductRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for SearchProducts.
// OperationSearchProductsSoapOut was auto-generated from WSDL.
type OperationSearchProductsSoapOut struct {
	Return *SearchProductResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SetBookingOffline.
// OperationSetBookingOfflineSoapIn was auto-generated from WSDL.
type OperationSetBookingOfflineSoapIn struct {
	Input_arr *SetBookingOfflineRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for SetBookingOffline.
// OperationSetBookingOfflineSoapOut was auto-generated from WSDL.
type OperationSetBookingOfflineSoapOut struct {
	Return *SetBookingOfflineResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SetPaxAdditionalInfo.
// OperationSetPaxAdditionalInfoSoapIn was auto-generated from
// WSDL.
type OperationSetPaxAdditionalInfoSoapIn struct {
	Input_arr *SetPaxAdditionalInfoRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for SetPaxAdditionalInfo.
// OperationSetPaxAdditionalInfoSoapOut was auto-generated from
// WSDL.
type OperationSetPaxAdditionalInfoSoapOut struct {
	Return *SetPaxAdditionalInfoResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SetReserveHotels.
// OperationSetReserveHotelsSoapIn was auto-generated from WSDL.
type OperationSetReserveHotelsSoapIn struct {
	Input_arr *SetReserveHotelsRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for SetReserveHotels.
// OperationSetReserveHotelsSoapOut was auto-generated from WSDL.
type OperationSetReserveHotelsSoapOut struct {
	Return *SetReserveHotelsResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SimulCancelReserve.
// OperationSimulCancelReserveSoapIn was auto-generated from WSDL.
type OperationSimulCancelReserveSoapIn struct {
	Input_arr *SimulCancelReserveRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for SimulCancelReserve.
// OperationSimulCancelReserveSoapOut was auto-generated from WSDL.
type OperationSimulCancelReserveSoapOut struct {
	Return *SimulCancelReserveResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticCreateReserve.
// OperationStaticCreateReserveSoapIn was auto-generated from WSDL.
type OperationStaticCreateReserveSoapIn struct {
	Input_arr *StaticCreateReserveRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticCreateReserve.
// OperationStaticCreateReserveSoapOut was auto-generated from
// WSDL.
type OperationStaticCreateReserveSoapOut struct {
	Return *StaticCreateReserveResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticGetPaxsStruct.
// OperationStaticGetPaxsStructSoapIn was auto-generated from WSDL.
type OperationStaticGetPaxsStructSoapIn struct {
	Input_arr *StaticGetPaxsStructRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticGetPaxsStruct.
// OperationStaticGetPaxsStructSoapOut was auto-generated from
// WSDL.
type OperationStaticGetPaxsStructSoapOut struct {
	Return *StaticGetPaxsStructResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticGetProductParameters.
// OperationStaticGetProductParametersSoapIn was auto-generated
// from WSDL.
type OperationStaticGetProductParametersSoapIn struct {
	Input_arr *StaticGetProductParametersRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticGetProductParameters.
// OperationStaticGetProductParametersSoapOut was auto-generated
// from WSDL.
type OperationStaticGetProductParametersSoapOut struct {
	Return *StaticGetProductParametersResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticGetReserveFlights.
// OperationStaticGetReserveFlightsSoapIn was auto-generated from
// WSDL.
type OperationStaticGetReserveFlightsSoapIn struct {
	Input_arr *StaticGetReserveFlightsRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticGetReserveFlights.
// OperationStaticGetReserveFlightsSoapOut was auto-generated from
// WSDL.
type OperationStaticGetReserveFlightsSoapOut struct {
	Return *StaticGetReserveFlightsResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticGetReserveHotels.
// OperationStaticGetReserveHotelsSoapIn was auto-generated from
// WSDL.
type OperationStaticGetReserveHotelsSoapIn struct {
	Input_arr *StaticGetReserveHotelsRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticGetReserveHotels.
// OperationStaticGetReserveHotelsSoapOut was auto-generated from
// WSDL.
type OperationStaticGetReserveHotelsSoapOut struct {
	Return *StaticGetReserveHotelsResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticGetReserveRooming.
// OperationStaticGetReserveRoomingSoapIn was auto-generated from
// WSDL.
type OperationStaticGetReserveRoomingSoapIn struct {
	Input_arr *StaticGetReserveRoomingRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticGetReserveRooming.
// OperationStaticGetReserveRoomingSoapOut was auto-generated from
// WSDL.
type OperationStaticGetReserveRoomingSoapOut struct {
	Return *StaticGetReserveRoomingResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticGetRoomsStruct.
// OperationStaticGetRoomsStructSoapIn was auto-generated from
// WSDL.
type OperationStaticGetRoomsStructSoapIn struct {
	Input_arr *StaticGetRoomsStructRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticGetRoomsStruct.
// OperationStaticGetRoomsStructSoapOut was auto-generated from
// WSDL.
type OperationStaticGetRoomsStructSoapOut struct {
	Return *StaticGetRoomsStructResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticGetSimulation.
// OperationStaticGetSimulationSoapIn was auto-generated from WSDL.
type OperationStaticGetSimulationSoapIn struct {
	Input_arr *StaticGetSimulationRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticGetSimulation.
// OperationStaticGetSimulationSoapOut was auto-generated from
// WSDL.
type OperationStaticGetSimulationSoapOut struct {
	Return *StaticGetSimulationResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticSetReserveExtraNights.
// OperationStaticSetReserveExtraNightsSoapIn was auto-generated
// from WSDL.
type OperationStaticSetReserveExtraNightsSoapIn struct {
	Input_arr *StaticSetReserveExtraNightsRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticSetReserveExtraNights.
// OperationStaticSetReserveExtraNightsSoapOut was auto-generated
// from WSDL.
type OperationStaticSetReserveExtraNightsSoapOut struct {
	Return *StaticSetReserveExtraNightsResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticSetReserveFlights.
// OperationStaticSetReserveFlightsSoapIn was auto-generated from
// WSDL.
type OperationStaticSetReserveFlightsSoapIn struct {
	Input_arr *StaticSetReserveFlightsRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticSetReserveFlights.
// OperationStaticSetReserveFlightsSoapOut was auto-generated from
// WSDL.
type OperationStaticSetReserveFlightsSoapOut struct {
	Return *StaticSetReserveFlightsResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for StaticSetReserveRooming.
// OperationStaticSetReserveRoomingSoapIn was auto-generated from
// WSDL.
type OperationStaticSetReserveRoomingSoapIn struct {
	Input_arr *StaticSetReserveRoomingRequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for StaticSetReserveRooming.
// OperationStaticSetReserveRoomingSoapOut was auto-generated from
// WSDL.
type OperationStaticSetReserveRoomingSoapOut struct {
	Return *StaticSetReserveRoomingResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for ValidateVISA.
// OperationValidateVISASoapIn was auto-generated from WSDL.
type OperationValidateVISASoapIn struct {
	Input_arr *ValidateVISARequest `xml:"input_arr,omitempty" json:"input_arr,omitempty" yaml:"input_arr,omitempty"`
}

// Operation wrapper for ValidateVISA.
// OperationValidateVISASoapOut was auto-generated from WSDL.
type OperationValidateVISASoapOut struct {
	Return *ValidateVISAResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// wbs_pkt_methodsSoap implements the Wbs_pkt_methodsSoap interface.
type wbs_pkt_methodsSoap struct {
	cli *soap.Client
}

// CancelReserve was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) CancelReserve(input_arr *CancelReserveRequest) (*CancelReserveResponse, error) {
	α := struct {
		M OperationCancelReserveSoapIn `xml:"tns:CancelReserve"`
	}{
		OperationCancelReserveSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationCancelReserveSoapOut `xml:"CancelReserveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/CancelReserve", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynCreateReserve was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynCreateReserve(input_arr *DynCreateReserveRequest) (*DynCreateReserveResponse, error) {
	α := struct {
		M OperationDynCreateReserveSoapIn `xml:"tns:DynCreateReserve"`
	}{
		OperationDynCreateReserveSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynCreateReserveSoapOut `xml:"DynCreateReserveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynCreateReserve", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynGetOptionalsSelected was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynGetOptionalsSelected(input_arr *DynGetOptionalsSelectedRequest) (*DynGetOptionalsSelectedResponse, error) {
	α := struct {
		M OperationDynGetOptionalsSelectedSoapIn `xml:"tns:DynGetOptionalsSelected"`
	}{
		OperationDynGetOptionalsSelectedSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynGetOptionalsSelectedSoapOut `xml:"DynGetOptionalsSelectedResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynGetOptionalsSelected", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynGetProductParameters was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynGetProductParameters(input_arr *DynProductParametersRequest) (*DynProductParametersResponse, error) {
	α := struct {
		M OperationDynGetProductParametersSoapIn `xml:"tns:DynGetProductParameters"`
	}{
		OperationDynGetProductParametersSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynGetProductParametersSoapOut `xml:"DynGetProductParametersResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynGetProductParameters", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynGetSimulation was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynGetSimulation(input_arr *DynGetSimulationRequest) (*DynGetSimulationResponse, error) {
	α := struct {
		M OperationDynGetSimulationSoapIn `xml:"tns:DynGetSimulation"`
	}{
		OperationDynGetSimulationSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynGetSimulationSoapOut `xml:"DynGetSimulationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynGetSimulation", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynPreBooking was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynPreBooking(input_arr *DynPreBookingRequest) (*DynPreBookingResponse, error) {
	α := struct {
		M OperationDynPreBookingSoapIn `xml:"tns:DynPreBooking"`
	}{
		OperationDynPreBookingSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynPreBookingSoapOut `xml:"DynPreBookingResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynPreBooking", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynPriceAvailResume was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynPriceAvailResume(input_arr *DynPriceAvailResumeRequest) (*DynPriceAvailResumeResponse, error) {
	α := struct {
		M OperationDynPriceAvailResumeSoapIn `xml:"tns:DynPriceAvailResume"`
	}{
		OperationDynPriceAvailResumeSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynPriceAvailResumeSoapOut `xml:"DynPriceAvailResumeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynPriceAvailResume", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynProductCache was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynProductCache(input_arr *DynProductCacheRequest) (*DynProductCacheResponse, error) {
	α := struct {
		M OperationDynProductCacheSoapIn `xml:"tns:DynProductCache"`
	}{
		OperationDynProductCacheSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynProductCacheSoapOut `xml:"DynProductCacheResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynProductCache", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynProductOptionals was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynProductOptionals(input_arr *DynProductOptionalsRequest) (*DynProductOptionalsResponse, error) {
	α := struct {
		M OperationDynProductOptionalsSoapIn `xml:"tns:DynProductOptionals"`
	}{
		OperationDynProductOptionalsSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynProductOptionalsSoapOut `xml:"DynProductOptionalsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynProductOptionals", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynRemoveOptionals was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynRemoveOptionals(input_arr *DynRemoveOptionalsRequest) (*DynRemoveOptionalsResponse, error) {
	α := struct {
		M OperationDynRemoveOptionalsSoapIn `xml:"tns:DynRemoveOptionals"`
	}{
		OperationDynRemoveOptionalsSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynRemoveOptionalsSoapOut `xml:"DynRemoveOptionalsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynRemoveOptionals", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynSearchProductAvailableServices was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynSearchProductAvailableServices(input_arr *DynProductAvailableServicesRequest) (*DynProductAvailableServicesResponse, error) {
	α := struct {
		M OperationDynSearchProductAvailableServicesSoapIn `xml:"tns:DynSearchProductAvailableServices"`
	}{
		OperationDynSearchProductAvailableServicesSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynSearchProductAvailableServicesSoapOut `xml:"DynSearchProductAvailableServicesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynSearchProductAvailableServices", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// DynSetServicesSelected was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) DynSetServicesSelected(input_arr *DynServicesSelectedRequest) (*DynServicesSelectedResponse, error) {
	α := struct {
		M OperationDynSetServicesSelectedSoapIn `xml:"tns:DynSetServicesSelected"`
	}{
		OperationDynSetServicesSelectedSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationDynSetServicesSelectedSoapOut `xml:"DynSetServicesSelectedResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/DynSetServicesSelected", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// GetClients was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) GetClients(input_arr *GetClientsRequest) (*GetClientsResponse, error) {
	α := struct {
		M OperationGetClientsSoapIn `xml:"tns:GetClients"`
	}{
		OperationGetClientsSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationGetClientsSoapOut `xml:"GetClientsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/GetClients", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// GetHotelDetails was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) GetHotelDetails(input_arr *GetHotelDetailsRequest) (*GetHotelDetailsResponse, error) {
	α := struct {
		M OperationGetHotelDetailsSoapIn `xml:"tns:GetHotelDetails"`
	}{
		OperationGetHotelDetailsSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationGetHotelDetailsSoapOut `xml:"GetHotelDetailsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/GetHotelDetails", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// GetInfoPax was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) GetInfoPax(input_arr *GetPaxInfoRequest) (*GetPaxInfoResponse, error) {
	α := struct {
		M OperationGetInfoPaxSoapIn `xml:"tns:GetInfoPax"`
	}{
		OperationGetInfoPaxSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationGetInfoPaxSoapOut `xml:"GetInfoPaxResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/GetInfoPax", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// GetPaxAdditionalInfo was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) GetPaxAdditionalInfo(input_arr *GetPaxAdditionalInfoRequest) (*GetPaxAdditionalInfoResponse, error) {
	α := struct {
		M OperationGetPaxAdditionalInfoSoapIn `xml:"tns:GetPaxAdditionalInfo"`
	}{
		OperationGetPaxAdditionalInfoSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationGetPaxAdditionalInfoSoapOut `xml:"GetPaxAdditionalInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/GetPaxAdditionalInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// GetStores was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) GetStores(input_arr *GetStoresRequest) (*GetStoresResponse, error) {
	α := struct {
		M OperationGetStoresSoapIn `xml:"tns:GetStores"`
	}{
		OperationGetStoresSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationGetStoresSoapOut `xml:"GetStoresResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/GetStores", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// ReserveList was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) ReserveList(input_arr *ReserveListRequest) (*ReserveListResponse, error) {
	α := struct {
		M OperationReserveListSoapIn `xml:"tns:ReserveList"`
	}{
		OperationReserveListSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationReserveListSoapOut `xml:"ReserveListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/ReserveList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// RetrieveReserve was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) RetrieveReserve(input_arr *RetrieveReserveRequest) (*RetrieveReserveResponse, error) {
	α := struct {
		M OperationRetrieveReserveSoapIn `xml:"tns:RetrieveReserve"`
	}{
		OperationRetrieveReserveSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationRetrieveReserveSoapOut `xml:"RetrieveReserveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/RetrieveReserve", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// SearchProductByCalendar was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) SearchProductByCalendar(input_arr *SearchProductByCalendarRequest) (*SearchProductByCalendarResponse, error) {
	α := struct {
		M OperationSearchProductByCalendarSoapIn `xml:"tns:SearchProductByCalendar"`
	}{
		OperationSearchProductByCalendarSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationSearchProductByCalendarSoapOut `xml:"SearchProductByCalendarResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/SearchProductByCalendar", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// SearchProductMasterData was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) SearchProductMasterData(input_arr *SearchProductMasterDataRequest) (*SearchProductMasterDataResponse, error) {
	α := struct {
		M OperationSearchProductMasterDataSoapIn `xml:"tns:SearchProductMasterData"`
	}{
		OperationSearchProductMasterDataSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationSearchProductMasterDataSoapOut `xml:"SearchProductMasterDataResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/SearchProductMasterData", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// SearchProductTypes was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) SearchProductTypes(input_arr *SearchProductTypesRequest) (*SearchProductTypesResponse, error) {
	α := struct {
		M OperationSearchProductTypesSoapIn `xml:"tns:SearchProductTypes"`
	}{
		OperationSearchProductTypesSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationSearchProductTypesSoapOut `xml:"SearchProductTypesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/SearchProductTypes", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// SearchProducts was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) SearchProducts(input_arr *SearchProductRequest) (*SearchProductResponse, error) {
	α := struct {
		M OperationSearchProductsSoapIn `xml:"tns:SearchProducts"`
	}{
		OperationSearchProductsSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationSearchProductsSoapOut `xml:"SearchProductsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/SearchProducts", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// SetBookingOffline was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) SetBookingOffline(input_arr *SetBookingOfflineRequest) (*SetBookingOfflineResponse, error) {
	α := struct {
		M OperationSetBookingOfflineSoapIn `xml:"tns:SetBookingOffline"`
	}{
		OperationSetBookingOfflineSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationSetBookingOfflineSoapOut `xml:"SetBookingOfflineResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/SetBookingOffline", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// SetPaxAdditionalInfo was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) SetPaxAdditionalInfo(input_arr *SetPaxAdditionalInfoRequest) (*SetPaxAdditionalInfoResponse, error) {
	α := struct {
		M OperationSetPaxAdditionalInfoSoapIn `xml:"tns:SetPaxAdditionalInfo"`
	}{
		OperationSetPaxAdditionalInfoSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationSetPaxAdditionalInfoSoapOut `xml:"SetPaxAdditionalInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/SetPaxAdditionalInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// SetReserveHotels was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) SetReserveHotels(input_arr *SetReserveHotelsRequest) (*SetReserveHotelsResponse, error) {
	α := struct {
		M OperationSetReserveHotelsSoapIn `xml:"tns:SetReserveHotels"`
	}{
		OperationSetReserveHotelsSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationSetReserveHotelsSoapOut `xml:"SetReserveHotelsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/SetReserveHotels", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// SimulCancelReserve was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) SimulCancelReserve(input_arr *SimulCancelReserveRequest) (*SimulCancelReserveResponse, error) {
	α := struct {
		M OperationSimulCancelReserveSoapIn `xml:"tns:SimulCancelReserve"`
	}{
		OperationSimulCancelReserveSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationSimulCancelReserveSoapOut `xml:"SimulCancelReserveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/SimulCancelReserve", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticCreateReserve was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticCreateReserve(input_arr *StaticCreateReserveRequest) (*StaticCreateReserveResponse, error) {
	α := struct {
		M OperationStaticCreateReserveSoapIn `xml:"tns:StaticCreateReserve"`
	}{
		OperationStaticCreateReserveSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticCreateReserveSoapOut `xml:"StaticCreateReserveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticCreateReserve", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticGetPaxsStruct was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticGetPaxsStruct(input_arr *StaticGetPaxsStructRequest) (*StaticGetPaxsStructResponse, error) {
	α := struct {
		M OperationStaticGetPaxsStructSoapIn `xml:"tns:StaticGetPaxsStruct"`
	}{
		OperationStaticGetPaxsStructSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticGetPaxsStructSoapOut `xml:"StaticGetPaxsStructResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticGetPaxsStruct", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticGetProductParameters was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticGetProductParameters(input_arr *StaticGetProductParametersRequest) (*StaticGetProductParametersResponse, error) {
	α := struct {
		M OperationStaticGetProductParametersSoapIn `xml:"tns:StaticGetProductParameters"`
	}{
		OperationStaticGetProductParametersSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticGetProductParametersSoapOut `xml:"StaticGetProductParametersResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticGetProductParameters", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticGetReserveFlights was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticGetReserveFlights(input_arr *StaticGetReserveFlightsRequest) (*StaticGetReserveFlightsResponse, error) {
	α := struct {
		M OperationStaticGetReserveFlightsSoapIn `xml:"tns:StaticGetReserveFlights"`
	}{
		OperationStaticGetReserveFlightsSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticGetReserveFlightsSoapOut `xml:"StaticGetReserveFlightsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticGetReserveFlights", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticGetReserveHotels was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticGetReserveHotels(input_arr *StaticGetReserveHotelsRequest) (*StaticGetReserveHotelsResponse, error) {
	α := struct {
		M OperationStaticGetReserveHotelsSoapIn `xml:"tns:StaticGetReserveHotels"`
	}{
		OperationStaticGetReserveHotelsSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticGetReserveHotelsSoapOut `xml:"StaticGetReserveHotelsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticGetReserveHotels", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticGetReserveRooming was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticGetReserveRooming(input_arr *StaticGetReserveRoomingRequest) (*StaticGetReserveRoomingResponse, error) {
	α := struct {
		M OperationStaticGetReserveRoomingSoapIn `xml:"tns:StaticGetReserveRooming"`
	}{
		OperationStaticGetReserveRoomingSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticGetReserveRoomingSoapOut `xml:"StaticGetReserveRoomingResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticGetReserveRooming", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticGetRoomsStruct was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticGetRoomsStruct(input_arr *StaticGetRoomsStructRequest) (*StaticGetRoomsStructResponse, error) {
	α := struct {
		M OperationStaticGetRoomsStructSoapIn `xml:"tns:StaticGetRoomsStruct"`
	}{
		OperationStaticGetRoomsStructSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticGetRoomsStructSoapOut `xml:"StaticGetRoomsStructResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticGetRoomsStruct", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticGetSimulation was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticGetSimulation(input_arr *StaticGetSimulationRequest) (*StaticGetSimulationResponse, error) {
	α := struct {
		M OperationStaticGetSimulationSoapIn `xml:"tns:StaticGetSimulation"`
	}{
		OperationStaticGetSimulationSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticGetSimulationSoapOut `xml:"StaticGetSimulationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticGetSimulation", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticSetReserveExtraNights was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticSetReserveExtraNights(input_arr *StaticSetReserveExtraNightsRequest) (*StaticSetReserveExtraNightsResponse, error) {
	α := struct {
		M OperationStaticSetReserveExtraNightsSoapIn `xml:"tns:StaticSetReserveExtraNights"`
	}{
		OperationStaticSetReserveExtraNightsSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticSetReserveExtraNightsSoapOut `xml:"StaticSetReserveExtraNightsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticSetReserveExtraNights", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticSetReserveFlights was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticSetReserveFlights(input_arr *StaticSetReserveFlightsRequest) (*StaticSetReserveFlightsResponse, error) {
	α := struct {
		M OperationStaticSetReserveFlightsSoapIn `xml:"tns:StaticSetReserveFlights"`
	}{
		OperationStaticSetReserveFlightsSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticSetReserveFlightsSoapOut `xml:"StaticSetReserveFlightsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticSetReserveFlights", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// StaticSetReserveRooming was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) StaticSetReserveRooming(input_arr *StaticSetReserveRoomingRequest) (*StaticSetReserveRoomingResponse, error) {
	α := struct {
		M OperationStaticSetReserveRoomingSoapIn `xml:"tns:StaticSetReserveRooming"`
	}{
		OperationStaticSetReserveRoomingSoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationStaticSetReserveRoomingSoapOut `xml:"StaticSetReserveRoomingResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/StaticSetReserveRooming", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// ValidateVISA was auto-generated from WSDL.
func (p *wbs_pkt_methodsSoap) ValidateVISA(input_arr *ValidateVISARequest) (*ValidateVISAResponse, error) {
	α := struct {
		M OperationValidateVISASoapIn `xml:"tns:ValidateVISA"`
	}{
		OperationValidateVISASoapIn{
			input_arr,
		},
	}

	γ := struct {
		M OperationValidateVISASoapOut `xml:"ValidateVISAResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://www3.optitravel.net/optitravel/export/wbs/pkt/ValidateVISA", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}
