package model

type SupplierResponseDetail struct {
	Suppliers []SupplierContactInfo `json:"suppliers"`
}

type SupplierInfo struct {
	Reference   string              `json:"reference"`
	Name        string              `json:"name"`
	Type        string              `json:"type"`
	ProductCode string              `json:"productCode"`
	Contact     SupplierContactInfo `json:"contact"`
}

type SupplierContactInfo struct {
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
