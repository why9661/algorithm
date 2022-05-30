package lru

import "container/list"

type Lru struct {
	mBytes    int64                         // max memory allowed
	nBytes    int64                         // used memory
	queue     *list.List                    // double linked list
	cache     map[string]*list.Element      // k-v map
	onEvicted func(key string, value Value) // callback function executed when an entry is purged
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(mBytes int64, onEvicted func(string, Value)) *Lru {
	return &Lru{
		mBytes:    mBytes,
		queue:     list.New(),
		cache:     make(map[string]*list.Element),
		onEvicted: onEvicted,
	}
}

func (l *Lru) Add(key string, value Value) {
	if ele, ok := l.cache[key]; ok {
		l.queue.MoveToFront(ele)
		e := ele.Value.(*entry)
		l.nBytes += int64(value.Len()) - int64(e.value.Len())
		if l.nBytes > l.mBytes {
			l.removeOldest()
		}
		e.value = value
	} else {
		l.nBytes += int64(len(key) + value.Len())
		if l.nBytes > l.mBytes {
			l.removeOldest()
		}
		ele := l.queue.PushFront(&entry{key: key, value: value})
		l.cache[key] = ele
	}
}

func (l *Lru) Get(key string) (Value, bool) {
	if ele, ok := l.cache[key]; ok {
		value := ele.Value.(*entry).value
		l.queue.MoveToFront(ele)
		return value, true
	}
	return nil, false
}

func (l *Lru) removeOldest() {
	ele := l.queue.Back()
	if ele != nil {
		e := ele.Value.(*entry)
		delete(l.cache, e.key)
		l.queue.Remove(ele)
		l.nBytes -= int64(len(e.key) + e.value.Len())
		if l.onEvicted != nil {
			l.onEvicted(e.key, e.value)
		}
	}
}
