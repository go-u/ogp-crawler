package memstore

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"server/domain/service"
	"time"
)

type MemStoreService struct {
	Pool *redis.Pool
}

func NewMemStoreService(port int) service.MemStoreService {
	optDial := redis.DialConnectTimeout(1 * time.Second)
	optRead := redis.DialReadTimeout(1 * time.Second)
	optWrite := redis.DialWriteTimeout(1 * time.Second)
	address := fmt.Sprintf("localhost:%d", port)
	pool := redis.Pool{
		MaxIdle:   1000,
		MaxActive: 1000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address, optDial, optRead, optWrite)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	// 接続テスト&初期化
	err := initialize(&pool)
	if err != nil {
		log.Fatalln(err)
	}

	memStoreService := &MemStoreService{
		Pool: &pool,
	}
	return memStoreService
}

func (s *MemStoreService) HasCache(field string, id string) (bool, error) {
	con := s.Pool.Get()
	defer con.Close()

	key := field + ":" + id // tweet:001

	exist, err := redis.Bool(con.Do("EXISTS", key))
	if err == redis.ErrNil {
		return false, nil
	}
	if err != nil { // err != redis.ErrNil
		return false, err
	}

	return exist, nil
}

func (s *MemStoreService) Get(field string, id string) (interface{}, error) {
	con := s.Pool.Get()
	defer con.Close()

	key := field + ":" + id // tweet:001
	result, err := con.Do("GET", key)
	if err == redis.ErrNil {
		return nil, nil
	}
	if err != nil { // err != redis.ErrNil
		return nil, err
	}
	return result, nil
}

func (s *MemStoreService) Add(field string, id string, value interface{}, sec int) error {
	con := s.Pool.Get()
	defer con.Close()

	key := field + ":" + id
	_, err := con.Do("SET", key, value, "EX", sec)
	if err != nil {
		return err
	}
	return nil
}
