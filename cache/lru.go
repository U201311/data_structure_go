package cache

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}
type DLinkedNode struct {
	key, val   int
	prev, next *DLinkedNode
}

func InitDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		val: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		head:     InitDLinkedNode(0, 0),
		tail:     InitDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.removeToHead(node)
	return node.val
}

func (this *LRUCache) removeToHead(node *DLinkedNode) {
	this.remove(node)
	this.addToHead(node)
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := InitDLinkedNode(key, value)
		this.cache[key] = node

	} else {
		node := this.cache[key]
		node.val = value
	}
	this.addToHead(node)
	this.size++
	if this.size > this.capacity {
		this.removeTail()
	}

}
