package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type downloadResult struct {
	index int
	url   string
	data  []byte
	err   error
}

// downloadOne 下载一个文件。它同时受 parent 和 timeout 控制。
func downloadOne(parent context.Context, url string, timeout time.Duration) ([]byte, error) {
	ctx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("HTTP 状态码异常: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}
	return data, nil
}

// downloadAll 并发下载所有 URL；任意一个任务超时，就取消其他任务。
func downloadAll(urls []string, timeout time.Duration) ([]downloadResult, error) {
	if len(urls) == 0 {
		return nil, errors.New("至少需要提供一个下载地址")
	}
	if timeout <= 0 {
		return nil, errors.New("超时时间必须大于 0")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make(chan downloadResult, len(urls))
	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func(index int, url string) {
			defer wg.Done()
			data, err := downloadOne(ctx, url, timeout)
			if err != nil && errors.Is(err, context.DeadlineExceeded) {
				cancel() // 一个任务超时，取消所有其他任务。
			}
			results <- downloadResult{index: index, url: url, data: data, err: err}
		}(i, url)
	}

	wg.Wait()
	close(results)

	allResults := make([]downloadResult, 0, len(urls))
	var failed int
	for result := range results {
		allResults = append(allResults, result)
		if result.err != nil {
			failed++
		}
	}
	if failed > 0 {
		return allResults, fmt.Errorf("%d/%d 个下载任务失败", failed, len(urls))
	}
	return allResults, nil
}

func saveResults(results []downloadResult, dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("创建输出目录失败: %w", err)
	}

	for _, result := range results {
		if result.err != nil {
			continue
		}
		name := filepath.Base(strings.TrimSuffix(result.url, "/"))
		if name == "." || name == "" || name == string(filepath.Separator) {
			name = fmt.Sprintf("download-%d", result.index+1)
		}
		path := filepath.Join(dir, fmt.Sprintf("%02d-%s", result.index+1, name))
		if err := os.WriteFile(path, result.data, 0o644); err != nil {
			return fmt.Errorf("保存 %s 失败: %w", result.url, err)
		}
		fmt.Printf("下载成功: %s -> %s\n", result.url, path)
	}
	return nil
}

func main() {
	timeout := flag.Duration("timeout", 5*time.Second, "每个下载任务的超时时间")
	outDir := flag.String("out", "downloads", "文件保存目录")
	flag.Parse()

	results, err := downloadAll(flag.Args(), *timeout)
	if saveErr := saveResults(results, *outDir); saveErr != nil {
		fmt.Fprintln(os.Stderr, saveErr)
		os.Exit(1)
	}
	if err != nil {
		for _, result := range results {
			if result.err != nil {
				fmt.Fprintf(os.Stderr, "下载失败: %s: %v\n", result.url, result.err)
			}
		}
		os.Exit(1)
	}
}
