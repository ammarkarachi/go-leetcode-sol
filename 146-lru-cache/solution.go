// https://leetcode.com/problems/lru-cache/description/?envType=company&envId=facebook&favoriteSlug=facebook-thirty-days
package solution

type LRUCache struct {
	Capacity int
	Map      map[int]*DNode
	Head     *DNode
	Tail     *DNode
}

type DNode struct {
	Key   int
	Value int
	Prev  *DNode
	Next  *DNode
}

func (this *LRUCache) addToHead(node *DNode) {
	currentHead := this.Head.Next
	currentHead.Prev = node
	node.Next = currentHead
	this.Head.Next = node
	node.Prev = this.Head
}

func (this *LRUCache) removeNode(node *DNode) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
}

func Constructor(capacity int) LRUCache {
	head := &DNode{}
	tail := &DNode{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Capacity: capacity,
		Map:      make(map[int]*DNode),
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Map[key]; ok {
		this.removeNode(node)
		this.addToHead(node)
		return this.Map[key].Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Map[key]; ok {
		this.Map[key].Value = value
		this.removeNode(node)
		this.addToHead(node)
	} else {
		newNode := &DNode{
			Key:   key,
			Value: value,
		}
		this.addToHead(newNode)
		this.Map[key] = newNode
	}
	if len(this.Map) > this.Capacity {
		leastRecentlyUsedNode := this.Tail.Prev
		delete(this.Map, leastRecentlyUsedNode.Key)
		this.removeNode(leastRecentlyUsedNode)
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
