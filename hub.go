package main

import (
	"context"
	"sync"
)

type hub struct {
	// 互斥鎖
	sync.Mutex
	// key:指向 subscriber type 的 pointer, value: empty struct{}
	subs map[*subscriber]struct{}
}

// 將 subscriber 加進 hub 的訂閱列表, 並啟動 subscriber 的消息處理循環(run)
func (h *hub) subscribe(ctx context.Context, s *subscriber) error {
	h.Lock()
	h.subs[s] = struct{}{} // 將 subscriber 加進名為 subs 的 map 當作 key
	h.Unlock()

	go s.run(ctx) // context: 上下文,控制 subscriber 的生命週期
	return nil    // nil 是訂閱成功情況下, 回傳的訊息
}

// 建立並返回一個新的 hub instance
func newHub() *hub {
	return &hub{
		subs: map[*subscriber]struct{}{},
	}
}
