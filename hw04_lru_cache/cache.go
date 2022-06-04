package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	// Cache

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l lruCache) Set(key Key, value interface{}) bool {
	// TODO implement me
	panic("implement me")
}

func (l lruCache) Get(key Key) (interface{}, bool) {
	// TODO implement me
	panic("implement me")
}

func (l lruCache) Clear() {
	// TODO implement me
	panic("implement me")
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
