package random

import (
	"math/rand"
	"time"
)

type RandBalance struct {
	peers []string
}

func New() *RandBalance {
	return new(RandBalance)
}

func (rb *RandBalance) Add(keys ...string) {
	for _, key := range keys {
		rb.peers = append(rb.peers, key)
	}
}

func (rb *RandBalance) Get() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(rb.peers) - 1)
	return rb.peers[index]
}
