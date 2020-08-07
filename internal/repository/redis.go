package repository

import (
	"errors"
	"fmt"
	"strconv"
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

func (cr *cacheRepo) Del(key string) error {
	_, err := cr.cache.Do("DEL", key)
	return err
}

func (cr *cacheRepo) TTL(key string) (int64, error) {
	ttl, err := cr.cache.Do("TTL", key)
	if err != nil {
		return 0, err
	}

	switch ttl := ttl.(type) {
	case []byte:
		res, err := strconv.ParseInt(string(ttl), 10, 64)
		return res, err
	case string:
		res, err := strconv.ParseInt(ttl, 10, 64)
		return res, err
	case nil:
		return 0, nil
	}
	return 0, err
}

func (cr *cacheRepo) GetInt64(key string) (int64, error) {
	ttl, err := cr.cache.Do("GET", key)
	if err != nil {
		return 0, err
	}

	switch ttl := ttl.(type) {
	case []byte:
		res, err := strconv.ParseInt(string(ttl), 10, 64)
		return res, err
	case string:
		res, err := strconv.ParseInt(ttl, 10, 64)
		return res, err
	case nil:
		return 0, nil
	}
	return 0, err
}
