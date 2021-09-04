package roundrobin

type roundRobin struct {
	curIndex int
	peers    []string
}

func New() *roundRobin {
	return new(roundRobin)
}

func (rrb *roundRobin) Add(keys ...string) {
	for _, key := range keys {
		rrb.peers = append(rrb.peers, key)
	}
}

func (rrb *roundRobin) Get() string {
	rrb.curIndex++
	rrb.curIndex = rrb.curIndex % len(rrb.peers)
	return rrb.peers[rrb.curIndex]
}
