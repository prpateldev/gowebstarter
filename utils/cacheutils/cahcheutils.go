package cacheutils

import (
	"gowebstarter/configs/cacheconfig"

	"github.com/gomodule/redigo/redis"

	// "github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

var cache redis.Conn

func GetConnection() redis.Conn {
	return cache
}

func SetConnection(c redis.Conn) {
	cache = c
}

func Connect(config cacheconfig.CacheConfiguration) (redis.Conn, error) {
	conn, err := redis.DialURL(config.REDIS_URL)
	return conn, err
}

func Init(config cacheconfig.CacheConfiguration) {
	// Initialize the redis connection to a redis instance running on your local machine
	conn, err := Connect(config)
	if err != nil {
		panic(err)
	}
	// Assign the connection to the package level `Cache` variable
	SetConnection(conn)
}

func SetSession(value string) (string, error) {
	// Create a new random session token
	sessionToken := uuid.New().String()
	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 120 seconds
	_, err := cache.Do("SETEX", sessionToken, cacheconfig.CacheConfig.SESSION_TIME_OUT, value)

	if err != nil {
		return "", err
	}

	return sessionToken, nil
}

func GetSession(sessionToken string) (interface{}, error) {
	response, err := cache.Do("GET", sessionToken)

	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	return response, nil
}

func DeleteSession(sessionToken string) (interface{}, error) {
	return cache.Do("DEL", sessionToken)
}
