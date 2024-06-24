package main

import "context"

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
}
