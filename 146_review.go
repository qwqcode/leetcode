type Node struct {
	Prev *Node
	Next *Node
	Key int
	Value int
}

type LRUCache struct {
	Tail *Node
	Head *Node
	Cache map[int]*Node
	Capacity int
	Size int
}

func initNode(key, val int) *Node {
	return &Node{Key: key, Value: val}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		Tail: initNode(0, 0),
		Head: initNode(0, 0),
		Size: 0,
		Capacity: capacity,
		Cache: map[int]*Node{},
	}

	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head

	return l
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}
	this.PutFirst(node)
	return node.Value
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.Cache[key]
	if ok {
		this.PutFirst(node)
		node.Value = value
	} else {
		node = initNode(key, value)
		this.AddFirst(node)
		this.Cache[key] = node
		this.Size++
		if this.Size > this.Capacity {
			this.DeleteTail()
			this.Size--
		}
	}
}


func (this *LRUCache) PutFirst(node *Node) {
	this.Delete(node)
	this.AddFirst(node)
}
func (this *LRUCache) AddFirst(node *Node) {
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
}
func (this *LRUCache) Delete(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
func (this *LRUCache) DeleteTail() {
	node := this.Tail.Prev
	this.Delete(node)
	delete(this.Cache, node.Key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */