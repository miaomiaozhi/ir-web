package runner

import (
	conf "ir-web/conf"
	mlog "ir-web/internal/wrapper/mlog"
	"net/http"
	"strconv"
	"time"
)

func InitBackgournTaskWithConfig(conf *conf.IrConfig) {
	// 指定 Health URL
	healthURL := "http://" + conf.GetString("web.host", "127.0.0.1") +
		":" + strconv.FormatInt(conf.GetInt("web.port", 8080), 10) + "/api/v1/health"
	mlog.Info("health url is", healthURL)

	// 每隔 5s 发送心跳请求
	ticker := time.NewTicker(time.Second * time.Duration(conf.GetInt("web.heartbeat", 10)))
	defer ticker.Stop()

	for range ticker.C {
		res, err := http.Get(healthURL)
		if err != nil {
			mlog.Error("heart beat error:", err)
			continue
		}
		mlog.Info("heart beat success")
		res.Body.Close()
	}
}
