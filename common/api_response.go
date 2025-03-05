package common

type successResponse struct {
	Data    interface{} `json:"data"`
	Paging  interface{} `json:"paging,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
}

func NewSuccessResponse(data interface{}, paging interface{}, filters interface{}) *successResponse {

	return &successResponse{Data: data, Paging: paging, Filters: filters}
}
func SimpleSuccessResponse(data interface{}) *successResponse {
	return &successResponse{Data: data}
}
