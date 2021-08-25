type Node struct{
    key int
    val int
    next *Node
    prev *Node
}

type DLinkedList struct{
    head,tail *Node
}

func (l *DLinkedList) PushFront(n *Node){
    if l.head==nil{
        l.tail=n
    }else{
        n.next=l.head
        l.head.prev=n
    }
    
    l.head=n
}

func (l *DLinkedList) PushAfter(n *Node){
    if l.tail==nil{
        l.head=n
    }else{
        n.prev=l.tail
        l.tail.next=n
    }
    
    l.tail=n
}

func (l *DLinkedList) MoveToFront(n *Node){
    if n==l.head{
        return
    }
    
    l.Remove(n)
    
    l.head.prev=n
    n.next=l.head
    n.prev=nil
    l.head=n
}

func (l *DLinkedList) Remove(n *Node){
    if n==l.head{
        return
    }
    
    if n==l.tail{
        l.tail.prev.next=nil
        l.tail=n.prev
    }else{
        n.prev.next=n.next
        n.next.prev=n.prev
    }
}

func (l *DLinkedList) Tail() *Node{
    return l.tail
}

type LRUCache struct {
    m map[int]*Node
    cap int
    size int
    list *DLinkedList
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        m: make(map[int]*Node),
        cap: capacity,
        list: &DLinkedList{},
    }
}


func (this *LRUCache) Get(key int) int {
    // check if exist
    if node,ok:=this.m[key];ok{
        this.list.MoveToFront(node)
        return node.val
    }
    
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    // check if exist
    if node,ok:=this.m[key];ok{
        node.val=value
        this.list.MoveToFront(node)
        return
    }
    
    // if key is new
    node:=&Node{key:key,val:value}
    this.list.PushFront(node)
    this.m[key]=node
    this.size++
    
    if this.size>this.cap{
        this.evict()
    }
}

func (this *LRUCache) evict()  {
    if tail:=this.list.Tail();tail!=nil{
        this.removeNode(tail)
    }
}

func (this *LRUCache) removeNode(n *Node)  {
    delete(this.m,n.key)
    this.list.Remove(n)
    this.size--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */