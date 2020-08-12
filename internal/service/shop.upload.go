package service

import (
	"context"
	"encoding/base64"
	"log"
)

func (ss *shopSvc) UploadImage(ctx context.Context, image []byte) (interface{}, error) {
	base64Img := base64.StdEncoding.EncodeToString(image)
	resp, err := ss.cloudinaryRepo.UploadImage(ctx, base64Img)
	if err != nil {
		log.Println("error when upload image to cloudinary", err.Error())
		return nil, err
	}
	return resp, nil
}
