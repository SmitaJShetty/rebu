// CartTripCacheService construct for services provided for caching
type CartTripCacheService interface {
	Get(key string) (*model.CartTrip, error)
	Set(key string, value *model.CartTrip) error
	Del(key string) error
}