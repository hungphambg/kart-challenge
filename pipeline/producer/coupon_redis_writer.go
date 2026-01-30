package main

import (
	"context"
	"fmt"
	"hash/crc32"
	"log"

	"github.com/go-redis/redis/v8"
)

const (
	ShardNum       = 64
	RedisBatchSize = 300
)

func NewCouponRedisWriter(redisAddr string, fileIdx int) *CouponRedisWriter {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr, // Redis address from docker-compose.pipeline.yaml
		Password: "",        // no password set
		DB:       0,         // use default DB
	})
	return &CouponRedisWriter{
		fileIdx:     fileIdx,
		redisClient: rdb,
	}
}

func (c *CouponRedisWriter) Write(ctx context.Context, batchCode []string) {
	pipe := c.redisClient.Pipeline()
	count := 0
	for _, code := range batchCode {
		_, key := getRedisKeyByCouponCode(code)
		pipe.SetBit(ctx, key, int64(c.fileIdx), 1)
		count++
		if count >= RedisBatchSize {
			_, err := pipe.Exec(ctx)
			if err != nil {
				log.Fatal(err)
			}
			pipe = c.redisClient.Pipeline()
			count = 0
		}
	}

	if count > 0 {
		_, err := pipe.Exec(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ShardOf(promo string) int {
	return int(crc32.ChecksumIEEE([]byte(promo)) % ShardNum)
}

func getRedisKeyByCouponCode(code string) (shard int, key string) {
	shard = ShardOf(code)
	key = fmt.Sprintf("promo:%d:%s", shard, code)
	return
}
