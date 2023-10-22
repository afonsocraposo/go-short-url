package helpers

import (
	"context"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
)

var client = redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
var ctx = context.Background()

func SetVar(key string, value string) {
    err := client.Set(ctx, key, value, 24 * time.Hour).Err()
    if  err != nil {
        panic(err)
    }
}

func GetVar(key string) string {
    val, err := client.Get(ctx, key).Result()
    if err != nil {
        panic(err)
    }
    return val
}

func CloseRedis() {
    fmt.Println("Closing redis connection")
    defer client.Close()
}
