package main

import (
	"context"
	"log"
	"sync"
)

type hub struct {
	// 互斥鎖
	sync.Mutex
	// key:指向 subscriber type 的 pointer, value: empty struct{}
	subs map[*subscriber]struct{}
}

// 發佈訊息
func (h *hub) publish(ctx context.Context, msg *message) error {
	log.Println("starting publish function, receiver hub 🅷 🅷 🅷")
	h.Lock()
	for s := range h.subs {
		s.publish(ctx, msg)
	}
	h.Unlock()
	return nil
}

// 將 subscriber 加進 hub 的訂閱列表, 並啟動 subscriber 的消息處理循環(run)
func (h *hub) subscribe(ctx context.Context, s *subscriber) error {
	log.Println("starting subscribe function, receiver hub 🅷 🅷 🅷")
	h.Lock()
	h.subs[s] = struct{}{} // 將 subscriber 加進名為 subs 的 map 當作 key
	h.Unlock()

	go func() {
		select {
		case <-s.quit:
		case <-ctx.Done():
			h.Lock()
			delete(h.subs, s) // 從 hub 的訂閱者清單中, 刪除此筆
			h.Unlock()
		}
	}()

	go s.run(ctx) // context: 上下文,控制 subscriber 的生命週期
	return nil    // nil 是訂閱成功情況下, 回傳的訊息
}

// 從 hub 中取消指定的訂閱者，如果操作成功返回nil，否則傳回錯誤訊息
func (h *hub) unsubscribe(ctx context.Context, s *subscriber) error {
	log.Println("starting 💔 unsubscribe function, receiver hub and sub: ", s)
	h.Lock()
	delete(h.subs, s)
	h.Unlock()
	close(s.quit)
	return nil
}

func (h *hub) Subscribers() int {
	h.Lock()
	c := len(h.subs) // 傳回這個hub有幾個subscriber的數量
	h.Unlock()
	return c
}

// 建立並返回一個新的 hub instance
func newHub() *hub {
	return &hub{
		subs: map[*subscriber]struct{}{},
	}
}
