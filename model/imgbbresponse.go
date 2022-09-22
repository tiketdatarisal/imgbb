package model

type ImgBBResponse struct {
	Data    *ImgBBData `json:"data,omitempty"`
	Success bool       `json:"success"`
	Status  int        `json:"status"`
}
