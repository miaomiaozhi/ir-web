package middleware

import (
	"container/heap"
	"ir-web/pkg"
	"math"
)

func (e *Engine) QueryIndexListByToken(token string) []int {
	words := pkg.SplitWorkByLanguage(token)

	list := make([]int, 0)
	for _, word := range words {
		if ids, has := e.PostingList[word]; has {
			if len(list) == 0 {
				list = ids
			} else {
				list = pkg.Intersect(list, ids)
			}
		}
	}

	used := make(map[int]bool, 0)
	for _, id := range list {
		used[id] = true
	}

	cosineSimilarity := e.GetCosineSimlarity(token, words)

	infos := make([]*pkg.Info, 0)
	for id := range used {
		infos = append(infos, &pkg.Info{
			Key: id,
			Val: cosineSimilarity[id],
		})
	}

	mheap := pkg.MyHeap(infos)

	heap.Init(&mheap)
	topK := mheap.GetTopK(10)
	res := make([]int, 0)
	for _, v := range topK {
		res = append(res, int(v.(*pkg.Info).Key))
	}
	return res
}

func (e *Engine) GetCosineSimlarity(token string, terms []string) map[int]float64 {
	res := make(map[int]float64)
	tf := TF_IDF_ForQuery(token, terms)
	for i := range e.Documents {
		n1, n2 := e.TfIdfMatrix[i], tf
		n1Value, n2Value := 0.0, 0.0
		cosine := 0.0
		for word := range n2 {
			if _, ok := n1[word]; ok {
				cosine += n1[word] * n2[word]
				n1Value += n1[word] * n1[word]
				n2Value += n2[word] * n2[word]
			}
		}
		value := math.Sqrt(n1Value) * math.Sqrt(n2Value)
		if value < 1e-3 {
			res[i] = 0
		} else {
			res[i] = float64(cosine) / value
		}
	}
	return res
}

// TODO
func TF_IDF_ForQuery(query string, queryWords []string) map[string]float64 {
	// 计算Query的TF-IDF
	// query_idf := make(map[string]float64)
	query_tf := make(map[string]float64)
	for _, word := range queryWords {
		query_tf[word]++
	}
	for index := range query_tf {
		query_tf[index] /= float64(len(queryWords))
	}
	return query_tf
}
