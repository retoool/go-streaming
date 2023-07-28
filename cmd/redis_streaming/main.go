package main

import (
	"context"
	redis2 "go-streaming/streams/redis"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"go-streaming/streams/flow"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	timer := time.NewTimer(time.Minute)
	go func() {
		<-timer.C
		cancelFunc()
	}()

	config := &redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	}

	source, err := redis2.NewRedisSource(ctx, config, "test")
	if err != nil {
		log.Fatal(err)
	}

	toUpperMapFlow := flow.NewMap(toUpper, 1)

	sink := redis2.NewRedisSink(ctx, config, "test2")

	source.
		Via(toUpperMapFlow).
		To(sink)
}

var toUpper = func(msg *redis.Message) string {
	return strings.ToUpper(msg.Payload)
}
