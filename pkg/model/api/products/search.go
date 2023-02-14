package model

import "github.com/jinzhu/gorm/dialects/postgres"

type SearchProductRequest struct {
	Filtering  FilteringDetail  `json:"filtering,omitempty"`
	Sorting    SortingDetail    `json:"sorting,omitempty"`
	Pagination PaginationDetail `json:"pagination,omitempty"`
	Currency   string           `json:"currency,omitempty"`
}

type FilteringDetail struct {
	Destination  string   `json:"destination,omitempty"`
	Tags         []int    `json:"tags,omitempty"`
	Flags        []string `json:"flags,omitempty"`
	HighestPrice *int     `json:"highestPrice,omitempty"`
	StartDate    *string  `json:"startDate,omitempty"`
	EndDate      *string  `json:"endDate,omitempty"`
}

type SortingDetail struct {
	Sort  *string `json:"sort,omitempty"`
	Order *string `json:"order,omitempty"`
}

type PaginationDetail struct {
	Start int `json:"start,omitempty"`
	Count int `json:"count,omitempty"`
}

type SearchProductResponse struct {
	Products   []ProductsDetail `json:"products"`
	TotalCount int              `json:"totalCount"`
}

type ProductsDetail struct {
	ProductCode     string                `json:"productCode"`
	Title           string                `json:"title"`
	Description     string                `json:"description"`
	Images          []ImageDetail         `json:"images"`
	Reviews         ReviewDetail          `json:"reviews"`
	Pricing         PricingDetail         `json:"pricing"`
	Destinations    []DestinationDetail   `json:"destinations"`
	Tags            []int                 `json:"tags"`
	Flags           []string              `json:"flags"`
	TranslationInfo TranslationInfoDetail `json:"translationInfo"`
}

type TranslationInfoDetail struct {
	ContainsMachineTranslatedText bool `json:"containsMachineTranslatedText"`
}

type ImageDetail struct {
	ImageSource string         `json:"imageSource"`
	Caption     string         `json:"caption"`
	IsCover     bool           `json:"isCover"`
	Variants    postgres.Jsonb `json:"variants"`
}

type PricingDetail struct {
	Summary             PricingSummaryDetail `json:"summary"`
	PartnerNetFromPrice float64              `json:"partnerNetFromPrice"`
	Currency            string               `json:"currency"`
}

type DestinationDetail struct {
	Ref     string `json:"ref"`
	Primary bool   `json:"primary"`
}

type PricingSummaryDetail struct {
	FromPrice               float64 `json:"fromPrice"`
	FromPriceBeforeDiscount float64 `json:"fromPriceBeforeDiscount"`
}

type VariantDetail struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Url    string `json:"url"`
}

type ReviewDetail struct {
	Sources               []ReivewSourceDetail `json:"sources"`
	ReviewCountTotals     []ReviewCountDetail  `json:"reviewCountTotals"`
	TotalViews            int                  `json:"totalViews"`
	CombinedAverageRating float64              `json:"combinedAverageRating"`
}

type ReivewSourceDetail struct {
	Provider      string  `json:"provider"`
	TotalCount    int     `json:"totalCount"`
	AverageRating float64 `json:"averageRating"`
}

type ReviewCountDetail struct {
	Rating float64 `json:"rating"`
	Count  int     `json:"count"`
}
