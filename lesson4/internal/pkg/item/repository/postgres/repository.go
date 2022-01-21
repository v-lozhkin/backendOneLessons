package postgres

import (
	"backendOneLessons/lesson4/internal/pkg/item"
	statpkg "backendOneLessons/lesson4/pkg/stat"

	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Stat struct {
	MethodDuration *statpkg.Timer
}

type repository struct {
	db   *sqlx.DB
	stat Stat
}

func New(db *sqlx.DB, stat promauto.Factory) item.Repository {
	ret := repository{
		db: db,
		stat: Stat{MethodDuration: &statpkg.Timer{HistogramVec: stat.NewHistogramVec(
			prometheus.HistogramOpts{Name: "repo_postgres_method_duration"},
			[]string{"method_name"},
		)}},
	}

	prometheus.MustRegister(ret.stat.MethodDuration)

	return ret
}
