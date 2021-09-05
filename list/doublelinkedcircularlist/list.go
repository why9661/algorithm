package doublelinkedcircularlist

type Element struct {
	prev, next *Element
	value      interface{}
	list       *List
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.head {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.head {
		return p
	}
	return nil
}

type List struct {
	head Element //头结点
	len  int
}

// 初始化(清空)链表
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

// 返回链表的第一个结点
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// 返回链表的最后一个结点
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.prev
}

// 插入结点，并返回该结点
func (l *List) Insert(e, at *Element) *Element {
	e.next = at.next
	e.prev = at
	e.next.prev = e
	e.prev.next = e
	e.list = l
	l.len++
	return e
}

// 插入值，并返回该结点
func (l *List) InsertValue(value interface{}, at *Element) *Element {
	return l.Insert(&Element{value: value}, at)
}

// 删除结点，并返回该结点
func (l *List) Remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
	e.list = nil
	l.len--
	return e
}

// 移动结点，并返回该结点
func (l *List) Move(e, at *Element) *Element {
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

// 在链表头部插入值，并返回生成的结点
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.InsertValue(v, &l.head)
}

// 在链表尾部插入值，并返回生成的结点
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.InsertValue(v, l.head.prev)
}

// 将结点移动到链表头
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.head.next == e {
		return
	}
	l.Move(e, &l.head)
}

// 将结点移动到链表尾
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.head.prev == e {
		return
	}
	l.Move(e, l.head.prev)
}
