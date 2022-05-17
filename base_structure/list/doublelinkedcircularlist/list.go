package doublelinkedcircularlist

type Node struct {
	prev, next *Node
	value      interface{}
	list       *List
}

func (e *Node) Next() *Node {
	if p := e.next; e.list != nil && p != &e.list.head {
		return p
	}
	return nil
}

func (e *Node) Prev() *Node {
	if p := e.prev; e.list != nil && p != &e.list.head {
		return p
	}
	return nil
}

type List struct {
	head Node //头结点
	len  int
}

// Init 初始化(清空)链表
func (l *List) Init() *List {
	l.head.next = &l.head
	l.head.prev = &l.head
	l.len = 0
	return l
}

func (l *List) lazyInit() {
	if l.head.next == nil {
		l.Init()
	}
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

// Front 返回链表的第一个结点
func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// Back 返回链表的最后一个结点
func (l *List) Back() *Node {
	if l.len == 0 {
		return nil
	}
	return l.head.prev
}

// Insert 插入结点，并返回该结点
func (l *List) Insert(e, at *Node) *Node {
	e.next = at.next
	e.prev = at
	e.next.prev = e
	e.prev.next = e
	e.list = l
	l.len++
	return e
}

// InsertValue 插入值，并返回该结点
func (l *List) InsertValue(value interface{}, at *Node) *Node {
	return l.Insert(&Node{value: value}, at)
}

// Remove 删除结点，并返回该结点
func (l *List) Remove(e *Node) *Node {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
	e.list = nil
	l.len--
	return e
}

// Move 移动结点，并返回该结点
func (l *List) Move(e, at *Node) *Node {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = at.next
	e.prev = at
	e.next.prev = e
	e.prev.next = e

	return e
}

// PushFront 在链表头部插入值，并返回生成的结点
func (l *List) PushFront(v interface{}) *Node {
	l.lazyInit()
	return l.InsertValue(v, &l.head)
}

// PushBack 在链表尾部插入值，并返回生成的结点
func (l *List) PushBack(v interface{}) *Node {
	l.lazyInit()
	return l.InsertValue(v, l.head.prev)
}

// MoveToFront 将结点移动到链表头
func (l *List) MoveToFront(e *Node) {
	if e.list != l || l.head.next == e {
		return
	}
	l.Move(e, &l.head)
}

// MoveToBack 将结点移动到链表尾
func (l *List) MoveToBack(e *Node) {
	if e.list != l || l.head.prev == e {
		return
	}
	l.Move(e, l.head.prev)
}
