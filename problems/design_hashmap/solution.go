type Node struct {
	Key   int
	Value int
	Next  *Node
}

type MyHashMap struct {
	Arrays   []*Node
	HashFunc func(key int) int
	Size     uint32
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	size := 100
	hm := MyHashMap{
		Arrays: make([]*Node, size),
		Size:   100,
		HashFunc: func(key int) int {
			return key % size
		},
	}

	return hm
}

func (this *MyHashMap) NewNode(key int, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	var prev *Node
	index := this.HashFunc(key)
	entry := this.Arrays[index]
	for entry != nil && entry.Key != key {
		prev = entry
		entry = entry.Next
	}

	if entry == nil {
		if prev == nil {
			prev = this.NewNode(key, value)
			this.Arrays[index] = prev
		} else {
			entry = this.NewNode(key, value)
			prev.Next = entry
		}
	} else {
		entry.Value = value
	}

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	entry := this.Arrays[this.HashFunc(key)]
	for entry != nil {
		if entry.Key == key {
			return entry.Value
		}

		entry = entry.Next
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	var prev *Node
	entry := this.Arrays[this.HashFunc(key)]
	for entry != nil && entry.Key != key {
		prev = entry
		entry = entry.Next
	}

	if entry == nil {
		return
	}

	if prev == nil {
		this.Arrays[this.HashFunc(key)] = entry.Next
	} else {
		prev.Next = entry.Next
	}

	return
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */