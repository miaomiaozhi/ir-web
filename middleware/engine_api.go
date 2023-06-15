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
		Terms:         make(map[int][]string),
		Title:         make([]string, 0),
		Urls:          make([]string, 0),
		VocabularySet: make(map[string]bool),
		PostingList:   make(InvertedIndex),
		Documents:     make([]string, 0),
		TfIdfMatrix:   make(map[int]map[string]float64),
	}
	readFile(conf, engine)
	build(conf, engine)
	calcTF_IDF(conf, engine)

	mlog.Info("init engine with config success")
}

func prepross(doc string) string {
	return doc
}

func readFile(conf *conf.IrConfig, engine *Engine) {
	dataPath := conf.MustGetString("data.path")
	mlog.Info("data path is", dataPath)
	mlog.Info("reading data files")

	cnt := 0
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
		doc := prepross(string(content))
		if doc == "" {
			return nil
		}

		engine.Documents = append(engine.Documents, doc)
		url, title, body := pkg.SplitDocument(doc)

		if url == "" || title == "" || body == "" {
			return nil
		}

		if conf.GetString("debug", "false") == "true" {
			if cnt < 5 {
				mlog.Info("url, title", url, title)
			} else if cnt == 5 {
				mlog.Info("log only 5 pages")
			}
			cnt += 1
		}

		engine.Urls = append(engine.Urls, url)
		engine.Title = append(engine.Title, title)

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
