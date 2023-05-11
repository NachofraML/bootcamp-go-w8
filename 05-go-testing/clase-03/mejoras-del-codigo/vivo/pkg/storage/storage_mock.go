package storage

// StorageMock Manually Mock Storage
type StorageMock struct {
	Values map[string]interface{}
}

func (sm *StorageMock) GetValue(key string) interface{} {
	return sm.Values[key]
}

func NewStorageMock() *StorageMock {
	return &StorageMock{}
}

// StorageBetterMock Testify Mock Storage, uncomment when used
//type StorageBetterMock struct {
//	mock.Mock
//}
//
//func (sm *StorageBetterMock) GetValue(key string) interface{} {
//	args := sm.Called(key)
//
//	r0 := args.Get(0)
//
//	return r0
//}
//
//func NewStorageBetterMock() *StorageBetterMock {
//	return &StorageBetterMock{}
//}
