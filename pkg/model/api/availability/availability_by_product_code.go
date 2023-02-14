package model

type ProductAvailabilityResponse struct {
	ProductCode   string               `json:"productCode"`
	BookableItems []BookableItemDetail `json:"bookableItems"`
	Currency      string               `json:"currency"`
}

type BookableItemDetail struct {
	ProductOptionCode string         `json:"productOptionCode"`
	Seasons           []SeasonDetail `json:"seasons"`
}

type SeasonDetail struct {
	StartDate      string                `json:"startDate"`
	EndDate        string                `json:"endDate"`
	PricingRecords []PricingRecordDetail `json:"pricingRecords"`
	OperatingHours []OperatingHourDetail `json:"operatingHours"`
}

type OperatingHourDetail struct {
	DayOfWeek      string                    `json:"dayOfWeek"`
	OperatingHours []OperatingHourDetailInfo `json:"operatingHours"`
}

type OperatingHourDetailInfo struct {
	OpensAt  string `json:"opensAt"`
	ClosesAt string `json:"closesAt"`
}

type PricingRecordDetail struct {
	DaysOfWeek       []string                `json:"daysOfWeek"`
	PricingDetails   []PricingDetailInfo     `json:"pricingDetails"`
	TimedEntries     []TimedEntriesDetail    `json:"timedEntries"`
	UnavailableDates []UnavailableDateDetail `json:"unavailableDates"`
}

type TimedEntriesDetail struct {
	StartTime string `json:"startTime"`
}

type UnavailableDateDetail struct {
	Date   string `json:"date"`
	Reason string `json:"reason"`
}

type PricingDetailInfo struct {
	PricingPackageType string      `json:"pricingPackageType"`
	MinTravelers       int         `json:"minTravelers"`
	MaxTravelers       int         `json:"maxTravelers"`
	AgeBand            string      `json:"ageBand"`
	Price              PriceDetail `json:"price"`
}

type PriceDetail struct {
	Original PriceDetailOriginal `json:"original"`
}

type PriceDetailOriginal struct {
	RecommendedRetailPrice float64 `json:"recommendedRetailPrice"`
	PartnerNetPrice        float64 `json:"partnerNetPrice"`
	BookingFee             float64 `json:"bookingFee"`
	PartnerTotalPrice      float64 `json:"partnerTotalPrice"`
}
