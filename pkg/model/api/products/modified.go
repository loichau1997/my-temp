package model

type ModifiedResponse struct {
	Products []ProductModifiedDetail `json:"products"`
}

type ProductModifiedDetail struct {
	Status        string `json:"status"`
	ProductCode   string `json:"productCode"`
	Language      string `json:"language"`
	CreatedAt     string `json:"createdAt"`
	lastUpdatedAt string `json:"lastUpdatedAt"`
}
