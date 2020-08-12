package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/yeremiaaryo/platform/internal/entity"
	"github.com/yeremiaaryo/platform/internal/utils"
)

func (cr *cloudinaryRepo) UploadImage(ctx context.Context, image string) (interface{}, error) {
	ts := time.Now().Unix()
	stringToSign := fmt.Sprintf("timestamp=%v%s", ts, entity.CloudinaryAPISecret)
	signature := utils.GenerateSHA1(stringToSign)

	image = "data:image/png;base64," + image
	params := url.Values{}
	params.Add("file", image)
	fileParams := params.Encode()

	bodyData := fmt.Sprintf("%s&api_key=%s&timestamp=%v&signature=%s", fileParams, entity.CloudinaryAPIKey, ts, signature)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, entity.CloudinaryBaseURL, bytes.NewBufferString(bodyData))
	if err != nil {
		return nil, err
	}

	resp, err := cr.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
