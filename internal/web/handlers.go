package web

import (
	"fmt"

	"github.com/DumbNoxx/goxe/internal/dashboard"
)

type Handler struct {
	Data *dashboard.Metrics
}

// TODO: save metrics to a file, then read the file to compare metrics
func (h *Handler) Hola() {
	fmt.Println(h.Data.GetMetrics())
}


func NewHandler(m *dashboard.Metrics) Handler {
	return Handler{
		Data: m,
	}
}

