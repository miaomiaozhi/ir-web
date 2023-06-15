package middleware

import (
	"ir-web/conf"
	mlog "ir-web/internal/wrapper/mlog"
)

type InvertedIndex map[string][]int

type Engine struct {
	Terms         map[int][]string
	Title         []string
	Urls          []string
	VocabularySet map[string]bool
	PostingList   InvertedIndex
	Documents     []string
	TfIdfMatrix   map[int]map[string]float64
}

var engine *Engine

func GetEngine() *Engine {
	if engine == nil {
		mlog.Error("init engine first")
		InitEngineWithConfig(conf.GetConfig().Conf)
		if engine == nil {
			panic("init engine error")
		}
		return engine
	}
	return engine
}
