package cachingstore

// CartTripCacheService construct for services provided for caching
type CartTripCacheService interface {
	Get(key string) (int, error)
	Set(key string, value int) error
	Del(key string) error
	ClearCache() error
}
