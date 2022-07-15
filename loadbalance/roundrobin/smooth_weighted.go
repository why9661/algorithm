package roundrobin

type SmoothWeightedRoundRobin struct {
	peers       []*node
	totalWeight int
}

type node struct {
	key             string
	weight          int // configure
	curWeight       int
	effectiveWeight int
}

func NewSmoothWeightedRoundRobin() *SmoothWeightedRoundRobin {
	return new(SmoothWeightedRoundRobin)
}

func (w *SmoothWeightedRoundRobin) Add(key string, weight int) {
	n := &node{
		key:             key,
		weight:          weight,
		curWeight:       weight, // or 0
		effectiveWeight: weight,
	}

	w.totalWeight += weight

	w.peers = append(w.peers, n)
}

func (w *SmoothWeightedRoundRobin) Get() string {
	var maxWeight *node

	for _, peer := range w.peers {
		peer.curWeight += peer.effectiveWeight
		if maxWeight == nil || peer.curWeight >= maxWeight.curWeight {
			maxWeight = peer
		}
	}

	maxWeight.curWeight -= w.totalWeight
	return maxWeight.key
}
