package repository

type RedisConn interface {
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
}

func (cr *cacheRepo) Get(key string) (string, error) {
	res, err := cr.cache.Do("GET", key)
	if err != nil {
		return "", err
	}

	return res.(string), err
}

func (cr *cacheRepo) Set(key, value string, expired int) error {
	_, err := cr.cache.Do("SETEX", key, expired, value)
	return err
}
