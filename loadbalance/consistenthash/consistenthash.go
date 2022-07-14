package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type cHash struct {
	opt     Option
	circle  []uint32
	mapping map[uint32]string
}

type Option struct {
	hash    hashFunc
	replica int
}

type hashFunc func([]byte) uint32

type modOption func(*Option)

func WithHashFunc(hash hashFunc) modOption {
	return func(option *Option) {
		option.hash = hash
	}
}

func WithReplica(replica int) modOption {
	return func(option *Option) {
		option.replica = replica
	}
}

func NewCHash(modOption ...modOption) *cHash {
	instance := &cHash{
		opt: Option{
			hash:    crc32.ChecksumIEEE,
			replica: 0,
		},
		mapping: make(map[uint32]string),
	}

	for _, fn := range modOption {
		fn(&instance.opt)
	}

	return instance
}

func (c *cHash) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < c.opt.replica; i++ {
			hashData := []byte(strconv.Itoa(i) + key)
			hash := c.opt.hash(hashData)
			c.circle = append(c.circle, hash)
			c.mapping[hash] = key
		}
	}

	sort.Slice(c.circle, func(i, j int) bool {
		return c.circle[i] < c.circle[j]
	})
}

func (c *cHash) Get(key string) string {
	if len(c.circle) == 0 {
		return ""
	}

	hash := c.opt.hash([]byte(key))
	index := sort.Search(len(c.circle), func(i int) bool {
		return c.circle[i] >= hash
	})

	return c.mapping[c.circle[index%len(c.circle)]]
}

//type hashFunc func([]byte) uint32
//
//type Map struct {
//	circle  []uint32
//	hash    hashFunc
//	replica int
//	mapping map[uint32]string
//}
//
//func New(hash hashFunc, replica int) *Map {
//	m := &Map{
//		hash:    hash,
//		replica: replica,
//		mapping: make(map[uint32]string),
//	}
//
//	if m.hash == nil {
//		m.hash = crc32.ChecksumIEEE
//	}
//
//	return m
//}
//
//func (m *Map) Add(keys ...string) {
//	for _, v := range keys {
//		for i := 0; i < m.replica; i++ {
//			hashData := []byte(strconv.Itoa(i) + v)
//			hash := m.hash(hashData)
//			m.mapping[hash] = v
//			m.circle = append(m.circle, hash)
//		}
//	}
//
//	sort.Slice(m.circle, func(i, j int) bool {
//		return m.circle[i] < m.circle[j]
//	})
//}
//
//func (m *Map) Get(key string) string {
//	if len(m.circle) == 0 {
//		return ""
//	}
//
//	hash := m.hash([]byte(key))
//
//	index := sort.Search(len(m.circle), func(i int) bool {
//		return m.circle[i] >= hash
//	})
//
//	return m.mapping[m.circle[index%len(m.circle)]]
//}
