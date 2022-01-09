package fs

import (
	"backendOneLessons/lesson4/internal/pkg/image"
	statpkg "backendOneLessons/lesson4/pkg/stat"
	"context"
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Stat struct {
	MethodDuration *statpkg.Timer
}

type fs struct {
	basePath string
	stat     Stat
}

func (f fs) Save(_ context.Context, filename string, data []byte) (string, error) {
	defer f.stat.MethodDuration.WithLabels(map[string]string{"method_name": "Save"}).Start().Stop()

	path := fmt.Sprintf("%s/%s", f.basePath, filename)
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err = file.Write(data); err != nil {
		return "", err
	}

	return path, nil
}

func New(path string, stat promauto.Factory) image.Storage {
	ret := fs{
		basePath: path,
		stat: Stat{MethodDuration: &statpkg.Timer{HistogramVec: stat.NewHistogramVec(
			prometheus.HistogramOpts{Name: "storage_method_duration"},
			[]string{"method_name"},
		)}},
	}

	prometheus.MustRegister(ret.stat.MethodDuration)

	return ret
}
