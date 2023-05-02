package setredis

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

/*type RedisData struct {
	Id          string
	Name        string
	Description string
	Price       string
}*/

var ctx = context.Background()

func SetRedis(Id, Name string, wg sync.WaitGroup, chs chan string) bool {
	defer wg.Done()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, Id, Name, 0).Err()
	if err != nil {
		panic(err)
	}
	chs <- "Succeed"
	return true
}
