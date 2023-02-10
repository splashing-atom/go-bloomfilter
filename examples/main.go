package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	bloomfilter "github.com/splashing-atom/go-bloomfilter"
)

func main() {
	MemoryBloomFilterExample()
	RedisBloomFilterExample()
}

func MemoryBloomFilterExample() {
	bloom := bloomfilter.NewMemoryBloomFilter(10000)
	bs := []byte("bloom")
	_ = bloom.Put(bs)
	exists, err := bloom.MightContain(bs)
	fmt.Println(exists, err)
}

func RedisBloomFilterExample() {
	cli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	key := "redis bloomfilter"
	ctx := context.Background()
	bloom := bloomfilter.NewRedisBloomFilter(cli, "test", 10000)
	bs := []byte(key)
	_ = bloom.PutCtx(ctx, bs)
	exists, err := bloom.MightContainCtx(ctx, bs)
	fmt.Println(exists, err)
}
