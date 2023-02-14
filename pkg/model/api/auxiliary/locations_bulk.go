package model

type LocationBulkResponse struct {
	Locations []LocationDetail `json:"locations"`
}

type LocationDetail struct {
	Provider          string        `json:"provider"`
	Reference         string        `json:"reference"`
	ProviderReference string        `json:"providerReference"`
	Name              string        `json:"name"`
	Address           AddressDetail `json:"address"`
	Center            CenterDetail  `json:"center"`
}

type CenterDetail struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type AddressDetail struct {
	Street             string `json:"street"`
	AdministrativeArea string `json:"administrativeArea"`
	Country            string `json:"country"`
	CountryCode        string `json:"countryCode"`
	Postcode           string `json:"postcode"`
}
