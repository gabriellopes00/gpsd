package path

import (
	d "gps-worker/domain"
	"time"
)

func GetPositionPath(entrypoint *d.Position, positions []*d.Position, paths chan<- *d.Position) {
	for _, p := range positions {
		time.Sleep(time.Millisecond * 1475)
		paths <- p
	}

	close(paths)
}
