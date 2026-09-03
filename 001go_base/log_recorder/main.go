package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// LogLevel 表示日志级别。
type LogLevel string

const (
	InfoLevel    LogLevel = "INFO"
	WarningLevel LogLevel = "WARNING"
	ErrorLevel   LogLevel = "ERROR"
)

// Config 保存命令行配置。
type Config struct {
	Level      LogLevel
	OutputPath string
}

// Logger 是日志记录器。
type Logger struct {
	Level  LogLevel
	Writer io.Writer
	file   *os.File
}

func parseArgs() (Config, error) {
	level := flag.String("level", "", "日志级别：INFO、WARNING 或 ERROR")
	outputPath := flag.String("output", "", "日志输出文件路径，不填写则输出到控制台")
	flag.Parse()

	parsedLevel := LogLevel(strings.ToUpper(strings.TrimSpace(*level)))
	if parsedLevel == "" {
		return Config{}, nil
	}

	switch parsedLevel {
	case InfoLevel, WarningLevel, ErrorLevel:
		return Config{
			Level:      parsedLevel,
			OutputPath: *outputPath,
		}, nil
	default:
		return Config{}, fmt.Errorf("不支持的日志级别: %s", *level)
	}
}

func newLogger(config Config) (*Logger, error) {
	if config.Level == "" {
		return &Logger{}, nil
	}

	logger := &Logger{
		Level:  config.Level,
		Writer: os.Stdout,
	}

	if config.OutputPath == "" {
		return logger, nil
	}

	file, err := os.OpenFile(
		config.OutputPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, fmt.Errorf("打开日志文件失败: %w", err)
	}

	logger.Writer = file
	logger.file = file
	return logger, nil
}

func (logger *Logger) log(level LogLevel, message string) {
	if logger == nil || logger.Level == "" || logger.Writer == nil {
		return
	}

	priority := func(current LogLevel) int {
		switch current {
		case InfoLevel:
			return 1
		case WarningLevel:
			return 2
		case ErrorLevel:
			return 3
		default:
			return 0
		}
	}

	// 只输出大于等于当前配置级别的日志。
	if priority(level) < priority(logger.Level) {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(logger.Writer, "%s [%s] %s\n", timestamp, level, message)
}

func (logger *Logger) Info(message string) {
	logger.log(InfoLevel, message)
}

func (logger *Logger) Warning(message string) {
	logger.log(WarningLevel, message)
}

func (logger *Logger) Error(message string) {
	logger.log(ErrorLevel, message)
}

func (logger *Logger) Close() error {
	if logger == nil || logger.file == nil {
		return nil
	}

	if err := logger.file.Close(); err != nil {
		return fmt.Errorf("关闭日志文件失败: %w", err)
	}
	return nil
}

func main() {
	config, err := parseArgs()
	if err != nil {
		fmt.Println("参数错误:", err)
		return
	}
	logger, err := newLogger(config)
	if err != nil {
		fmt.Println("创建日志记录器失败:", err)
		return
	}
	defer logger.Close()

	logger.Info("程序启动")
	logger.Warning("这是一条警告信息")
	logger.Error("这是一条错误信息")
}
