package random

import (
	"math/rand"
	"time"
)

type RandomBalance struct {
	peers []string
}

func New() *RandomBalance {
	return new(RandomBalance)
}

func (rb *RandomBalance) Add(keys ...string) {
	for _, key := range keys {
		rb.peers = append(rb.peers, key)
	}
}

func (rb *RandomBalance) Get() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(rb.peers) - 1)
	return rb.peers[random]
}
