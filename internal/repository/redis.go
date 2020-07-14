package repository

import (
	"errors"
	"fmt"
)

type RedisConn interface {
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
}

func (cr *cacheRepo) Get(key string) (string, error) {
	res, err := cr.cache.Do("GET", key)
	if err != nil {
		return "", err
	}
	switch res := res.(type) {
	case []byte:
		return string(res), nil
	case string:
		return res, nil
	case nil:
		return "", errors.New("Nil value")
	}

	return "", fmt.Errorf("unexpected type for String, got type %T", res)
}

func (cr *cacheRepo) Set(key, value string, expired int) error {
	_, err := cr.cache.Do("SETEX", key, expired, value)
	return err
}
