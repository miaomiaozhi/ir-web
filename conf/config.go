package conf

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type IrConfig struct {
	root gjson.Result
	file string
}

type globalConf struct {
	Conf *IrConfig
	// others
}

var globalConfInfo *globalConf

// 读取全局配置文件数据
func GetConfig() *globalConf {
	if globalConfInfo == nil {
		panic("init global config must be called before GET")
	}
	return globalConfInfo
}

func Read(filePath string) (*IrConfig, error) {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s, err: %s", filePath, err)
	}

	conf, err := Parse(string(dat))

	if err != nil {
		return nil, errors.WithMessagef(err, "in %s", filePath)
	}
	conf.file = filePath
	log.Println("read config file path success")
	return conf, nil
}

func Parse(jsonData string) (*IrConfig, error) {
	if !gjson.Valid(jsonData) {
		return nil, fmt.Errorf("invalid config json")
	}
	res := gjson.Parse(jsonData)

	c := &IrConfig{
		root: res,
	}
	return c, nil
}

func InitGlobalConfig(conf *IrConfig) {
	globalConfInfo = &globalConf{
		Conf: conf,
	}
	log.Println("init global config success")
}

func (h *IrConfig) GetInt(path string, def int64) int64 {
	val := h.root.Get(path)
	if !val.Exists() {
		return def
	}
	return val.Int()
}

func (h *IrConfig) GetBool(path string, def bool) bool {
	val := h.root.Get(path)
	if !val.Exists() {
		return def
	}
	return val.Bool()
}

func (h *IrConfig) GetFloat(path string, def float64) float64 {
	val := h.root.Get(path)
	if !val.Exists() {
		return def
	}
	return val.Float()
}

func (h *IrConfig) GetString(path string, def string) string {
	val := h.root.Get(path)
	if !val.Exists() {
		return def
	}
	return val.String()
}

func (h *IrConfig) MustGetAny(path string) gjson.Result {
	val := h.root.Get(path)
	if !val.Exists() {
		panic(fmt.Sprintf("cannot get config in %s %s", h.file, path))
	}
	return val
}

func (h *IrConfig) MustGetInt(path string) int64 {
	val := h.root.Get(path)
	if !val.Exists() {
		panic(fmt.Sprintf("cannot get config in %s %s", h.file, path))
	}
	if val.Type != gjson.Number {
		panic(fmt.Sprintf("expect number type in %s %s", h.file, path))
	}
	return val.Int()
}

func (h *IrConfig) MustGetFloat(path string) float64 {
	val := h.root.Get(path)
	if !val.Exists() {
		panic(fmt.Sprintf("cannot get config in %s %s", h.file, path))
	}
	if val.Type != gjson.Number {
		panic(fmt.Sprintf("expect number type in %s %s", h.file, path))
	}
	return val.Float()
}

func (h *IrConfig) MustGetString(path string) string {
	val := h.root.Get(path)
	if !val.Exists() {
		panic(fmt.Sprintf("cannot get config in %s %s", h.file, path))
	}
	if val.Type != gjson.String {
		panic(fmt.Sprintf("expect string type in %s %s", h.file, path))
	}
	return val.String()
}
