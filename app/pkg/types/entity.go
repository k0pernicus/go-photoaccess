package types

import "time"

// required is a special structure to be composed with all the other ones, which
// contains specific and required fields to register (like the ID, the creation date, etc...)
type RequiredFields struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Photo is the raw structure of a photo, which can contains annotations.
// A Photo entity can contains annotations, or an annotation id, but it can't contains both.
// The content of a photo must be base64.
type Photo struct {
	Content     string       `json:"content"` // Base64
	Annotations []Annotation `json:"annotations,omitempty"`
	RequiredFields
}

// PhotoCreationRequest is a simple structure that contains all the information
// to create a photo entity.
// The content of the data is returned as a raw base64 string.
type PhotoCreationRequest struct {
	Data string `json:"data"`
}

// PhotoCreationResponse is a simple structure that contains all the information
// when a photo entity has been created
type PhotoCreationResponse struct {
	ID string `json:"id"`
}

// PhotoGetResponse is a structure that contains all the information
type PhotoGetResponse struct {
	Photo       Photo                   `json:"photo"`
	Annotations []AnnotationGetResponse `json:"annotations,omitempty"`
}

// Coordinates allows to fusion all coordinates informations about an
// Annotation in a picture, in pixels.
// X is the pixel at left, and X2 at right (the last one of the bounding box).
// Y is the most upper pixel, and Y2 the most lower bottom one.
type Coordinates struct {
	X  int `json:"x"`
	X2 int `json:"x2"`
	Y  int `json:"y"`
	Y2 int `json:"y2"`
}

// An annotation must be linked to a photo, and contains text
type Annotation struct {
	PhotoID     int         `json:"photo_id"` // Linked to an existing Photo entity
	Content     string      `json:"content"`
	Coordinates Coordinates `json:"coordinates"`
	RequiredFields
}

// AnnotationCreationRequest is a simple structure that contains an annotation
// for an existing photo
type AnnotationCreationRequest struct {
	Text        string      `json:"text"`
	Coordinates Coordinates `json:"coordinates"`
}

// AnnotationCreationResponse is a simple structure that contains all the information
// when an Annotation entity has been created
type AnnotationCreationResponse struct {
	ID string `json:"id"`
}

// AnnotationGetResponse is a simple structure that contains all the information
// about an Annotation entity
type AnnotationGetResponse struct {
	Annotation Annotation `json:"annotation"`
}
