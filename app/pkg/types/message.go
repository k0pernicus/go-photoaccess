package types

// Message is a type that contains a specific message when receiving a message from the service
type Message string

const (
	// OK seems... fine
	OK Message = "OK"
	// InternalError signifies that something gone wrong
	InternalError Message = "Internal Error"
	// CannotDecodeMessage if the body message is not correctly formated
	CannotDecodeMessage Message = "Cannot decode message"
	// EntityNotFound is sent back if the requested entity has not been found in DB operation(s)
	EntityNotFound Message = "Entity not found"
	// When information is missing in query parameter for example
	MissingInformation Message = "Missing information"
	// When user / developer send us invalid information (in payload for example)
	InvalidInformation Message = "Invalid information received"
)
