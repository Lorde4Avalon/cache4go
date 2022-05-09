package cache4go

import (
	"log"
	"sync"
)

type CacheTable struct {
	//seft lock
	sync.RWMutex

	//table name
	name string
	//table's items
	items map[interface{}]*CacheItem

	loger *log.Logger

	

}

//table's name
func (table *CacheTable) Name() string {
	return table.name
}

//get table's item count
func (table *CacheTable) Count() int {
	table.RLock()
	defer table.RUnlock()
	return len(table.items)
}


