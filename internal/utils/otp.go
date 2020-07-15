package utils

import (
	"crypto/rand"

	"github.com/yeremiaaryo/platform/internal/entity"
)

const otpChars = "1234567890"

func GenerateOTP() (string, error) {
	buffer := make([]byte, entity.OTPLength)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < entity.OTPLength; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
