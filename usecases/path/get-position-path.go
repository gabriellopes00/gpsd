package path

import (
	"gps-worker/domain"
	"time"
)

func GetPositionPath(entrypoint domain.Position, channel chan domain.Position) {
	for p := range channel {
		time.Sleep(time.Millisecond * 2)
		channel <- p
	}
}
