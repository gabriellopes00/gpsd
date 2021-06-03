package path

import (
	"fmt"
	d "gps-worker/domain"
	"time"
)

func GetPositionPath(entrypoint *d.Position, positions []*d.Position, paths chan<- *d.Position) {
	for _, p := range positions {
		go getPath(p, paths)
	}
	close(paths)
}

func getPath(position *d.Position, channel chan<- *d.Position) {
	time.Sleep(time.Millisecond * 1475)
	channel <- position
	fmt.Println(position.Name)
}
