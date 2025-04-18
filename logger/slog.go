// Package logger 框架自用日志包
package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

// Config 日志配置结构体
type Config struct {
	Level  slog.Leveler // 日志等级，如 slog.LevelInfo
	Format string       // 输出格式: "json" or "text"
}

// Init 初始化 logger
func Init(cfg Config) {
	var handler slog.Handler

	switch cfg.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: cfg.Level,
		})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: cfg.Level,
		})
	}

	Logger = slog.New(handler)
	slog.SetDefault(Logger) // 设置为全局默认 logger
}
