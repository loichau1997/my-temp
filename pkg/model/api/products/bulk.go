package model

type BulkProductRequest struct {
	ProductCodes []string         `json:"productCodes,omitempty"`
	Sorting      SortingDetail    `json:"sorting,omitempty"`
	Pagination   PaginationDetail `json:"pagination,omitempty"`
}

type BulkProductResponse struct {
	BulkProductDetail []BulkProductDetail
}

type BulkProductDetail struct {
	Status        string            `json:"status"`
	ProductCode   string            `json:"productCode"`
	Language      string            `json:"language"`
	CreatedAt     string            `json:"createdAt"`
	LastUpdatedAt string            `json:"lastUpdatedAt"`
	Title         string            `json:"title"`
	TicketInfo    TicketInfoDetail  `json:"ticketInfo"`
	PricingInfo   PricingInfoDetail `json:"pricingInfo"`
	Images        []ImageDetail     `json:"images"`
}

type LogisticDetail struct {
	Start                       []LogisticStartDetail            `json:"start"`
	End                         []LogisticEndDetail              `json:"end"`
	Redemption                  RedemptionDetail                 `json:"redemption"`
	TravelerPickup              TravelerPickupDetail             `json:"travelerPickup"`
	TimeZone                    string                           `json:"timeZone"`
	Description                 string                           `json:"description"`
	Inclusions                  []InclusionDetail                `json:"inclusions"`
	Exclusions                  []ExclusionDetail                `json:"exclusions"`
	AdditionalInfo              []AdditionalInfoDetail           `json:"additionalInfo"`
	CancellationPolicy          CancellationPolicyDetail         `json:"cancellationPolicy"`
	BookingConfirmationSettings BookingConfirmationSettingDetail `json:"bookingConfirmationSettings"`
	BookingRequirements         BookingRequirementDetail         `json:"bookingRequirements"`
	BookingQuestions            []string                         `json:"bookingQuestions"`
	Tags                        []int                            `json:"tags"`
	Destinations                []DestinationDetail              `json:"destinations"`
	Itinerary                   ItineraryDetail                  `json:"itinerary"`
	ProductOptions              []ProductOptionDetail            `json:"productOptions"`
	TranslationInfo             TranslationInfoDetail            `json:"translationInfo"`
	Supplier                    SupplierDetail                   `json:"supplier"`
	Reviews                     ReviewDetail                     `json:"reviews"`
}

type SupplierDetail struct {
	Name         string `json:"name"`
	Reference    string `json:"reference"`
	Contact      string `json:"contact"`
	SupplierType string `json:"supplierType"`
}

type ProductOptionDetail struct {
	ProductOptionCode string `json:"productOptionCode"`
	Description       string `json:"description"`
	Title             string `json:"title"`
}

type ItineraryDetail struct {
	ItineraryType            string                           `json:"itineraryType"`
	SkipTheLine              bool                             `json:"skipTheLine"`
	PrivateTour              bool                             `json:"privateTour"`
	Duration                 ItineraryDurationDetail          `json:"duration"`
	ItineraryItems           []ItineraryItemDetail            `json:"itineraryItems"`
	ActivityInfo             ItineraryActivityDetail          `json:"activityInfo"`
	Days                     []ItineraryDayDetail             `json:"days"`
	PointOfInterestLocations []PointOfInterestLocationsDetail `json:"pointOfInterestLocations"`
	PointOfInterest          []LocationDetail                 `json:"pointsOfInterest"`
	Routes                   []RouteDetail                    `json:"routes"`
}

type RouteDetail struct {
	OperatingSchedule string                          `json:"operatingSchedule"`
	Duration          ItineraryDurationDetail         `json:"duration"`
	Name              string                          `json:"name"`
	Stops             []RouteStop                     `json:"stops"`
	PointOfInterest   []PointOfInterestLocationDetail `json:"pointsOfInterest"`
}

type RouteStop struct {
	StopLocation LocationDetail `json:"stopLocation"`
	Description  string         `json:"description"`
}

type PointOfInterestLocationsDetail struct {
	Location LocationDetail `json:"location"`
}

type ItineraryActivityDetail struct {
	Location    LocationDetail `json:"location"`
	Description string         `json:"description"`
}

type ItineraryItemDetail struct {
	PointOfInterestLocation PointOfInterestLocationDetail `json:"pointOfInterestLocation"`
	Duration                ItineraryDurationDetail       `json:"duration"`
	PassByWithoutStopping   bool                          `json:"passByWithoutStopping"`
	AdmissionIncluded       string                        `json:"admissionIncluded"`
	Description             string                        `json:"description"`
}

