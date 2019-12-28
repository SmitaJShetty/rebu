
// NewCacheService returns a new cache service
func NewCacheService() CartTripCacheService {
	return &RedisCache{
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

// RedisCache caching service with redis
type RedisCache struct {
	Client *redis.Client
}

func (rc *RedisCache) isCacheActive() (bool, error) {
	_, err := rc.Client.Ping().Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

//Get returns a value mapped to key
func (rc *RedisCache) Get(key string) (*model.CartTrip, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Get: Key's empty")
	}

	isCacheAlive, aliveErr := rc.isCacheActive()
	if aliveErr != nil {
		return nil, aliveErr
	}

	if !isCacheAlive {
		return nil, fmt.Errorf("Get:cache server not alive")
	}

	var value model.CartTrip
	valueBytes, valueErr := rc.Client.Get(key).Bytes()
	if valueErr == redis.Nil {
		log.Println(fmt.Errorf("Get:key (%s) does not exist;", key))
		return nil, nil
	}

	if valueErr != nil {
		return nil, valueErr
	}

	err := json.Unmarshal(valueBytes, &value)
	return &value, err
}

//Set saves a key,value into redis
func (rc *RedisCache) Set(key string, value *model.CartTrip) error {
	if key == "" {
		return fmt.Errorf("Set:Invalid key")
	}

	if value == nil {
		return fmt.Errorf("Set:Invalid value(%v) for key (%s)", value, key)
	}

	if rc.Client == nil {
		return fmt.Errorf("Set:Invalid client reference")
	}

	valueMarshaled, valMarshaledErr := json.Marshal(*value)
	if valMarshaledErr != nil {
		return valMarshaledErr
	}

	setValueErr := rc.Client.Set(key, valueMarshaled, 0).Err()
	if setValueErr != nil {
		return setValueErr
	}

	return nil
}

//Del Deletes a key,value from redis
func (rc *RedisCache) Del(key string) error {
	if key == "" {
		return fmt.Errorf("Set:Invalid key")
	}

	if rc.Client == nil {
		return fmt.Errorf("Set:Invalid client reference")
	}

	setValueErr := rc.Client.Del(key).Err()
	if setValueErr != nil {
		return setValueErr
	}

	return nil
}