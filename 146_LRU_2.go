// 定义双向链表
type MyListNode struct {
	key, value int
	pre, next  *MyListNode
}

type LRUCache struct {
	size, capacity           int
	cache                    map[int]*MyListNode
	protectHead, protectTail *MyListNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		size:        0,
		capacity:    capacity,
		cache:       map[int]*MyListNode{},
		protectHead: &MyListNode{},
		protectTail: &MyListNode{},
	}

	lruCache.protectHead.next = lruCache.protectTail
	lruCache.protectTail.pre = lruCache.protectHead
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToTail(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToTail(node)
		return
	}
	if this.size == this.capacity {
		headNode := this.protectHead.next
		this.remove(headNode)
		delete(this.cache, headNode.key)
		this.size--
	}

	node := &MyListNode{
		key:   key,
		value: value,
	}

	this.addToTail(node)
	this.cache[key] = node
	this.size++
}

// 删除指定节点
func (this *LRUCache) remove(node *MyListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre

	node.pre = nil
	node.next = nil
}

// 添加到末尾
func (this *LRUCache) addToTail(node *MyListNode) {
	this.protectTail.pre.next = node
	node.pre = this.protectTail.pre

	node.next = this.protectTail
	this.protectTail.pre = node
}

// 从当前位置移动到末尾
func (this *LRUCache) moveToTail(node *MyListNode) {
	this.remove(node)
	this.addToTail(node)
}