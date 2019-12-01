package cacheconfig

type CacheConfiguration struct {
	// PORT             string
	REDIS_URL        string
	SESSION_TIME_OUT int8
}

var CacheConfig CacheConfiguration = CacheConfiguration{
	REDIS_URL:        "redis://localhost",
	SESSION_TIME_OUT: 120,
}
