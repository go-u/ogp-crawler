package service

type MemStoreService interface {
	HasCache(field string, id string) (bool, error)
	Get(field string, id string) (interface{}, error)
	Add(field string, id string, value interface{}, sec int) error
}
