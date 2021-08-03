package backoffmanager

import (
	"math"
	"math/rand"
	"time"
)

type BackoffManager interface {
	Backoff() *time.Timer
}

type Backoff struct {
	Duration time.Duration
	Factor   float64
	Jitter   float64
	Steps    int
	Cap      time.Duration
}

type exponentialBackoffManager struct {
	backoff              *Backoff
	backoffTimer         *time.Timer
	lastBackoffStart     time.Time
	initialBackoff       time.Duration
	backoffResetDuration time.Duration
	clock                time.Time
	nowFunc              func() time.Time
}

// NewExponentialBackoffManager 需要設置最初始延遲時間initBackoff，最大延遲時間maxBackoff。
// 若是 backoff 時間超過，最大延遲時間 maxBackoff 之後的 backoff 時間只會是加上抖動 jitter。
// resetDuration 表示，捼果 backoff 時間超過所所設定重置時間表示 backoff 工作做太久需要重置，backoff 會退回 initBackoff。
// backoffFactor 倍數因子，每次 backoff 需要增量的倍數。
func NewExponentialBackoffManager(initBackoff, maxBackoff, resetDuration time.Duration, backoffFactor, jitter float64, nowFunc func() time.Time) BackoffManager {
	return &exponentialBackoffManager{
		backoff: &Backoff{
			Duration: initBackoff,
			Factor:   backoffFactor,
			Jitter:   jitter,
			Steps:    math.MaxInt32,
			Cap:      maxBackoff,
		},
		backoffTimer:         nil,
		initialBackoff:       initBackoff,
		lastBackoffStart:     nowFunc(),
		backoffResetDuration: resetDuration,
		clock:                nowFunc(),
		nowFunc:              nowFunc,
	}
}

func (b *exponentialBackoffManager) getNextBackoff() time.Duration {
	b.clock = b.nowFunc()
	if b.clock.Sub(b.lastBackoffStart) > b.backoffResetDuration {
		b.backoff.Steps = math.MaxInt32
		b.backoff.Duration = b.initialBackoff
	}
	b.lastBackoffStart = time.Now()
	return b.backoff.Step()
}

func (b *exponentialBackoffManager) Backoff() *time.Timer {
	if b.backoffTimer == nil {
		b.backoffTimer = time.NewTimer(b.getNextBackoff())
	} else {
		b.backoffTimer.Reset(b.getNextBackoff())
	}
	return b.backoffTimer
}

func (b *Backoff) Step() time.Duration {
	if b.Steps < 1 {
		if b.Jitter > 0 {
			return Jitter(b.Duration, b.Jitter)
		}
		return b.Duration
	}
	b.Steps--

	duration := b.Duration

	// calculate the next step
	if b.Factor != 0 {
		b.Duration = time.Duration(float64(b.Duration) * b.Factor)
		if b.Cap > 0 && b.Duration > b.Cap {
			b.Duration = b.Cap
			b.Steps = 0
		}
	}

	if b.Jitter > 0 {
		duration = Jitter(duration, b.Jitter)
	}
	return duration
}

func Jitter(duration time.Duration, maxFactor float64) time.Duration {
	if maxFactor <= 0.0 {
		maxFactor = 1.0
	}
	wait := duration + time.Duration(rand.Float64()*maxFactor*float64(duration))
	return wait
}
