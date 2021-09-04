package random

import (
	"math/rand"
	"time"
)

type rBalance struct {
	peers []string
}

func New() *rBalance {
	return new(rBalance)
}

func (rb *rBalance) Add(keys ...string) {
	for _, key := range keys {
		rb.peers = append(rb.peers, key)
	}
}

func (rb *rBalance) Get() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(rb.peers) - 1)
	return rb.peers[random]
}
