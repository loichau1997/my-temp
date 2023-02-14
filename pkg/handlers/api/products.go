package api

import (
	"encoding/json"
	"evendo-viator/conf"
	"evendo-viator/pkg/model/api/products"
	"evendo-viator/pkg/utils"
	"fmt"
	jfCommon "gitlab.com/jfcore/common/common"
	"gitlab.com/jfcore/common/ginext"
)

type ViatorAPIHandlers struct {
}

func NewViatorAPIHandlers() *ViatorAPIHandlers {
	return &ViatorAPIHandlers{}
}

func (h *ViatorAPIHandlers) TestAPI(r *ginext.Request) *ginext.Response {
	return ginext.NewResponseData(jfCommon.CODE_SUCCESS, "Success")
}

func (h *ViatorAPIHandlers) RequestProductSearch(r *ginext.Request) *ginext.Response {
	req := model.SearchProductRequest{}
	r.MustBind(&req)
	res, err := h.ProductSearch(req)
	if err != nil {
		return ginext.NewErrorResponse(jfCommon.CODE_SERVER_ERROR, fmt.Errorf("While calling API got error : %v", err).Error())
	}
	return ginext.NewResponseData(jfCommon.CODE_SUCCESS, res)
}

func (h *ViatorAPIHandlers) RequestProductBulk(r *ginext.Request) *ginext.Response {
	req := model.BulkProductRequest{}
	r.MustBind(&req)
	res, err := h.ProductBulk(req)
	if err != nil {
		return ginext.NewErrorResponse(jfCommon.CODE_SERVER_ERROR, fmt.Errorf("While calling API got error : %v", err).Error())
	}
	return ginext.NewResponseData(jfCommon.CODE_SUCCESS, res)
}

func (h *ViatorAPIHandlers) RequestProductByCode(r *ginext.Request) *ginext.Response {
	productCode := r.GinCtx.Param("id")
	res, err := h.GetProductByCode(productCode)
	if err != nil {
		return ginext.NewErrorResponse(jfCommon.CODE_SERVER_ERROR, fmt.Errorf("While calling API got error : %v", err).Error())
	}
	return ginext.NewResponseData(jfCommon.CODE_SUCCESS, res)
}

func (h *ViatorAPIHandlers) ProductSearch(requestObj model.SearchProductRequest) (response *model.SearchProductResponse, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.ViatorProductSearchPath

	requestBody, err := utils.ModelToObj(requestObj)
	if err != nil {
		return nil, fmt.Errorf("While calling API gor error : %s", err.Error())
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "POST", utils.ViatorAPICommonHeader, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	searchProductResponse := model.SearchProductResponse{}
	err = json.Unmarshal([]byte(rawResponse), &searchProductResponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &searchProductResponse, nil
}

func (h *ViatorAPIHandlers) ProductBulk(requestObj model.BulkProductRequest) (response *[]model.BulkProductDetail, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.ViatorProductBulkPath

	requestBody, err := utils.ModelToObj(requestObj)
	if err != nil {
		return nil, fmt.Errorf("While parsing input, got error : %v", err)
	}
	rawResponse, _, err := jfCommon.SendRestAPI(url, "POST", utils.ViatorAPICommonHeader, nil, requestBody)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	searchProductResponse := []model.BulkProductDetail{}
	err = json.Unmarshal([]byte(rawResponse), &searchProductResponse)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &searchProductResponse, nil
}

func (h *ViatorAPIHandlers) GetProductByCode(requestProductCode string) (response *model.ProductDetailByCode, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.ViatorProductGetByCode + requestProductCode
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	productDetail := model.ProductDetailByCode{}
	err = json.Unmarshal([]byte(rawResponse), &productDetail)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &productDetail, nil
}

func (h *ViatorAPIHandlers) GetProductByCodeAndLanguage(requestProductCode string, language string) (response *model.ProductDetailByCode, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.ViatorProductGetByCode + requestProductCode
	header := utils.ViatorAPISpecificHeader
	header["Accept-Language"] = language
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", header, nil, nil)

	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	productDetail := model.ProductDetailByCode{}
	err = json.Unmarshal([]byte(rawResponse), &productDetail)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &productDetail, nil
}
func (h *ViatorAPIHandlers) GetListTags() (response *model.TagsResponse, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.ViatorProductTags
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	listTags := model.TagsResponse{}
	err = json.Unmarshal([]byte(rawResponse), &listTags)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &listTags, nil
}

func (h *ViatorAPIHandlers) GetProductModifiedSince() (response *model.ModifiedResponse, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.ViatorProductModifiedSince
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	modifiedSince := model.ModifiedResponse{}
	err = json.Unmarshal([]byte(rawResponse), &modifiedSince)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &modifiedSince, nil
}

func (h *ViatorAPIHandlers) BookingQuestion() (response *model.ModifiedResponse, err error) {
	url := conf.GetConfig().ViatorEndpoint + utils.ViatorProductBookingQuestion
	rawResponse, _, err := jfCommon.SendRestAPI(url, "GET", utils.ViatorAPICommonHeader, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("While calling API, got error : %v", err)
	}
	bookingQuestion := model.ModifiedResponse{}
	err = json.Unmarshal([]byte(rawResponse), &bookingQuestion)
	if err != nil {
		return nil, fmt.Errorf("While parsing API response, got error : %v", err)
	}
	return &bookingQuestion, nil
}
