package wsdl_to_dto

import (
	"tarmac/dto"
	"tarmac/wsdl"
)

// Helper function to convert a wsdl.DynProductParametersResponse struct with string pointers into a dto.DynProductParametersResponse with regular strings
func ToDynProductParametersResponseDTO(r *wsdl.DynProductParametersResponse) dto.DynProductParametersResponse {
	if r == nil {
		return dto.DynProductParametersResponse{}
	}

	return dto.DynProductParametersResponse{
		Code:            deref(r.Code),
		Name:            deref(r.Name),
		Localization:    deref(r.Localization),
		Country:         deref(r.Country),
		Zone:            deref(r.Zone),
		ExtraNightsFrom: deref(r.ExtraNightsFrom),
		ExtraNightsTo:   deref(r.ExtraNightsTo),
		MinPaxReserve:   deref(r.MinPaxReserve),
		MaxPaxReserve:   deref(r.MaxPaxReserve),
		DepartureLocals: mapDepartureLocals(r.DepartureLocals),
		DepartureDates:  mapDepartureDates(r.DepartureDates),
		RoomTypes:       mapRoomTypes(r.RoomTypes),
		BaseLocals:      mapBaseLocals(r.BaseLocals),
		ExtraLocals:     mapExtraLocals(r.ExtraLocals),
		RequestAges:     mapRequestAges(r.RequestAges),
		Errors:          mapErrors(r.Errors),
	}
}

// Helper function to map DepartureLocalDTOs
func mapDepartureLocals(locals *wsdl.DynDepartureLocalArray) []dto.DepartureLocalDTO {
	if locals == nil {
		return nil
	}

	dtoLocals := make([]dto.DepartureLocalDTO, len(locals.Items))
	for i, local := range locals.Items {
		dtoLocals[i] = dto.DepartureLocalDTO{
			ID:   deref(local.Code),
			Name: deref(local.Name),
		}
	}
	return dtoLocals
}

// Helper function to map DepartureDateDTOs
func mapDepartureDates(dates *wsdl.DynDepartureDateArray) []dto.DepartureDateDTO {
	if dates == nil {
		return nil
	}
	dtoDates := make([]dto.DepartureDateDTO, len(dates.Items))
	for i, date := range dates.Items {
		dtoDates[i] = dto.DepartureDateDTO{
			Date: deref(date.Date),
		}
	}
	return dtoDates
}

// Helper function to map RoomTypeDTOs
func mapRoomTypes(roomTypes *wsdl.DynRoomTypeArray) []dto.RoomTypeDTO {
	if roomTypes == nil {
		return nil
	}
	dtoRoomTypes := make([]dto.RoomTypeDTO, len(roomTypes.Items))
	for i, room := range roomTypes.Items {
		dtoRoomTypes[i] = dto.RoomTypeDTO{
			Code:         deref(room.Code),
			Name:         deref(room.Name),
			NumAdults:    deref(room.NumAdults),
			NumChilds:    deref(room.NumChilds),
			ChildAgeFrom: deref(room.ChildAgeFrom),
			ChildAgeTo:   deref(room.ChildAgeTo),
		}
	}
	return dtoRoomTypes
}

// Helper function to map BaseLocalDTOs
func mapBaseLocals(baseLocals *wsdl.DynBaseLocalArray) []dto.BaseLocalDTO {
	if baseLocals == nil {
		return nil
	}
	dtoBaseLocals := make([]dto.BaseLocalDTO, len(baseLocals.Items))
	for i, local := range baseLocals.Items {
		dtoBaseLocals[i] = dto.BaseLocalDTO{
			Code:            deref(local.Code),
			Name:            deref(local.Name),
			MinNights:       deref(local.MinNights),
			MaxNights:       deref(local.MaxNights),
			BeforeAfterBase: deref(local.BeforeAfterBase),
		}
	}
	return dtoBaseLocals
}

// Helper function to map ExtraLocalDTOs
func mapExtraLocals(extraLocals *wsdl.DynExtraLocalArray) []dto.ExtraLocalDTO {
	if extraLocals == nil {
		return nil
	}
	dtoExtraLocals := make([]dto.ExtraLocalDTO, len(extraLocals.Items))
	for i, local := range extraLocals.Items {
		dtoExtraLocals[i] = dto.ExtraLocalDTO{
			Code:   deref(local.Code),
			Name:   deref(local.Name),
			Locals: mapBaseLocals(local.Locals),
		}
	}
	return dtoExtraLocals
}

// Helper function to map RequestAgesDTO
func mapRequestAges(ages *wsdl.DynRequestAges) *dto.RequestAgesDTO {
	if ages == nil {
		return nil
	}
	return &dto.RequestAgesDTO{
		ReqAgeType:  deref(ages.ReqAgeType),
		MinAgeAdult: deref(ages.MinAgeAdult),
		MaxAgeAdult: deref(ages.MaxAgeAdult),
	}
}

// Helper function to map ErrorDTO
func mapErrors(errs *wsdl.ErrorStruct) *dto.ErrorDTO {
	if errs == nil {
		return nil
	}
	var dtoErrors []dto.ErrorObj
	if errs.ErrorList != nil && errs.ErrorList.Items != nil {
		dtoErrors = make([]dto.ErrorObj, len(errs.ErrorList.Items))
		for i, err := range errs.ErrorList.Items {
			dtoErrors[i] = dto.ErrorObj{
				Code: deref(err.Code),
				Desc: deref(err.Desc),
			}
		}
	}
	return &dto.ErrorDTO{
		HasErrors: deref(errs.HasErrors),
		ErrorList: dtoErrors,
	}
}
