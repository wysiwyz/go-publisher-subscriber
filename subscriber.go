package main

import (
	"context"
	"log"
	"sync"
)

// message 結構: 用來儲存 subscriber 的 message data
type message struct {
	data []byte
}

// subscriber 結構: 訂閱者具備的性質
type subscriber struct {
	// 互斥鎖
	sync.Mutex

	// 訂閱者名稱
	name string

	// 接收 message 的通道
	handler chan *message

	// 關閉這個 subscriber 的信息通道，型態是空結構體 struct{} (是不攜帶任何訊息的信號通道)
	quit chan struct{}
}

// 啟動 subscriber 的消息處理循環
func (s *subscriber) run(ctx context.Context) {
	for {
		select {
		// T <- ch operator用於從 channel 中接收數據
		// ch <- value 向 channel 發送數據
		case msg := <-s.handler:
			log.Println("[subscriber run]", s.name, string(msg.data))
		case <-s.quit:
			return
		case <-ctx.Done():
			return
		}
	}
}

func (s *subscriber) publish(ctx context.Context, msg *message) {
	log.Println("starting publish function, receiver subscriber 🆂 🆂 🆂")
	select {
	case <-ctx.Done():
		return
	case s.handler <- msg:
	default:
	}
}

// 建立一個新的訂閱者
func newSubscriber(name string) *subscriber {
	return &subscriber{
		name:    name,
		handler: make(chan *message, 100), // 建立一個容量為 100 的 message channel
		quit:    make(chan struct{}),      // 建立一個用來退出信號的 empty struct channel
	}
}
