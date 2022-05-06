package cache4go

import (
	"sync"
	"time"
)

type CacheItem struct {
	//seft lock
	sync.RWMutex

	key         interface{}
	data        interface{}
	lifeTime    time.Duration
	createTime  time.Time
	lastAccess  time.Time
	accessCount int64
}

func NewItem(key interface{}, data interface{}, lifeTime time.Duration) *CacheItem {
	time := time.Now()

	return &CacheItem{
		key:         key,
		data:        data,
		lifeTime:    lifeTime,
		createTime:  time,
		lastAccess:  time,
		accessCount: 0,
	}
}

func (item *CacheItem) KeepAlive() {
	item.Lock()
	defer item.Unlock()
	item.lastAccess = time.Now()
	item.accessCount++
}

//Geter
func (item *CacheItem) Key() interface{} {
	//inmuable
	return item.key
}

func (item *CacheItem) Data() interface{} {
	//inmuable
	return item.data
}

func (item *CacheItem) LifeTime() time.Duration {
	//inmuable
	return item.lifeTime
}

func (item *CacheItem) CreateTime() time.Time {
	//inmuable
	return item.createTime
}

func (item *CacheItem) LastAccess() time.Time {
	item.RLock()
	defer item.RUnlock()
	return item.lastAccess
}

func (item *CacheItem) AccessCount() int64 {
	item.RLock()
	defer item.RUnlock()
	return item.accessCount
}
