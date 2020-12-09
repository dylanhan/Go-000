package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	exit := make(chan os.Signal)
	g.Go(func() error {
		if err := http.ListenAndServe("127.0.0.0:8000", nil); err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		signal.Notify(exit)
		select {
		case <-exit:
			return errors.New("receive signal, exit")
		case <-ctx.Done():
			fmt.Println("signal all cancel")
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
