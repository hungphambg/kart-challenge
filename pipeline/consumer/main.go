package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	redis "github.com/go-redis/redis/v8"
	kafka "github.com/segmentio/kafka-go"
)

// CouponState tracks which topics a couponLib code has been seen in
type CouponState struct {
	mu          sync.Mutex
	couponMap   map[string]map[string]struct{} // couponCode -> {topic1, topic2, ...}
	redisClient *redis.Client
}

func NewCouponState(redisAddr string) *CouponState {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr, // Redis address from docker-compose.pipeline.yaml
		Password: "",        // no password set
		DB:       0,         // use default DB
	})
	return &CouponState{
		couponMap:   make(map[string]map[string]struct{}),
		redisClient: rdb,
	}
}

// AddObservation records that a couponCode was seen in a given topic
func (cs *CouponState) AddObservation(couponCode, topic string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if _, ok := cs.couponMap[couponCode]; !ok {
		cs.couponMap[couponCode] = make(map[string]struct{})
	}
	cs.couponMap[couponCode][topic] = struct{}{}
}

// GetObservations returns the set of topics a couponLib code has been seen in
func (cs *CouponState) GetObservations(couponCode string) map[string]struct{} {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.couponMap[couponCode]
}

// GetCouponCodesByObservationCount returns couponLib codes observed in 'count' or more topics
func (cs *CouponState) GetCouponCodesByObservationCount(count int) []string {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	var validCodes []string
	for code, topics := range cs.couponMap {
		if len(topics) >= count {
			validCodes = append(validCodes, code)
		}
	}
	return validCodes
}

// SaveValidCouponsToRedis saves a list of couponLib codes to Redis as a Set
// This operation is atomic using temporary keys and RENAME
func (cs *CouponState) SaveValidCouponsToRedis(ctx context.Context, validCodes []string) error {
	const (
		redisKeyCurrent = "valid_coupon_codes"
		redisKeyTemp    = "valid_coupon_codes_temp"
	)

	pipe := cs.redisClient.Pipeline()
	pipe.Del(ctx, redisKeyTemp) // Clear temp set
	if len(validCodes) > 0 {
		pipe.SAdd(ctx, redisKeyTemp, validCodes) // Add all valid codes to temp set
	}
	pipe.Rename(ctx, redisKeyTemp, redisKeyCurrent) // Atomically rename temp set to current

	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to save valid coupons to Redis: %w", err)
	}
	log.Printf("Saved %d valid couponLib codes to Redis key '%s'", len(validCodes), redisKeyCurrent)
	return nil
}

func main() {
	fmt.Println("Coupon Consumer Service started.")

	kafkaBrokers := []string{"kafka:9092"}
	topics := []string{"couponbase1.gz", "couponbase2.gz", "couponbase3.gz"}
	groupID := "couponLib-consumer-group"
	redisAddr := "redis:6379" // Redis address from docker-compose.pipeline.yaml

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        kafkaBrokers,
		GroupID:        groupID,
		GroupTopics:    topics,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: time.Second,
		Logger:         kafka.LoggerFunc(logf),
		ErrorLogger:    kafka.LoggerFunc(logf),
	})
	defer reader.Close()

	couponState := NewCouponState(redisAddr)

	fmt.Printf("Subscribed to topics: %v with consumer group: %s\n", topics, groupID)

	// Goroutine to periodically identify and log valid coupons
	go func() {
		ticker := time.NewTicker(5 * time.Second) // Check every 5 seconds
		defer ticker.Stop()
		for range ticker.C {
			valid := couponState.GetCouponCodesByObservationCount(2)
			if len(valid) > 0 {
				fmt.Printf("[%s] Valid coupons (observed in >= 2 topics): %v\n", time.Now().Format("15:04:05"), valid)
				err := couponState.SaveValidCouponsToRedis(context.Background(), valid)
				if err != nil {
					log.Printf("Error saving valid coupons to Redis: %v", err)
				}
			}
		}
	}()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		couponCode := string(m.Value)
		couponState.AddObservation(couponCode, m.Topic)

		fmt.Printf("Consumed message from topic %s, partition %d, offset %d: %s = %s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), couponCode)
	}

	fmt.Println("Coupon Consumer Service stopped.")
}

func logf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
