package api

import (
	"encoding/json"
	"evendo-viator/conf"
	"evendo-viator/pkg/utils"
	"fmt"
	jfCommon "gitlab.com/jfcore/common/common"
)

func (h *ViatorAPIHandlers) GetSupplier(productCode []string) (*map[string]interface{}, error) {
	url := conf.GetConfig().ViatorEndpoint + utils.SupplierDetail
	requestBody := map[string][]string{
		"productCodes": productCode,
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "POST", utils.ViatorAPICommonHeader, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	var supplierResponse map[string]interface{}
	err = json.Unmarshal([]byte(rawResponse), &supplierResponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &supplierResponse, nil
}
