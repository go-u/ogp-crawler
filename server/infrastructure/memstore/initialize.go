package memstore

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

func initialize(pool *redis.Pool) error {
	con := pool.Get()

	// キーをセット
	testKey := "testKey"
	_, err := con.Do("SET", testKey, true, "EX", 1)
	if err != nil {
		log.Println(err)
	}

	// キーが存在することを確認
	rExist, err := redis.Bool(con.Do("GET", testKey))
	if !rExist {
		return err
	}

	// 期限切れであることを確認
	time.Sleep(1500 * time.Millisecond)
	rExist, err = redis.Bool(con.Do("GET", testKey))
	if rExist {
		return err
	}

	// メモリを初期化
	_, err = con.Do("FLUSHALL")
	if err != nil {
		return err
	}

	return nil
}
