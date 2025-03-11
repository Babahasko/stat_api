package stat

import (
	"go/adv-demo/pkg/res"
	"net/http"
	"time"
)

const (
	GroupByDay = "day"
	GroupByMonth = "month"
)
type StatHandlerDeps struct {
	StatRepository *StatRepository
}
type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps *StatHandlerDeps ) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.HandleFunc("GET /stat", handler.GetStat())
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse(time.DateOnly, r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		to, err := time.Parse(time.DateOnly, r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by !=GroupByDay && by !=GroupByMonth {
			http.Error(w, "Invalid by parram", http.StatusBadRequest)
			return
		}
		stats := handler.StatRepository.GetStat(by, from, to)
		res.Json(w, stats, 200)
	}
}