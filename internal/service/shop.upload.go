package service

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/yeremiaaryo/platform/internal/entity"
)

func (ss *shopSvc) UploadImage(ctx context.Context, image []byte) (*entity.UploadImageResponse, error) {
	base64Img := base64.StdEncoding.EncodeToString(image)
	resp, err := ss.cloudinaryRepo.UploadImage(ctx, base64Img)
	if err != nil {
		log.Println("error when upload image to cloudinary", err.Error())
		return nil, err
	}
	return resp, nil
}
