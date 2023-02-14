package model

type BookingQuestionsResponse struct {
	LegacyBookingQuestionID int      `json:"legacyBookingQuestionId"`
	Id                      string   `json:"id"`
	Type                    string   `json:"type"`
	Group                   string   `json:"group"`
	Label                   string   `json:"label"`
	Required                string   `json:"required"`
	MaxLength               int      `json:"maxLength"`
	Units                   []string `json:"units"`
	Hint                    string   `json:"hint"`
}
