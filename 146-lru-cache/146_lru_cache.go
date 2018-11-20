package leetcode

/*
Source: https://leetcode.com/problems/lru-cache
Test Case:
["LRUCache","put","put","get","put","get","put","get","get","get"]
[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
*/

// DoublyLinkedList is a doubly linked list
type DoublyLinkedList struct {
	Prev *DoublyLinkedList
	Next *DoublyLinkedList
	Val  int
	Key  int
}

// LRUCache is a LRU cache
type LRUCache struct {
	capacity int
	length   int
	head     *DoublyLinkedList
	tail     *DoublyLinkedList
	hash     map[int]*DoublyLinkedList
}

// Constructor initializes a new LRUCache instance
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		hash:     make(map[int]*DoublyLinkedList),
		length:   0,
	}
}

// Get value of key, if not exists, return -1
func (cache *LRUCache) Get(key int) int {
	node := cache.hash[key]
	if node == nil {
		return -1
	}

	cache.moveToHead(node)
	return node.Val
}

// Put a value into the cache
func (cache *LRUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}

	if node := cache.hash[key]; node != nil {
		cache.moveToHead(node)
		node.Val = value
	} else {
		cache.ensureSpace()

		cache.head = &DoublyLinkedList{
			Val:  value,
			Key:  key,
			Prev: nil,
			Next: cache.head,
		}

		if cache.head.Next != nil {
			cache.head.Next.Prev = cache.head
		}

		if cache.tail == nil {
			cache.tail = cache.head
		}

		cache.hash[key] = cache.head
		cache.length++
	}
}

func (cache *LRUCache) ensureSpace() {
	if cache.length == 0 || cache.capacity == 0 || cache.length < cache.capacity {
		return
	}

	node := cache.tail

	if cache.length == 1 {
		cache.head = nil
		cache.tail = nil
	} else {
		cache.tail = node.Prev
		cache.tail.Next = nil
		node.Prev = nil
	}

	delete(cache.hash, node.Key)
	cache.length--
}

func (cache *LRUCache) moveToHead(node *DoublyLinkedList) {
	if node == nil || cache.hash[node.Key] != node || cache.head == node {
		return
	}

	node.Prev.Next = node.Next

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if cache.tail == node {
		cache.tail = node.Prev
	}

	node.Next = cache.head
	node.Next.Prev = node
	node.Prev = nil
	cache.head = node
}
