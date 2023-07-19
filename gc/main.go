package main

import (
	"context"
	"sync"
)

type Node struct {
}

func main() {
	root := make(chan Node)
	ctx := context.Background()
	expected := []string{"A", "B", "C", "D", "F", "H"}
	nodes := make([]Node, len(expected))
	for i := range expected {
		nodes[i] = Node{}
	}

	go func() {
		for _, n := range nodes {
			select {
			case root <- n:
			case <-ctx.Done():
				return
			}
		}
		close(root)
	}()
	ConcurrentMark(ctx, root, func(ctx context.Context, ref Node, fn func(Node)) {
		fn(ref)
	})

}

func ConcurrentMark(ctx context.Context, root <-chan Node, refs func(context.Context, Node, func(Node))) (map[Node]struct{}, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var (
		grays = make(chan Node)
		seen  = map[Node]struct{}{} // or not "white", basically "seen"
		wg    sync.WaitGroup

		refErr error
	)

	go func() {
		for gray := range grays {
			if _, ok := seen[gray]; ok {
				wg.Done()
				continue
			}
			seen[gray] = struct{}{} // post-mark this as non-white

			go func(gray Node) {
				defer wg.Done()

				send := func(n Node) {
					wg.Add(1)
					select {
					case grays <- n:
					case <-ctx.Done():
						wg.Done()
					}
				}

				refs(ctx, gray, send)

			}(gray)
		}
	}()

	for r := range root {
		wg.Add(1)
		select {
		case grays <- r:
		case <-ctx.Done():
			wg.Done()
		}

	}

	// Wait for outstanding grays to be processed
	wg.Wait()

	close(grays)

	if refErr != nil {
		return nil, refErr
	}
	if cErr := ctx.Err(); cErr != nil {
		return nil, cErr
	}

	return seen, nil
}
