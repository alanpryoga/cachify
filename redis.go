package cachify

type redisCache struct {
}

func NewRedisCache() Cache {
	return &redisCache{}
}
