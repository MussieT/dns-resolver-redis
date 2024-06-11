package connect

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/redis/go-redis/v9"
)

func redisConf() redis.Client {
	fmt.Println(os.Getenv("REDIS_ADDRESS"))
	client := redis.NewClient(&redis.Options{
        Addr:	  os.Getenv("REDIS_ADDRESS"),
        Password: os.Getenv("PASSWORD"), // no password set
        DB:		  0,  // use default DB
    })

	return *client
}

func SetValue(domain string, ipkey []net.IP) {
	client := redisConf()

	ctx := context.Background()

	data, _ := json.Marshal(ipkey)

	err := client.Set(ctx, domain, data, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, domain).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)
}