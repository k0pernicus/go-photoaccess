package types

// ServiceResponse structure and allows to encode the service response for any query
type ServiceResponse struct {
	StatusCode int         `json:"status_code"`
	Response   interface{} `json:"response"`
}

type ErrorResponse struct {
	Message   Message `json:"message,omitempty"`
	ExtraInfo string  `json:"extra_info,omitempty"`
}

// GetResponse is a specific structure that handles the response for a "get" handler
type GetResponse struct {
	Data interface{} `json:"data"`
}

// PostResponse is a specific structure that handles the response for a "post" handler
type PostResponse struct {
	Data    interface{} `json:"data"`
	Message Message     `json:"message,omitempty"`
}
