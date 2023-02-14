package api

import (
	"encoding/json"
	"evendo-viator/conf"
	model "evendo-viator/pkg/model/api/availability"
	"evendo-viator/pkg/utils"
	"fmt"
	jfCommon "gitlab.com/jfcore/common/common"
)

func (h *ViatorAPIHandlers) AvailabilitySchedulesByProductCode(productCode string) (response *model.ProductAvailabilityResponse, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.AvailabilityScheduleByID + productCode
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	availabilitySchedule := model.ProductAvailabilityResponse{}
	err = json.Unmarshal([]byte(rawResponse), &availabilitySchedule)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &availabilitySchedule, nil
}

func (h *ViatorAPIHandlers) AvailabilitySchedulesModifiedBy(params interface{}) (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.AvailabilityScheduleModifiedSince

	requestParams, err := utils.ModelToParam(params)
	if err != nil {
		return nil, fmt.Errorf("While parsing input, got error : %v", err)
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, requestParams, nil)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var availabilityModifiedSince interface{}
	err = json.Unmarshal([]byte(rawResponse), &availabilityModifiedSince)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &availabilityModifiedSince, nil
}

func (h *ViatorAPIHandlers) AvailabilitySchedulesCheck(requestObj interface{}) (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.AvailabilityCheck

	requestBody, err := utils.ModelToObj(requestObj)
	if err != nil {
		return nil, fmt.Errorf("While parsing input, got error : %v", err)
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var availabilityScheduleCheck interface{}
	err = json.Unmarshal([]byte(rawResponse), &availabilityScheduleCheck)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &availabilityScheduleCheck, nil
}

func (h *ViatorAPIHandlers) AvailabilitySchedulesBulk(requestObj interface{}) (response interface{}, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.AvailabilityScheduleBulk

	requestBody, err := utils.ModelToObj(requestObj)
	if err != nil {
		return nil, fmt.Errorf("While parsing input, got error : %v", err)
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var availabilityScheduleBulk interface{}
	err = json.Unmarshal([]byte(rawResponse), &availabilityScheduleBulk)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &availabilityScheduleBulk, nil
}
