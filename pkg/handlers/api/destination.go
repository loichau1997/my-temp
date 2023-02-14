package api

import (
	"encoding/json"
	"evendo-viator/conf"
	model "evendo-viator/pkg/model/api/destination"
	"evendo-viator/pkg/utils"
	"fmt"
	jfCommon "gitlab.com/jfcore/common/common"
)

func (h *ViatorAPIHandlers) GetDestination() (response *model.APIDestinationResponse, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.ViatorDestination
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, utils.ViatorDestination)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var getDestiantionResponse model.APIDestinationResponse
	err = json.Unmarshal([]byte(rawResponse), &getDestiantionResponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &getDestiantionResponse, nil
}
