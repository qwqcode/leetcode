type LinkNode struct {
	Key int
	Value int
	Prev *LinkNode
	Next *LinkNode
}

type LRUCache struct {
	capacity int
	size int
	head *LinkNode
	tail *LinkNode
	cache map[int]*LinkNode
}

func initNode(k int, v int) *LinkNode {
	return &LinkNode{
		Key: k,
		Value: v,
	}
}


func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		head: initNode(0, 0),
		tail: initNode(0, 0),
		cache: map[int]*LinkNode{},
	}

	// 缺少头尾指针首尾相连
	l.head.Next = l.tail
	l.tail.Prev = l.head

	return l
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.PutFirst(node)
	return node.Value
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.cache[key]
	if ok {
		node.Value = value
		this.PutFirst(node)
	} else {
		node = initNode(key, value)
		this.AddToHead(node)
		this.cache[key] = node
		this.size++
		if this.size > this.capacity {
			this.RemoveTail()
			this.size--
		}
	}
}

func (this *LRUCache) AddToHead(node *LinkNode) {
	node.Prev = this.head
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) PutFirst(node *LinkNode) {
	// 这里直接写成了 add to head，而没有预先 removeNode
	this.RemoveNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) RemoveTail() {
	node := this.tail.Prev
	this.RemoveNode(node)
	delete(this.cache, node.Key)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */