package main

import (
	"context"
	"fmt"
	"hash/crc32"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	ShardNum       = 64
	Workers        = 4
	RedisBatchSize = 300
)

type Promo struct {
	Key   string
	Shard int
}

func ShardOf(promo string) int {
	return int(crc32.ChecksumIEEE([]byte(promo)) % ShardNum)
}

// CouponState tracks which topics a coupon code has been seen in
type CouponState struct {
	ctx         context.Context
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
		ctx:         context.Background(),
		couponMap:   make(map[string]map[string]struct{}),
		redisClient: rdb,
	}
}

var batchScript = redis.NewScript(`
local mask = tonumber(ARGV[1])

for i = 1, #KEYS do
  local old = redis.call("GET", KEYS[i])

  if old then
    redis.call("SET", KEYS[i], bit.bor(tonumber(old), mask))
  else
    redis.call("SET", KEYS[i], mask)
  end
end

return #KEYS
`)

// AddObservation records that a couponCode was seen in a given topic
func (cs *CouponState) AddObservation(batchCode []string, fileID int) {
	//mask := int64(1 << fileID)
	pipe := cs.redisClient.Pipeline()
	count := 0
	start := time.Now()
	fmt.Println("start set bit for leng ", len(batchCode))
	for _, code := range batchCode {
		shard := ShardOf(code)
		key := fmt.Sprintf("promo:%d:%s", shard, code)
		pipe.SetBit(cs.ctx, key, int64(fileID), 1)
		count++

		if count >= RedisBatchSize {
			_, err := pipe.Exec(cs.ctx)
			if err != nil {
				log.Fatal(err)
			}
			pipe = cs.redisClient.Pipeline()
			count = 0
		}
	}

	if count > 0 {
		_, err := pipe.Exec(cs.ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("end set bit int time ", time.Since(start))
}

func (cs *CouponState) RedisWorker(id int, rdb *redis.Client, ch <-chan Promo, wg *sync.WaitGroup, fileId int64) {
	defer wg.Done()

	pipe := rdb.Pipeline()
	count := 0
	start := time.Now()

	for promo := range ch {
		pipe.SetBit(cs.ctx, promo.Key, fileId, 1)
		count++

		if count >= BatchSize {
			_, err := pipe.Exec(cs.ctx)
			if err != nil {
				log.Fatal(err)
			}
			pipe = rdb.Pipeline()
			count = 0
		}
	}

	if count > 0 {
		_, err := pipe.Exec(cs.ctx)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf(
		"[worker-%d] done in %s\n",
		id,
		time.Since(start),
	)
}

// GetObservations returns the set of topics a coupon code has been seen in
func (cs *CouponState) GetObservations(couponCode string) map[string]struct{} {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.couponMap[couponCode]
}

// GetCouponCodesByObservationCount returns coupon codes observed in 'count' or more topics
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

// SaveValidCouponsToRedis saves a list of coupon codes to Redis as a Set
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
	log.Printf("Saved %d valid coupon codes to Redis key '%s'", len(validCodes), redisKeyCurrent)
	return nil
}
