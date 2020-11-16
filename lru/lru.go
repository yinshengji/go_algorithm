package lru
/*
 * leetcode 146. LRU缓存机制
 */

import "container/list"

type LRUCache struct {
	capacity int
	cache map[int]*list.Element
	linkedList *list.List
}

type Pair struct {
	key int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache {
		capacity: capacity,
		linkedList: list.New(),
		cache: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.linkedList.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if elem, ok := this.cache[key]; ok {
		this.linkedList.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if len(this.cache) >= this.capacity {
			delete(this.cache, this.linkedList.Back().Value.(Pair).key)
			this.linkedList.Remove(this.linkedList.Back())
		}
		this.cache[key] = this.linkedList.PushFront(Pair{key, value})
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
