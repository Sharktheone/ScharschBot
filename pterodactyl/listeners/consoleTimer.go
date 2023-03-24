package listeners

import "time"

type consoleTimer struct {
	enabled bool
	t       *time.Timer
	started bool
	maxTime time.Duration
	c       <-chan time.Time
}

func newTimer(maxTime string) (*consoleTimer, error) {
	if maxTime == "" {
		return &consoleTimer{
			enabled: false,
		}, nil
	}
	m, err := time.ParseDuration(maxTime)
	if err != nil {
		return nil, err
	}
	return &consoleTimer{
		maxTime: m,
	}, nil
}

func (t *consoleTimer) start() {
	if !t.enabled {
		return
	}
	if t.started {
		return
	}
	t.t = time.NewTimer(t.maxTime)
	t.c = t.t.C
	t.started = true
}
func (t *consoleTimer) reset() {
	if !t.enabled {
		return
	}
	t.t.Reset(t.maxTime)
}
func (t *consoleTimer) stop() {
	if !t.enabled {
		return
	}
	t.t.Stop()
	t.started = false
}
