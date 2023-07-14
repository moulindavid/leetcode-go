package main

type Node struct {
	value int
	key   int
	next  *Node
	prev  *Node
}
type LRUCache struct {
	cache       map[int]*Node
	capacity    int
	lru         *Node
	most_recent *Node
}

func Constructor(capacity int) LRUCache {
	lru_cache := LRUCache{
		capacity:    capacity,
		cache:       make(map[int]*Node),
		lru:         &Node{key: 0, value: 0},
		most_recent: &Node{key: 0, value: 0}}

	lru_cache.lru.next, lru_cache.most_recent.prev = lru_cache.most_recent, lru_cache.lru
	return lru_cache
}

func (this *LRUCache) Remove(node *Node) {
	previous, next := node.prev, node.next
	next.prev, previous.next = previous, next
}

// insert node at right
func (this *LRUCache) Insert(node *Node) {
	previous, next := this.most_recent.prev, this.most_recent
	previous.next, next.prev = node, node
	node.prev, node.next = previous, next
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; ok {
		//update key to be the most most recent
		this.Remove(this.cache[key])
		this.Insert(this.cache[key])
		return this.cache[key].value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		this.Remove(this.cache[key])
	}
	this.cache[key] = &Node{key: key, value: value}
	this.Insert(this.cache[key])

	if len(this.cache) > this.capacity {
		lru := this.lru.next
		this.Remove(lru)
		delete(this.cache, lru.key)
	}
}
