package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDownloadAllCancelsOthersWhenOneTimesOut(t *testing.T) {
	canceled := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/slow":
			select {
			case <-time.After(200 * time.Millisecond):
				_, _ = w.Write([]byte("slow"))
			case <-r.Context().Done():
				close(canceled)
			}
		case "/fast":
			_, _ = w.Write([]byte("fast"))
		}
	}))
	defer server.Close()

	results, err := downloadAll([]string{server.URL + "/slow", server.URL + "/fast"}, 30*time.Millisecond)
	if err == nil {
		t.Fatal("期望出现超时错误")
	}
	if len(results) != 2 {
		t.Fatalf("结果数量 = %d，期望为 2", len(results))
	}
	select {
	case <-canceled:
	case <-time.After(time.Second):
		t.Fatal("其他请求没有收到取消信号")
	}

	var timeoutFound bool
	for _, result := range results {
		if errors.Is(result.err, context.DeadlineExceeded) {
			timeoutFound = true
		}
	}
	if !timeoutFound {
		t.Fatal("结果中没有发现超时错误")
	}
}

func TestDownloadAllRejectsInvalidInput(t *testing.T) {
	if _, err := downloadAll(nil, time.Second); err == nil {
		t.Fatal("空 URL 列表应该报错")
	}
	if _, err := downloadAll([]string{"http://example.com"}, 0); err == nil {
		t.Fatal("非正超时时间应该报错")
	}
}
