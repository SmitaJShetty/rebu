package cachingstore

import "github.com/stretchr/testify/mock"

//MockCache construct for mock cache
type MockCache struct {
	mock.Mock
}

//Get mock get cache
func (m *MockCache) Get(key string) (int, error) {
	return 2, nil
}

//Set mock set cache
func (m *MockCache) Set(key string, value int) error {
	return nil
}

//Del mock delete cache
func (m *MockCache) Del(key string) error {
	return nil
}

//ClearCache mock clear cache
func (m *MockCache) ClearCache() error {
	return nil
}

//NewMockCache constructor for mock cache
func NewMockCache() *MockCache {
	return &MockCache{}
}
