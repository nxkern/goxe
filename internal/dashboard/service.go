package dashboard

import (
	"github.com/DumbNoxx/goxe/internal/utils"
)

type Metrics struct {
	Origin string
	Data DataLog
}
type DataLog struct {
	Level map[string]int
}

func (m *Metrics) SetMetrics(Data utils.Logs) {
	var count = make(map[string]int)
	for origin, messages := range Data{
		m.Origin = origin
		for _, stats := range messages {
			count[string(stats.Level)]++
		}
	m.Data.Level = count
	}
}

func (m *Metrics) GetMetrics() Metrics {
	return *m
}

