package lru

import "testing"

/*
 * leetcode 146. LRU缓存机制
 */

func TestLRUCache_Get(t *testing.T) {
	obj := Constructor(3)
	obj.Put(1, 1)
	obj.Put(2, 2)
	if obj.Get(1) != 1 {
		t.Fail()
	}
	if obj.Get(3) != -1 {
		t.Fail()
	}
}

func TestLRUCache_Put(t *testing.T) {
	obj := Constructor(3)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)
	if obj.Get(1) != -1 {
		t.Fail()
	}
	if obj.Get(2) != 2 {
		t.Fail()
	}
	if obj.Get(3) != 3 {
		t.Fail()
	}
	if obj.Get(4) != 4 {
		t.Fail()
	}
	obj.Put(4, 5)
	if obj.Get(4) != 5 {
		t.Fail()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
