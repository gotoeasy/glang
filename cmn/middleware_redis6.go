package cmn

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// Redis6客户端结构体
type Redis6Client struct {
	ctx context.Context
	rdb *redis.Client
}

// 创建Redis6客户端对象
func NewRedis6Client(opt *redis.Options) *Redis6Client {
	// redis6是github.com/go-redis/redis/v8，redis7是github.com/go-redis/redis/v9
	return &Redis6Client{
		ctx: context.Background(),
		rdb: redis.NewClient(opt),
	}
}

// 设定
func (r *Redis6Client) Set(key string, value string, expiration time.Duration) error {
	return r.rdb.Set(r.ctx, key, value, expiration).Err()
}

// 删除
func (r *Redis6Client) Del(key string) error {
	return r.rdb.Del(r.ctx, key).Err()
}

// 判断是否存在
func (r *Redis6Client) Exists(key string) *redis.IntCmd {
	return r.rdb.Exists(r.ctx, key)
}

// 获取（空时也是error，可通过err==redis.Nil判断）
func (r *Redis6Client) Get(key string) (string, error) {
	return r.rdb.Get(r.ctx, key).Result()
}

// 关闭客户端
func (r *Redis6Client) Close(key string) error {
	return r.rdb.Close()
}
