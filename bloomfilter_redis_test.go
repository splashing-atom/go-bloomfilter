package bloomfilter

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

func Test_redisBloomFilter_MightContain(t *testing.T) {
	cli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	key := "redis bloomfilter"
	ctx := context.Background()
	bloom := NewRedisBloomFilter(cli, "_test", 100)
	_ = bloom.PutCtx(ctx, []byte(key))
	exists, err := bloom.MightContainCtx(ctx, []byte(key))
	if err != nil {
		t.Error(err)
	}
	if !exists {
		t.Error("key not exists")
	}

	exists, err = bloom.MightContainCtx(ctx, []byte(key+"abc"))
	if err != nil {
		t.Error(err)
	}
	if exists {
		t.Error("key not exists")
	}
}
