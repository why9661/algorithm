package smooth

type WRoundRobin struct {
	peers []*node
}

type node struct {
	key             string
	weight          int
	curWeight       int
	effectiveWeight int
}

func New() *WRoundRobin {
	return new(WRoundRobin)
}

func (w *WRoundRobin) Add(key string, weight int) {
	n := &node{
		key:             key,
		weight:          weight,
		curWeight:       weight,
		effectiveWeight: weight,
	}

	w.peers = append(w.peers, n)
}

func (w *WRoundRobin) Get() string {
	totalWeight := 0

	var maxWeight *node

	for _, peer := range w.peers {
		totalWeight += peer.effectiveWeight
		peer.curWeight += peer.effectiveWeight
		if maxWeight == nil || peer.curWeight >= maxWeight.curWeight {
			maxWeight = peer
		}
	}

	maxWeight.curWeight -= totalWeight
	return maxWeight.key
}
