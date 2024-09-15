package cmn

import (
	"sync"
	"time"
)

// 缓存项结构
type cacheItem struct {
	value      any
	expiration time.Time
}

// 缓存结构
type Cache struct {
	mu         sync.RWMutex
	items      map[string]cacheItem
	duration   time.Duration
	ignorecase bool
}

// 新建内存缓存(有存活期，会定期清理失效缓存)
func NewCache(duration time.Duration, ignorecases ...bool) *Cache {
	cache := &Cache{
		items:      make(map[string]cacheItem),
		duration:   duration,
		ignorecase: false,
	}
	if len(ignorecases) > 0 {
		cache.ignorecase = ignorecases[0]
	}

	go cache.cleanupExpired()
	return cache
}

// 添加缓存项
func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expiration := time.Now().Add(c.duration)
	if c.ignorecase {
		key = ToLower(key)
	}
	c.items[key] = cacheItem{value: value, expiration: expiration}
}

// 获取缓存项
func (c *Cache) Get(key string) (any, bool) {
	if c.ignorecase {
		key = ToLower(key)
	}

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
	return item.value, true
}

// 删除缓存项
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.ignorecase {
		key = ToLower(key)
	}
	delete(c.items, key)
}

// 取缓存的所有有效期内的键
func (c *Cache) Keys() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	var keys []string
	for key, item := range c.items {
		if !time.Now().After(item.expiration) {
			keys = append(keys, key)
		}
	}
	return keys
}

// 取缓存的所有有效期内的值
func (c *Cache) Values() []any {
	c.mu.Lock()
	defer c.mu.Unlock()

	var values []any
	for _, item := range c.items {
		if !time.Now().After(item.expiration) {
			values = append(values, item.value)
		}
	}
	return values
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
