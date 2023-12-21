package main

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	queue    *list.List                    // doubly ll is reperesnted by queue
	elements map[interface{}]*list.Element // map to point to the elements in queue
}

// entry to holdd keyvalue
type entry struct {
	key   interface{}
	value interface{}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		queue:    list.New(),
		elements: make(map[interface{}]*list.Element),
	}
}

// Get retrieves the value for a key
func (c *LRUCache) Get(key interface{}) interface{} {
	// if key is present
	if e, found := c.elements[key]; found {
		// move to front as this is the recently used element
		c.queue.MoveToFront(e)
		return e.Value.(*entry).value
	}

	// if key is not present
	return nil
}

// PUT inserts key value into the LRUCache
func (c *LRUCache) Put(key interface{}, value interface{}) {
	// check if key is already present
	if e, found := c.elements[key]; found {
		// move to front as recently used
		c.queue.MoveToFront(e)
		// update value after key is found
		e.Value.(*entry).value = value
		return
	}

	//if cache is full remove lru
	if c.queue.Len() == c.capacity {
		lastelem := c.queue.Back()
		c.queue.Remove(lastelem)
		delete(c.elements, lastelem.Value.(*entry).key)
	}

	//create new entry for key value
	newEntry := &entry{key, value}
	e := c.queue.PushFront(newEntry)
	c.elements[key] = e
}

// GetMRU
func (c *LRUCache) GetMRU() interface{} {
	// if queue is not empty return the front element key
	if c.queue.Front() != nil {
		return c.queue.Front().Value.(*entry).key
	}
	//if queue is empty return nil
	return nil
}

func main() {
	cache := Constructor(3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	// Insert key "b" with value 2.
	cache.Get("a")
	// Returns 1 and makes "a" the most recent key.
	cache.Put("c", 3)
	// Cache is now ["a":1, "b":2, "c":3].
	cache.GetMRU()
	// Should return "a" as the most recent key.
	cache.Put("d", 4) // Evicts "b" and cache is now ["a":1, "c":3, "d":4].
	cache.Get("b")
	// Returns nil, as "b" was evicted.
	cache.Put("a", 5)
	// Updates "a" to 5 and makes it the most recent key.
	cache.Get("a")

}
