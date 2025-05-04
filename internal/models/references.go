package models

import "time"

// ReferencesModel References model for related data
type ReferencesModel struct {
	Agencies   []AgencyReference `json:"agencies"`
	Routes     []interface{}     `json:"routes"`
	Situations []interface{}     `json:"situations"`
	StopTimes  []interface{}     `json:"stopTimes"`
	Stops      []interface{}     `json:"stops"`
	Trips      []interface{}     `json:"trips"`
}

// NewEmptyReferences creates a new empty References model with initialized empty slices
func NewEmptyReferences() ReferencesModel {
	return ReferencesModel{
		Agencies:   []AgencyReference{},
		Routes:     []interface{}{},
		Situations: []interface{}{},
		StopTimes:  []interface{}{},
		Stops:      []interface{}{},
		Trips:      []interface{}{},
	}
}

// NewOKResponse is a helper function that returns a successful response.
func NewOKResponse(data interface{}) ResponseModel {
	return NewResponse(200, data, "OK")
}

// NewResponse Helper function to create a standard response
func NewResponse(code int, data interface{}, text string) ResponseModel {
	return ResponseModel{
		Code:        code,
		CurrentTime: time.Now().UnixNano() / int64(time.Millisecond),
		Data:        data,
		Text:        text,
		Version:     2,
	}
}
