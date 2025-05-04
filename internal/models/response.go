package models

// ResponseModel Base response structure that can be reused
type ResponseModel struct {
	Code        int         `json:"code"`
	CurrentTime int64       `json:"currentTime"`
	Data        interface{} `json:"data"`
	Text        string      `json:"text"`
	Version     int         `json:"version"`
}
