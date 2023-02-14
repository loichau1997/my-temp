package destination

//type ProductAvailabilityResponse struct {
//	ProductCode   string               `json:"productCode"`
//	BookableItems []BookableItemDetail `json:"bookableItems"`
//	Currency      string               `json:"currency"`
//}

type APIDestinationResponse struct {
	ErrorReference interface{}            `json:"error_reference"`
	Data           []APIDestinationDetail `json:"data"`
}

type APIDestinationDetail struct {
	SortOrder           int     `json:"sortOrder"`
	Selectable          bool    `json:"selectable"`
	DestinationURLName  string  `json:"destinationUrlName"`
	DefaultCurrencyCode string  `json:"defaultCurrencyCode"`
	LookupID            string  `json:"lookupId"`
	ParentID            int     `json:"parentId"`
	Timezone            string  `json:"timeZone"`
	IataCode            string  `json:"iataCode"`
	DestinationName     string  `json:"destinationName"`
	DestinationID       int     `json:"destinationID"`
	DestinationType     string  `json:"destinationType"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
}
