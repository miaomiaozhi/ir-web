package pkg

// Info 表示一个双关键字的键值对
type Info struct {
	Key int
	Val float64
}

// MyHeap 表示一个 Info 的堆
type MyHeap []*Info

func (h MyHeap) Len() int {
	return len(h)
}

func (h MyHeap) Less(i, j int) bool {
	return h[i].Val > h[j].Val
}

func (h MyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyHeap) Push(x interface{}) {
	*h = append(*h, x.(*Info))
}

func (h *MyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MyHeap) GetTopK(k int) []interface{} {
	res := make([]interface{}, 0, k)
	for h.Len() > 0 && k > 0 {
		res = append(res, h.Pop())
		k -= 1
	}
	return res
}
