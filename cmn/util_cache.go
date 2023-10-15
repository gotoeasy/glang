package cmn

import (
	"container/list"
	"sync"
)

type LruCache struct {
	maxEntries int
	cache      map[string]*list.Element
	lruList    *list.List
	mutex      sync.Mutex
}

type lruCacheEntry struct {
	key   string
	value string
}

// 内存缓存(最近使用优先)
func NewLruCache(maxEntries int) *LruCache {
	if maxEntries < 16 {
		maxEntries = 16
	}
	if maxEntries > 100*10000 {
		maxEntries = 100 * 10000
	}
	return &LruCache{
		maxEntries: maxEntries,
		cache:      make(map[string]*list.Element, maxEntries),
		lruList:    list.New(),
	}
}

func (c *LruCache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.lruList.MoveToFront(ele)
		return ele.Value.(*lruCacheEntry).value, true
	}

	return "", false
}

func (c *LruCache) Add(key string, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.lruList.Len() >= c.maxEntries {
		c.removeOldest()
	}

	e := c.lruList.PushFront(&lruCacheEntry{key, value})
	c.cache[key] = e
}

func (c *LruCache) removeOldest() {
	if c.lruList.Len() == 0 {
		return
	}

	ele := c.lruList.Back()
	if ele != nil {
		c.lruList.Remove(ele)
		delete(c.cache, ele.Value.(*lruCacheEntry).key)
	}
}
