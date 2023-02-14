package api

import (
	"encoding/json"
	"evendo-viator/conf"
	"evendo-viator/pkg/utils"
	"fmt"
	jfCommon "gitlab.com/jfcore/common/common"
)

func (h *ViatorAPIHandlers) BookingHold(requestObj interface{}) (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.BookingHold

	requestBody, err := utils.ModelToObj(requestObj)
	if err != nil {
		return nil, fmt.Errorf("While parsing input, got error : %v", err)
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "POST", utils.ViatorAPICommonHeader, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var bookingHoldResponse interface{}
	err = json.Unmarshal([]byte(rawResponse), &bookingHoldResponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &bookingHoldResponse, nil
}

func (h *ViatorAPIHandlers) BookingBook(requestObj interface{}) (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.BookingStatus

	requestBody, err := utils.ModelToObj(requestObj)
	if err != nil {
		return nil, fmt.Errorf("While parsing input, got error : %v", err)
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "POST", utils.ViatorAPICommonHeader, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var bookingStatus interface{}
	err = json.Unmarshal([]byte(rawResponse), &bookingStatus)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &bookingStatus, nil
}

func (h *ViatorAPIHandlers) BookingStatus(requestObj interface{}) (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.AvailabilityCheck

	requestBody, err := utils.ModelToObj(requestObj)
	if err != nil {
		return nil, fmt.Errorf("While parsing input, got error : %v", err)
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "POST", utils.ViatorAPICommonHeader, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var bookingStatusResponse interface{}
	err = json.Unmarshal([]byte(rawResponse), &bookingStatusResponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &bookingStatusResponse, nil
}

func (h *ViatorAPIHandlers) BookingCancel(bookingID string) (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.BookingCancel(bookingID)

	rawResponse, _, err := jfCommon.SendRestAPI(url, "POST", utils.ViatorAPICommonHeader, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var bookingCancelReponse interface{}
	err = json.Unmarshal([]byte(rawResponse), &bookingCancelReponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &bookingCancelReponse, nil
}

func (h *ViatorAPIHandlers) BookingCancelQuote(bookingID string) (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.BookingCancelQuote(bookingID)

	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var bookingCancelQuoteResponse interface{}
	err = json.Unmarshal([]byte(rawResponse), &bookingCancelQuoteResponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &bookingCancelQuoteResponse, nil
}

func (h *ViatorAPIHandlers) BookingModifiedSince() (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.BookingModifiedSinces

	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var bookingModifiedSinceResponse interface{}
	err = json.Unmarshal([]byte(rawResponse), &bookingModifiedSinceResponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &bookingModifiedSinceResponse, nil
}

func (h *ViatorAPIHandlers) BookingModifiedSinceAcknowledge(requestObj interface{}) (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.BookingModifiedSincesAcknowledge

	requestBody, err := utils.ModelToObj(requestObj)
	if err != nil {
		return nil, fmt.Errorf("While parsing input, got error : %v", err)
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var bookingModifiedSinceAcknowledgeResponse interface{}
	err = json.Unmarshal([]byte(rawResponse), &bookingModifiedSinceAcknowledgeResponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &bookingModifiedSinceAcknowledgeResponse, nil
}

func (h *ViatorAPIHandlers) BookingCancelReason(cancelType string) (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.AvailabilityScheduleModifiedSince

	requestParams := map[string]string{
		"type": cancelType,
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, requestParams, nil)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var bookingCancelReasonResponse interface{}
	err = json.Unmarshal([]byte(rawResponse), &bookingCancelReasonResponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &bookingCancelReasonResponse, nil
}
