package model

type LanguageGuideDetail struct {
	Language    string `json:"language"`
	LegacyGuide string `json:"legacyGuide"`
	Type        string `json:"type"`
}

type ProductDetailByCode struct {
	Status                      string                           `json:"status"`
	ProductCode                 string                           `json:"productCode"`
	Language                    string                           `json:"language"`
	CreatedAt                   string                           `json:"createdAt"`
	LastUpdatedAt               string                           `json:"lastUpdatedAt"`
	Title                       string                           `json:"title"`
	LanguageGuides              interface{}                      `json:"languageGuides"`
	TicketInfo                  TicketInfoDetail                 `json:"ticketInfo"`
	PricingInfo                 PricingInfoDetail                `json:"pricingInfo"`
	Images                      []ImageDetail                    `json:"images"`
	Logistics                   LogisticDetail                   `json:"logistics"`
	TimeZone                    string                           `json:"timeZone"`
	Description                 string                           `json:"description"`
	Inclusions                  []InclusionDetail                `json:"inclusions"`
	Exclusions                  []InclusionDetail                `json:"exclusions"`
	AdditionalInfo              []AdditionalInfoDetail           `json:"additionalInfo"`
	CancellationPolicy          CancellationPolicyDetail         `json:"cancellationPolicy"`
	BookingConfirmationSettings BookingConfirmationSettingDetail `json:"bookingConfirmationSettings"`
	BookingRequirements         BookingRequirementDetail         `json:"bookingRequirements"`
	Supplier                    SupplierDetail                   `json:"supplier"`
	Reviews                     ReviewDetail                     `json:"reviews"`
	Tags                        []int64                          `json:"tags"`
	ProductOptions              []ProductOptionDetail            `json:"productOptions"`
	Destinations                []DestinationDetail              `json:"destinations"`
	Itinerary                   ItineraryDetail                  `json:"itinerary"`
	//Itinerary interface{} `json:"itinerary"`
}

//type ProductOptionDetail struct {
//	ProductOptionCode string `json:"productOptionCode"`
//	Description       string `json:"description"`
//	Title             string `json:"title"`
//}
