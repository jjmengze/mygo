package backoffmanager

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBackoff_Step(t *testing.T) {
	tests := []struct {
		initial *Backoff
		want    []time.Duration
	}{
		{
			initial: &Backoff{
				Duration: time.Second,
				Factor:   2,
				Steps:    3,
				Cap:      3 * time.Second,
			},
			want: []time.Duration{
				1 * time.Second,
				2 * time.Second,
				3 * time.Second,
			},
		},
	}
	for seed := int64(0); seed < 5; seed++ {
		for _, tt := range tests {
			initial := *tt.initial
			t.Run(fmt.Sprintf("%#v seed=%d", initial, seed), func(t *testing.T) {
				rand.Seed(seed)
				for i := 0; i < len(tt.want); i++ {
					got := initial.Step()
					t.Logf("[%d]=%s", i, got)
					if initial.Jitter > 0 {
						if got == tt.want[i] {
							t.Errorf("Backoff.Step(%d) = %v, no jitter", i, got)
							continue
						}
						diff := float64(tt.want[i]-got) / float64(tt.want[i])
						if diff > initial.Jitter {
							t.Errorf("Backoff.Step(%d) = %v, want %v, outside range", i, got, tt.want)
							continue
						}
					} else {
						if got != tt.want[i] {
							t.Errorf("Backoff.Step(%d) = %v, want %v", i, got, tt.want)
							continue
						}
					}
				}
			})
		}
	}
}
