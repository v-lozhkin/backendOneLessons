package stat

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Timer struct {
	*prometheus.HistogramVec
	startTime time.Time
	labels    prometheus.Labels
}

func (t *Timer) WithLabels(labels prometheus.Labels) *Timer {
	t.labels = labels

	return t
}

func (t *Timer) Start() *Timer {
	t.startTime = time.Now()

	return t
}

func (t *Timer) Stop() {
	t.With(t.labels).Observe(time.Since(t.startTime).Seconds())
}
