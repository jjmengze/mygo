package signal

import (
	"context"
	"syscall"
	"testing"
	"time"
)

func TestSetupSignalContext(t *testing.T) {
	type args struct {
		fun func(ctx context.Context, t *testing.T)
	}
	tests := []struct {
		name      string
		situation args
		want      context.Context
	}{
		{
			name: "called twice",
			situation: args{
				fun: func(ctx context.Context, t *testing.T) {
					defer func() { recover() }()
					SetupSignalContext()
				},
			},
			want: nil,
		},
		{
			name: "cancel context",
			situation: args{
				fun: func(ctx context.Context, t *testing.T) {
					syscall.Kill(syscall.Getpid(), syscall.SIGINT)
					ctxTimeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
					defer cancel()
					select {
					case <-ctx.Done():
					case <-ctxTimeout.Done():
						t.Error("unexpected exit")
					}
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := SetupSignalContext()
			tt.situation.fun(ctx, t)
		})
	}
}
