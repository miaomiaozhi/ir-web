package middleware

import (
	"container/heap"
	"ir-web/conf"
	"ir-web/pkg"
	"math"
)

// 传入查询 token 返回查找结果
func (e *Engine) QueryIndexListByToken(token string) []int {
	words := pkg.SplitWorkByLanguage(token)
	return e.queryIndexListHelper(token, words)
}

// 传入查询 token 返回模糊查找结果
func (e *Engine) FuzzyQueryIndexListByToken(token string) []int {
	prev := pkg.SplitWorkByLanguage(token)
	words := make([]string, 0, len(prev))
	words = append(words, prev...)

	for term := range e.VocabularySet {
		for _, prevWord := range prev {
			if pkg.GetMinimalEditDistance(term, prevWord) <= int(conf.GetConfig().Conf.GetInt("engine.distance", 1)) {
				words = append(words, term)
			}
		}
	}
	return e.queryIndexListHelper(token, words)
}

// 传入 words 然后返回 index 列表
func (e *Engine) queryIndexListHelper(token string, words []string) []int {
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
	tf := calc(token, terms)
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

func calc(query string, queryWords []string) map[string]float64 {
	tf := make(map[string]float64)
	for _, word := range queryWords {
		tf[word]++
	}
	for index := range tf {
		tf[index] /= float64(len(queryWords))
	}
	return tf
}
