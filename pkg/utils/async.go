package utils

import (
	"context"
	"time"
)

func DoWithTimeout(call func(), timeout time.Duration) {
	// 创建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 使用goroutine调用传入的函数
	done := make(chan any)
	go func() {
		call()
		done <- struct{}{}
	}()

	// 等待结果或超时
	select {
	case <-done:
		return
	case <-ctx.Done():
		return
	}
}