type PointOfInterestLocationDetail struct {
	Location     LocationDetail `json:"location"`
	AttractionID int            `json:"attractionId"`
}

type ItineraryDayDetail struct {
	Title     string `json:"title"`
	DayNumber int    `json:"dayNumber"`
	Items     []struct {
		PointOfInterestLocation struct {
			Location LocationDetail `json:"location"`
		} `json:"pointOfInterestLocation"`
		Duration struct {
			FixedDurationInMinutes int `json:"fixedDurationInMinutes"`
		} `json:"duration"`
		PassByWithoutStopping bool   `json:"passByWithoutStopping"`
		AdmissionIncluded     string `json:"admissionIncluded"`
		Description           string `json:"description"`
	} `json:"items"`
	Accommodations []struct {
		Description string `json:"description"`
	} `json:"accommodations"`
}

type ItineraryDurationDetail struct {
	FixedDurationInMinutes int `json:"fixedDurationInMinutes"`
}

type BookingRequirementDetail struct {
	MinTravelersPerBooking  int  `json:"minTravelersPerBooking"`
	MaxTravelersPerBooking  int  `json:"maxTravelersPerBooking"`
	RequiresAdultForBooking bool `json:"requiresAdultForBooking"`
}

type BookingConfirmationSettingDetail struct {
	BookingCutoffType      string `json:"bookingCutoffType"`
	BookingCutoffInMinutes int    `json:"bookingCutoffInMinutes"`
	ConfirmationType       string `json:"confirmationType"`
}

type CancellationPolicyDetail struct {
	Type                          string                    `json:"type"`
	Description                   string                    `json:"description"`
	CancelIfBadWeather            bool                      `json:"cancelIfBadWeather"`
	CancelIfInSufficientTravelers bool                      `json:"cancelIfInSufficientTravelers"`
	RefundEligibility             []RefundEligibilityDetail `json:"refundEligibility"`
}

type RefundEligibilityDetail struct {
	DayRangeMin          int  `json:"dayRangeMin"`
	DayRangeMax          *int `json:"dayRangeMax"`
	PercentageRefundable int  `json:"percentageRefundable"`
}

type AdditionalInfoDetail struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type ExclusionDetail struct {
	Category            string `json:"category"`
	CategoryDescription string `json:"categoryDescription"`
	Type                string `json:"type"`
	TypeDescription     string `json:"typeDescription"`
	OtherDescription    string `json:"otherDescription"`
}

type InclusionDetail struct {
	Category            string `json:"category"`
	CategoryDescription string `json:"categoryDescription"`
	Type                string `json:"type"`
	TypeDescription     string `json:"typeDescription"`
	OtherDescription    string `json:"otherDescription"`
}

type TravelerPickupDetail struct {
	PickupOptionType          string                         `json:"pickupOptionType"`
	AllowCustomTravelerPickup bool                           `json:"allowCustomTravelerPickup"`
	Locations                 []TravelerPickupLocationDetail `json:"locations"`
}

type TravelerPickupLocationDetail struct {
	Location struct {
		Ref string `json:"ref"`
	} `json:"location"`
	PickupType string `json:"pickupType"`
}

type RedemptionDetail struct {
	RedemptionType      string `json:"redemptionType"`
	SpecialInstructions string `json:"specialInstructions"`
}

type LogisticStartDetail struct {
	Location LocationDetail `json:"location"`
}

type LogisticEndDetail struct {
	Location LocationDetail `json:"location"`
}

type LocationDetail struct {
	Ref string `json:"ref"`
}

type PricingInfoDetail struct {
	Type     string          `json:"type"`
	AgeBands []AgeBandDetail `json:"ageBands"`
}

type AgeBandDetail struct {
	AgeBand                string `json:"ageBand"`
	StartAge               int    `json:"startAge"`
	EndAge                 int    `json:"endAge"`
	MinTravelersPerBooking int    `json:"minTravelersPerBooking"`
	MaxTravelersPerBooking int    `json:"maxTravelersPerBooking"`
}

type TicketInfoDetail struct {
	TicketTypes                 []string `json:"ticketTypes"`
	TicketTypeDescription       string   `json:"ticketTypeDescription"`
	TicketPerBook               string   `json:"ticketsPerBooking"`
	TicketPerBookingDescription string   `json:"ticketsPerBookingDescription"`
}
