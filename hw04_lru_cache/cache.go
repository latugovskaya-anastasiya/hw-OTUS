package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	// если элемент присутствует в словаре
	if ent, ok := l.items[key]; ok {
		ent.Value.(*cacheItem).value = value
		l.queue.MoveToFront(ent)
		return true
	}
	// если элемента нет в словаре
	e := l.queue.PushFront(&cacheItem{key: key, value: value})
	l.items[key] = e

	exceeds := l.queue.Len() > l.capacity
	if exceeds {
		lastEnt := l.queue.Back()
		kv := lastEnt.Value.(*cacheItem)
		l.items[kv.key] = nil
		l.queue.Remove(lastEnt)
		delete(l.items, kv.key)
	}

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	// если элемент присутствует в словаре
	if ent, ok := l.items[key]; ok {
		l.queue.MoveToFront(ent)
		if ent.Value.(*cacheItem).value == nil {
			return nil, false
		}
		value := ent.Value.(*cacheItem).value.(int)
		return value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	for k, v := range l.items {
		v.Value = nil
		v.Prev = nil
		v.Next = nil
		l.items[k] = nil
		delete(l.items, k)
	}
}

func NewCache(capacity int) Cache {
	if capacity <= 0 {
		return nil
	}
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
