package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type hashFunc func([]byte) uint32

type Map struct {
	circle  []uint32
	hash    hashFunc
	replica int
	mapping map[uint32]string
}

func New(hash hashFunc, replica int) *Map {
	m := &Map{
		hash:    hash,
		replica: replica,
		mapping: make(map[uint32]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

func (m *Map) Add(keys ...string) {
	for _, v := range keys {
		for i := 0; i < m.replica; i++ {
			hashData := []byte(strconv.Itoa(i) + v)
			hash := m.hash(hashData)
			m.mapping[hash] = v
			m.circle = append(m.circle, hash)
		}
	}

	sort.Slice(m.circle, func(i, j int) bool {
		return m.circle[i] < m.circle[j]
	})
}

func (m *Map) Get(key string) string {
	if len(m.circle) == 0 {
		return ""
	}

	hash := m.hash([]byte(key))

	index := sort.Search(len(m.circle), func(i int) bool {
		return m.circle[i] >= hash
	})

	return m.mapping[m.circle[index%len(m.circle)]]
}
