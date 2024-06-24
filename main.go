package main

import (
	"context"
	"time"
)

func main() {
	// var ctx context.Context = context.Background()
	ctx := context.Background()
	h := newHub()
	sub01 := newSubscriber("sub001")
	sub02 := newSubscriber("sub002")
	sub03 := newSubscriber("sub003")

	h.subscribe(ctx, sub01)
	h.subscribe(ctx, sub02)
	h.subscribe(ctx, sub03)

	_ = h.publish(ctx, &message{data: []byte("test01")})
	_ = h.publish(ctx, &message{data: []byte("test02")})
	_ = h.publish(ctx, &message{data: []byte("test03")})
	time.Sleep(1 * time.Second)

	h.unsubscribe(ctx, sub03)
	_ = h.publish(ctx, &message{data: []byte("test04")})
	_ = h.publish(ctx, &message{data: []byte("test05")})

	time.Sleep(1 * time.Second)
}
