package cache

import (
	"gps-worker/domain"
	i "gps-worker/infra"
)

func CachePositions(positions []domain.Position) error {
	cache := i.RedisCache{}
	cache.Connect()

	for _, p := range positions {
		err := cache.Set(i.CacheData{Key: string(p.Name), Value: string(rune(p.Distance))})
		if err != nil {
			return err
		}
	}

	return nil
}
