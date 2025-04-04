package test

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"time"
)

func ExitWithCancelAfter(ctx context.Context, duration time.Duration) context.CancelFunc {
	ctx, cancelFn := context.WithCancel(ctx)

	go func() {
		select {
		case <-time.After(duration):
			slog.Error("timeout exceeded, terminating program.")
			os.Exit(1)
		case <-ctx.Done():
			return
		}
	}()

	return cancelFn
}

func ExpectPanic(t *testing.T) {
	if err := recover(); err == nil {
		t.Fatal("Expected a panic!")
	}
}

func ExpectNoPanic(t *testing.T) {
	if err := recover(); err != nil {
		t.Fatal("Don’t panic and always carry a towel!", "error:", err)
	}
}
