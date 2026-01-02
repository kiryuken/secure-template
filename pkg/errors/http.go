package errors

// HTTP error responses

type ErrorResponse struct {
	Code    string
	Message string
	Details interface{}
}

// ToHTTPResponse converts AppError to HTTP response
func ToHTTPResponse() ErrorResponse {
	// TODO: Implement
	return ErrorResponse{}
}

// GetHTTPStatusCode returns HTTP status code for error
func GetHTTPStatusCode() int {
	// TODO: Implement
	return 500
}
