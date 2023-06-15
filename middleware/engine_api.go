package middleware

import (
	"ir-web/conf"
	mlog "ir-web/internal/wrapper/mlog"
	"ir-web/pkg"
	"math"
	"os"
	"path/filepath"
	"regexp"

	"github.com/kljensen/snowball"
	"github.com/yanyiwu/gojieba"
)

func InitEngineWithConfig(conf *conf.IrConfig) {
	engine = &Engine{
		Terms:                  make(map[int][]string),
		CosineSimilarityMatrix: make(map[int]map[int]float64),
		VocabularySet:          make(map[string]bool),
		PostingList:            make(InvertedIndex),
		Documents:              make([]string, 0),
		TfIdfMatrix:            make(map[int]map[string]float64),
	}
	readFile(conf, engine)
	build(conf, engine)
	calcTF_IDF(conf, engine)
}

func prepross(doc string) string {
	return doc
}

func readFile(conf *conf.IrConfig, engine *Engine) {
	dataPath := conf.MustGetString("data.path")
	mlog.Info("data path is", dataPath)
	if err := filepath.Walk(dataPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !info.Mode().IsRegular() || filepath.Ext(path) != ".txt" {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		engine.Documents = append(engine.Documents, prepross(string(content)))
		return nil
	}); err != nil {
		mlog.Fatal("read data error", err)
	}
	mlog.Info("read data file success")
}

func build(conf *conf.IrConfig, engine *Engine) {
	jieba := gojieba.NewJieba()
	for index, doc := range engine.Documents {
		words := jieba.CutForSearch(doc, true)
		english := regexp.MustCompile(`\b\w+\b`).FindAllString(doc, -1)

		for _, word := range words {
			if regexp.MustCompile(`^[\p{Han}]+$`).MatchString(word) {
				engine.Terms[index] = append(engine.Terms[index], word)
				engine.PostingList[word] = append(engine.PostingList[word], index)
				engine.VocabularySet[word] = true
			}
		}
		for _, word := range english {
			word, _ = snowball.Stem(word, "english", true)
			engine.Terms[index] = append(engine.Terms[index], word)
			engine.PostingList[word] = append(engine.PostingList[word], index)
			engine.VocabularySet[word] = true
		}
	}
	mlog.Info("build intverted index success")
}

func calcTF_IDF(conf *conf.IrConfig, engine *Engine) {
	// calc IDF
	IDF := make(map[string]float64)
	for vocab := range engine.VocabularySet {
		IDF[vocab] = math.Log10(float64(len(engine.Documents)) /
			float64(pkg.GetDF(engine.PostingList[vocab])))
	}
	for id := range engine.Documents {
		engine.TfIdfMatrix[id] = make(map[string]float64)
		for _, term := range engine.Terms[id] {
			val := pkg.GetTF(id, engine.PostingList[term])
			if val == 0 {
				engine.TfIdfMatrix[id][term] = 0
			} else {
				engine.TfIdfMatrix[id][term] = float64(val) * IDF[term]
			}
		}
	}
	mlog.Info("calc TF-IDF success")
}

