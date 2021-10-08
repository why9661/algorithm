package roundrobin

type RoundRobin struct {
	curIndex int
	peers    []string
}

func New() *RoundRobin {
	return new(RoundRobin)
}

func (rrb *RoundRobin) Add(keys ...string) {
	for _, key := range keys {
		rrb.peers = append(rrb.peers, key)
	}
}

func (rrb *RoundRobin) Get() string {
	rrb.curIndex++
	rrb.curIndex = rrb.curIndex % len(rrb.peers)
	return rrb.peers[rrb.curIndex]
}
