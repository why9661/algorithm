package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type hashFunc func([]byte) uint32

type cHash struct {
	keys    []int
	replica int
	mapping map[int]string
	hash    hashFunc
}

func New(hash hashFunc, replica int) *cHash {
	instance := cHash{
		replica: replica,
		mapping: make(map[int]string),
		hash:    hash,
	}

	if instance.hash == nil {
		instance.hash = crc32.ChecksumIEEE
	}

	return &instance
}

func (ch *cHash) Add(peers ...string) {
	for _, peer := range peers {
		for i := 0; i < ch.replica; i++ {
			hashData := []byte(strconv.Itoa(i) + peer)
			hash := int(ch.hash(hashData))
			ch.keys = append(ch.keys, hash)
			ch.mapping[hash] = peer
		}
	}
	sort.Ints(ch.keys)
}

func (ch *cHash) Get(key string) string {
	if ch.keys == nil {
		return ""
	}

	hash := int(ch.hash([]byte(key)))
	index := sort.Search(len(ch.keys), func(i int) bool {
		return ch.keys[i] >= hash
	})

	return ch.mapping[ch.keys[index%len(ch.keys)]]
}
