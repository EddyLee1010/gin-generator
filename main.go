package main

import (
	"github.com/eddylee1010/gin-generator/cmd"
	"github.com/eddylee1010/gin-generator/logger"
	"log/slog"
)

func main() {

	// 初始化日志系统
	logger.Init(logger.Config{
		Level:  slog.LevelDebug, // 输出 Debug 及以上日志
		Format: "json",          // 或 "json"
	})
	cmd.Execute()
}
