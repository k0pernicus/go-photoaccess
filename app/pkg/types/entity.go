package types

import "time"

// Photo is the raw structure of a photo, which can contains annotations
// A Photo entity can contains annotations, or an annotation id, but it can't contains both.
type Photo struct {
	ID           int          `json:"id"`
	Data         []byte       `json:"content"`
	Annotations  []Annotation `json:"annotations,omitempty"`
	AnnotationID int          `json:"annotation_id,omitempty"`
	IsAdditional bool         // No need to export it, as `AnnotationID` could be exported as well
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

// IsValid returns if the current structure is valid (and can be saved) or not,
// as it handles a specific information that could be a danger when loading the
// information from DB: cyclic load with AnnotationID, if the Annotation entity
// has already been linked to the current Photo entity...
func (p Photo) IsValid() bool {
	return !(p.IsAdditional && len(p.Annotations) > 0)
}

// PhotoCreationRequest is a simple structure that contains all the informations
// to create a photo entity
type PhotoCreationRequest struct {
	Data         []byte `json:"data"`
	IsAdditional bool   `json:"is_additional,omitempty"`
	AnnotationID int    `json:"annotation_id,omitempty"`
}

// PhotoCreationResponse is a simple structure that contains all the informations
// when a photo entity has been created
type PhotoCreationResponse struct {
	ID string `json:"id"`
}

// An annotation must be linked to a photo, and contains text
type Annotation struct {
	ID        int       `json:"id"`
	PhotoID   int       `json:"photo_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AnnotationCreationRequest is a simple structure that contains an annotation
// for an existing photo
type AnnotationCreationRequest struct {
	Text string `json:"text"`
}

// AnnotationCreationResponse is a simple structure that contains all the informations
// when an annotation entity has been created
type AnnotationCreationResponse struct {
	ID string `json:"id"`
}
