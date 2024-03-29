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

func (cr *cloudinaryRepo) UploadImage(ctx context.Context, image, folder string) (*entity.UploadImageResponse, error) {
	ts := time.Now().Unix()
	stringToSign := fmt.Sprintf("folder=%s&timestamp=%v%s", folder, ts, entity.CloudinaryAPISecret)
	signature := utils.GenerateSHA1(stringToSign)

	image = "data:image/png;base64," + image
	params := url.Values{}
	params.Add("file", image)
	fileParams := params.Encode()

	bodyData := fmt.Sprintf("%s&folder=%s&api_key=%s&timestamp=%v&signature=%s", fileParams, folder, entity.CloudinaryAPIKey, ts, signature)

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

	result := new(entity.UploadImageResponse)
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
