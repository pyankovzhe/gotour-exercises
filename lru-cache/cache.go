package lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListElement
	mu       sync.RWMutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListElement, capacity),
	}
}

// The return value is a boolean flag whether the item was present in the cache.
func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		el.Value = value
		c.queue.MoveToFront(el)

		return true
	} else {
		if c.queue.Len() == c.capacity {
			c.queue.Remove(c.queue.Last())
			delete(c.items, key)
		}

		newEl := c.queue.PushFront(value)
		c.items[key] = newEl

		return false
	}
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if el, ok := c.items[key]; ok {
		c.queue.MoveToFront(el)
		return el.Value, ok
	}

	return nil, false
}

func (c *lruCache) Clear() {
	// TODO: implement
}
