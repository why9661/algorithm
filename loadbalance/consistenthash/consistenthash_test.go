package consistenthash

import (
	"strconv"
	"testing"
)

func TestConsistentHash(t *testing.T) {
	hFunc := func(data []byte) uint32 {
		i, _ := strconv.Atoi(string(data))
		return uint32(i)
	}

	hash := New(hFunc, 3)

	//5: 5,15,25
	//3: 3,13,23
	//1: 1,11,21
	hash.Add("5", "3", "1")

	testCases := map[string]string{
		"2":  "3",
		"4":  "5",
		"6":  "1",
		"26": "1",
	}

	for k, v := range testCases {
		if peer := hash.Get(k); peer != v {
			t.Errorf("Wrong:%s-%s True:%s-%s", k, peer, k, v)
		}
	}
}
