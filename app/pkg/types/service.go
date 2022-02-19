package types

// ServiceResponse structure and allows to encode the service response for any query
type ServiceResponse struct {
	StatusCode int         `json:"status_code"`
	Response   interface{} `json:"response"`
}

type ErrorResponse struct {
	Message Message `json:"message,omitempty"`
}

// GetResponse is a specific structure that handles the response for a "get" handler
type GetResponse struct {
	Data interface{} `json:"data"`
}

// DeleteResponse is a specific structure that handles the response for a "delete" handler
// The message is only used when something goes wrong - otherwise, the status code is "truth"
type DeleteResponse struct {
	Message Message `json:"message,omitempty"`
}

// PostResponse is a specific structure that handles the response for a "post" handler
type PostResponse struct {
	Data    interface{} `json:"data"`
	Message Message     `json:"message,omitempty"`
}

// ExistsResponse is a specific structure that handles the response for an "exists" handler
type ExistsResponse struct {
	Exists  bool    `json:"exists"`
	Message Message `json:"message,omitempty"`
}
