package api

import (
	"encoding/json"
	"evendo-viator/conf"
	model "evendo-viator/pkg/model/api/auxiliary"
	"evendo-viator/pkg/utils"
	"fmt"
	jfCommon "gitlab.com/jfcore/common/common"
)

func (h *ViatorAPIHandlers) GetLocationBulk(listLocation []string) (response *model.LocationBulkResponse, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.AuxiliaryLocationBulk
	bodyInput := map[string][]string{
		"locations": listLocation,
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "POST", utils.ViatorAPICommonHeader, nil, bodyInput)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	bulkLocation := model.LocationBulkResponse{}
	err = json.Unmarshal([]byte(rawResponse), &bulkLocation)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &bulkLocation, nil
}
