package ginext

// BodyMeta represents a body meta data like pagination or extra response information
// it should always be rendered as a map of key: value
type BodyMeta map[string]interface{}

// GeneralBody defines a general response body
type GeneralBody struct {
	Code    int         `json:"code,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    BodyMeta    `json:"meta,omitempty"`
}

func NewBodyPaginated(code int, message string, data interface{}, pager *Pager) *GeneralBody {
	return &GeneralBody{
		Code:    code,
		Message: message,
		Data:    data,
		Meta: BodyMeta{
			"page":        pager.GetPage(),
			"total_pages": pager.GetTotalPages(),
			"page_size":   pager.GetPageSize(),
			"total":       pager.TotalRows,
			"pages":       pager.GetTotalPages(),
		},
	}
}
