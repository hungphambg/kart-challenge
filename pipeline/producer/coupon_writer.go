package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
)

type CouponWriter interface {
	Write(ctx context.Context, batchCode []string)
}

type CouponRedisWriter struct {
	fileIdx     int
	redisClient *redis.Client
}

type KafkaWriter struct {
	writer *kafka.Writer
}

func (c *KafkaWriter) Write(ctx context.Context, batchCode []string) {
	msgs := make([]kafka.Message, 0, len(batchCode))
	for _, code := range batchCode {
		msg := kafka.Message{
			Key:   []byte(code),
			Value: []byte(code),
			Time:  time.Now(),
		}
		msgs = append(msgs, msg)
	}
	err := c.writer.WriteMessages(ctx, msgs...)
	if err != nil {
		fmt.Printf("Failed to write message to Kafka: %v\n", err)
	}
}
