package model

import (
	"strconv"
	"time"
)

type ImgBBData struct {
	ID              string `json:"id,omitempty"`
	Title           string `json:"title,omitempty"`
	URLViewer       string `json:"url_viewer,omitempty"`
	URL             string `json:"url,omitempty"`
	DisplayURL      string `json:"display_url,omitempty"`
	DeleteURL       string `json:"delete_url,omitempty"`
	Width           string `json:"width,omitempty"`
	Height          string `json:"height,omitempty"`
	Size            int    `json:"size,omitempty"`
	Time            string `json:"time,omitempty"`
	Expiration      string `json:"expiration,omitempty"`
	Image           *Image `json:"image,omitempty"`
	MediumSizeImage *Image `json:"medium,omitempty"`
	Thumbnail       *Image `json:"thumb,omitempty"`
}

func (i ImgBBData) GetWidth() int64 {
	if i.Width == "" {
		return 0
	}

	w, _ := strconv.ParseInt(i.Width, 10, 64)
	return w
}

func (i ImgBBData) GetHeight() int64 {
	if i.Height == "" {
		return 0
	}

	h, _ := strconv.ParseInt(i.Height, 10, 64)
	return h
}

func (i ImgBBData) GetTimestamp() int64 {
	if i.Time == "" {
		return 0
	}

	t, _ := strconv.ParseInt(i.Time, 10, 64)
	return t
}

func (i ImgBBData) GetTime() time.Time {
	return time.Unix(i.GetTimestamp(), 0)
}

func (i ImgBBData) GetExpiration() *time.Duration {
	if i.Expiration == "" {
		return nil
	}

	d, _ := strconv.ParseInt(i.Expiration, 10, 64)
	if d < 60 || d > 15552000 {
		return nil
	}

	duration := time.Second * time.Duration(d)
	return &duration
}

func (i ImgBBData) GetTimeExpiration() *time.Time {
	d := i.GetExpiration()
	if d == nil {
		return nil
	}

	t := i.GetTime().Add(*d)
	return &t
}
