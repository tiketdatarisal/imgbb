package imgbb

import (
	"bytes"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/tiketdatarisal/imgbb/model"
)

const imgBbUrl = "https://api.imgbb.com/1/upload"

func Upload(apiKey, filename string, data []byte) (*model.ImgBBResponse, error) {
	client := resty.New()
	res, err := client.R().
		SetQueryParam("key", apiKey).
		SetFileReader("image", filename, bytes.NewReader(data)).
		Post(imgBbUrl)
	if err != nil {
		return nil, err
	}

	resBody := model.ImgBBResponse{}
	if err = json.Unmarshal(res.Body(), &resBody); err != nil {
		return nil, err
	}

	return &resBody, nil
}
