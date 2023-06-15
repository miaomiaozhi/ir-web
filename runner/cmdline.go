package runner

import (
	"flag"
	conf "ir-web/conf"
	mlog "ir-web/internal/wrapper/mlog"
	"os"
)

func MustReadConfigFromCmdFlags() (*conf.IrConfig, error) {
	var configFilePath string

	// 将命令行参数绑定到configFilePath变量上
	flag.StringVar(&configFilePath, "config", "", "path to config file")
	flag.Parse()

	// 如果未指定config参数，则输出提示信息
	mlog.Info("config file path", configFilePath)
	if configFilePath == "" {
		mlog.Fatal("config file path is empty")
	}
	if pathChecker(configFilePath) {
		return conf.Read(configFilePath)
	} else {
		mlog.Fatal("illegal config file path")
	}
	// unreachable
	return nil, nil
}

func pathChecker(path string) bool {
	mlog.Info("path checker", path)
	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		mlog.Fatalf("File %s does not exist\n", path)
		return false
	}

	// 检查文件权限
	if _, err := os.OpenFile(path, os.O_RDONLY, 0666); err != nil {
		mlog.Fatalf("File %s is not readable\n", path)
		return false
	}

	// 检查文件类型
	fileInfo, err := os.Stat(path)
	if err != nil {
		mlog.Fatalf("Failed to get file info for %s: %v\n", path, err)
		return false
	}

	if fileInfo.IsDir() {
		mlog.Fatalf("%s is not a file\n", path)
		return false
	}

	mlog.Info("config path is legal", path)
	return true
}
