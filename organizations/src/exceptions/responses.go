package exceptions

// ErrorResponse is the schema for data that is returned when an error occurs.
type ErrorResponse struct {
	Object     string `json:"object"`
	Name       string `json:"name"`
	StatusCode int    `json:"statusCode"`
	Details    string `json:"details"`
}

// MultiErrorResponse is the schema for data that is returned when multiple errors occur.
type MultiErrorResponse struct {
	Object     string `json:"object"`
	Name       string `json:"name"`
	StatusCode int `json:"statusCode"`
	Details    []ErrorDetail `json:"details"`
}

// ErrorDetail is the schema for data that is returned to explain specific error details.
type ErrorDetail struct {
	Object  string `json:"object"`
	Name    string `json:"name"`
	Details string `json:"details"`
}

// NewErrorResponse creates an ErrorResponse object
func NewErrorResponse(detail ErrorDetail, statusCode int) ErrorResponse {
	res := ErrorResponse{}
	res.Object = "error"
	res.Name = detail.Name
	res.StatusCode = statusCode
	res.Details = detail.Details
	return res
}

// NewMultiErrorResponse creates a MultiErrorResponse object.
func NewMultiErrorResponse(name string, statusCode int, details []ErrorDetail) MultiErrorResponse {
	res := MultiErrorResponse{}
	res.Object = "error"
	res.Name = name
	res.StatusCode = statusCode
	res.Details = details
	return res
}

// NewErrorDetail creates a new ErrorDetail object.
func NewErrorDetail(name, details string) ErrorDetail {
	return ErrorDetail{
		Object: "error-detail",
		Name: name,
		Details: details,
	}
}