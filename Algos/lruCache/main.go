package main

import "fmt"

// Do not edit the class below except for the insertKeyValuePair,
// getValueFromKey, and getMostRecentKey methods. Feel free
// to add new properties and methods to the class.

	type Node struct {
		key   string
		value int
		next  *Node
		prev  *Node
	}

	type LRUCache struct {
		maxSize int
		hash    map[string]*Node
		head    *Node
		curr    *Node
	}

	func NewLRUCache(size int) *LRUCache {
		return &LRUCache{
			maxSize: size,
			hash:    make(map[string]*Node, size),
			head:    nil,
			curr:    nil,
		}
	}

	func (cache *LRUCache) InsertKeyValuePair(key string, value int) {
		if len(cache.hash) == 0 {
			node := Node{
				key:   key,
				value: value,
				next:  nil,
				prev:  nil,
			}

			cache.hash[key] = &node
			cache.head = &node
			cache.curr = &node
			return
		}

		if node, ok := cache.hash[key]; ok {
			node.value = value
			return
		}

		if len(cache.hash) < cache.maxSize {

			node := Node{
				key:   key,
				value: value,
				next:  nil,
				prev:  nil,
			}

			cache.hash[key] = &node
			temp := cache.curr
			cache.curr.next = &node
			cache.curr = &node
			cache.curr.prev = temp

		} else {

			if cache.maxSize == 1 {
				delete(cache.hash, cache.head.key)
				cache.InsertKeyValuePair(key, value)
				return
			}

			temp := cache.head
			cache.head = cache.head.next
			cache.head.prev = nil

			delete(cache.hash, temp.key)

			node := Node{
				key:   key,
				value: value,
				next:  nil,
				prev:  nil,
			}

			temp = cache.curr
			cache.curr.next = &node
			cache.curr = &node
			cache.curr.prev = temp
			cache.hash[key] = &node
		}

	}

	// The second return value indicates whether or not the key was found
	// in the cache.
	func (cache *LRUCache) GetValueFromKey(key string) (int, bool) {
		if node, ok := cache.hash[key]; ok {

			if cache.maxSize == 1 {
				return node.value, true
			}

			if node == cache.head {
				node.next.prev = nil
				cache.head = node.next
				node.next = nil
				node.prev = cache.curr
				cache.curr.next = node
				cache.curr = node
			} else if node != cache.curr {
				prev := node.prev
				prev.next = node.next

				next := node.next
				next.prev = node.prev

				next.next = node
				node.next = nil
				node.prev = next

				cache.curr = node
			}

			return node.value, true

		}

		return 0, false
	}

	// The second return value is false if the cache is empty.
	func (cache *LRUCache) GetMostRecentKey() (string, bool) {
		if len(cache.hash) == 0 {
			return "", false
		}

		return cache.curr.key, true
	}

func main() {
	cache := NewLRUCache(3)

	cache.InsertKeyValuePair("a", 1)
	cache.InsertKeyValuePair("b", 2)
	cache.InsertKeyValuePair("c", 3)

	fmt.Println(cache.GetMostRecentKey())
	fmt.Println(cache.GetValueFromKey("a"))
	fmt.Println(cache.GetMostRecentKey())

	cache.InsertKeyValuePair("d", 4)
	fmt.Println(cache.GetValueFromKey("b"))
	cache.InsertKeyValuePair("a", 5)
	fmt.Println(cache.GetValueFromKey("a"))

	//cache.InsertKeyValuePair("c", 3)
	//
	//cache.InsertKeyValuePair("d", 4)
	//fmt.Println(cache.GetMostRecentKey())
	//
	//cache.InsertKeyValuePair("e", 5)
	//
	//fmt.Println(cache.hash)
	//fmt.Println(cache.GetMostRecentKey())
	//
	//
	//fmt.Println(cache.GetValueFromKey("f"))
	//fmt.Println(cache.GetValueFromKey("e"))

}
