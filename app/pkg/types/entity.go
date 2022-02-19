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
	Content      string       `json:"content"` // Base64 content
	Annotations  []Annotation `json:"annotations,omitempty"`
	AnnotationID int          `json:"annotation_id,omitempty"`
	IsAdditional bool         `json:"is_additional"`
	RequiredFields
}

// IsValid returns if the current structure is valid (and can be saved) or not,
// as it handles a specific information that could be a danger when loading the
// information from DB: cyclic load with AnnotationID, if the Annotation entity
// has already been linked to the current Photo entity...
func (p Photo) IsValid() bool {
	return !(p.IsAdditional && len(p.Annotations) > 0)
}

// PhotoCreationRequest is a simple structure that contains all the information
// to create a photo entity.
// The content of the data is returned as a raw base64 string.
type PhotoCreationRequest struct {
	Data         string `json:"data"`
	IsAdditional bool   `json:"is_additional,omitempty"`
	AnnotationID int    `json:"annotation_id,omitempty"`
}

// IsValid returns if the current structure is valid from the user (and can be saved) or not,
// as it handles a specific information that could be a danger when loading the
// information from DB: cyclic load with AnnotationID, if the Annotation entity
// has already been linked to the current Photo entity...
func (p PhotoCreationRequest) IsValid() bool {
	return p.IsAdditional && p.AnnotationID != 0 || !p.IsAdditional
}

// PhotoCreationResponse is a simple structure that contains all the information
// when a photo entity has been created
type PhotoCreationResponse struct {
	ID string `json:"id"`
}

// PhotoGetResponse is a structure that contains all the information
// for non-additional Photo type
type PhotoGetResponse struct {
	Photo       Photo                   `json:"photo"`
	Annotations []AnnotationGetResponse `json:"annotations,omitempty"`
}

// PhotoGetResponse is a structure that contains all the information
// for additional Photo type ony
type AdditionalPhotoGetResponse struct {
	Photo Photo `json:"additional_photo"`
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
	ID               string                       `json:"id"`
	AdditionalPhotos []AdditionalPhotoGetResponse `json:"additional_photos"`
}

// AnnotationGetResponse is a simple structure that contains all the information
// about an Annotation entity
type AnnotationGetResponse struct {
	Annotation Annotation `json:"annotation"`
}
