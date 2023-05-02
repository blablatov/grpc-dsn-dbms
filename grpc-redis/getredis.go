package getredis

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

/*type RedisData struct {
	Name        string
	Description string
	Price       string
}*/

var ctx = context.Background()

func GetRedis(Id string, wg sync.WaitGroup, chg chan string) {
	defer wg.Done()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get(ctx, Id).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	chs <- val
}
