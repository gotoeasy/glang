package cmn

import (
	"sync"
	"time"
)

// 缓存项结构
type CacheItem struct {
	Value      any
	expiration time.Time
}

// 缓存结构
type Cache struct {
	mu       sync.RWMutex
	items    map[string]CacheItem
	duration time.Duration
}

// 新建内存缓存(有存活期，会定期清理失效缓存)
func NewCache(duration time.Duration) *Cache {
	cache := &Cache{
		items:    make(map[string]CacheItem),
		duration: duration,
	}
	go cache.cleanupExpired()
	return cache
}

// 添加缓存项
func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expiration := time.Now().Add(c.duration)
	c.items[key] = CacheItem{Value: value, expiration: expiration}
}

// 获取缓存项
func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	if time.Now().After(item.expiration) {
		c.Delete(key)
		return nil, false
	}
	return item.Value, true
}

// 删除缓存项
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// 每5分钟检查清理一次过期缓存项
func (c *Cache) cleanupExpired() {
	ticker := time.NewTicker(time.Minute * 5)
	for {
		<-ticker.C
		c.mu.Lock()
		for key, item := range c.items {
			if time.Now().After(item.expiration) {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
