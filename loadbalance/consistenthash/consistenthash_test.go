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

	hash := NewCHash(WithHashFunc(hFunc), WithReplica(3))

	//5: 5,15,25
	//3: 3,13,23
	//1: 1,11,21
	//circle: 1,3,5,11,13,15,21,23,25
	hash.Add("5", "3", "1")

	testCases := map[string]string{
		"2":  "3",
		"4":  "5",
		"6":  "1", // 11--->1
		"26": "1", // len(circle)--->circle[0]=1
	}

	for k, v := range testCases {
		if peer := hash.Get(k); peer != v {
			t.Errorf("Wrong:%s-%s True:%s-%s", k, peer, k, v)
		}
	}
}
