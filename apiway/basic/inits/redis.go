package inits

import (
	"context"
	"zg6zk1/apiway/basic/global"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ExampleClient() {
	global.Client = redis.NewClient(&redis.Options{
		Addr:     "14.103.243.149:6379",
		Password: "2003225zyh", // no password set
		DB:       0,            // use default DB
	})

}
