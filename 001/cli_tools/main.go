package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// Config 保存命令行参数。
type Config struct {
	filePath   string
	operation  string
	outputPath string
}

func parseArgs() (Config, error) {
	filePath := flag.String("file", "", "输入文件路径")
	operation := flag.String("operation", "", "操作类型：count 或 upper")
	outputPath := flag.String("output", "", "输出文件路径，可选")
	flag.Parse()

	if strings.TrimSpace(*filePath) == "" {
		return Config{}, fmt.Errorf("必须输入 -file 参数")
	}
	if strings.TrimSpace(*operation) == "" {
		return Config{}, fmt.Errorf("必须输入 -operation 参数")
	}

	return Config{
		filePath:   *filePath,
		operation:  *operation,
		outputPath: *outputPath,
	}, nil
}

func readInputFile(filePath string) ([]byte, error) {
	if filePath == "" {
		return nil, fmt.Errorf("必须指定 -file 参数")
	}
	return os.ReadFile(filePath)
}

func countCharacters(data []byte) ([]byte, error) {
	count := utf8.RuneCount(data)
	return []byte(fmt.Sprintf("字符数: %d\n", count)), nil
}

func toUpper(data []byte) ([]byte, error) {
	return []byte(strings.ToUpper(string(data))), nil
}

// 根据 operation 分发到具体的处理函数。
func process(operation string, data []byte) ([]byte, error) {
	switch operation {
	case "count":
		return countCharacters(data)
	case "upper":
		return toUpper(data)
	default:
		return nil, fmt.Errorf("不支持的操作: %s", operation)
	}
}

func outputResult(outputPath string, data []byte) error {
	if outputPath == "" {
		_, err := os.Stdout.Write(data)
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

func main() {
	// 1. 解析参数
	config, err := parseArgs()
	if err != nil {
		fmt.Println("参数错误:", err)
		return
	}

	// 2. 读取输入文件
	data, err := readInputFile(config.filePath)
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	// 3. 判断 operation，并执行对应处理
	result, err := process(config.operation, data)
	if err != nil {
		fmt.Println("处理错误:", err)
		return
	}

	// 4. 输出到文件或终端
	if err := outputResult(config.outputPath, result); err != nil {
		fmt.Println("输出错误:", err)
	}
}
