package main

import (
	"fmt"
	"gps-worker/app/mock"
	d "gps-worker/domain"
	"gps-worker/infra"
	"log"
	"strconv"
	"strings"

	c "gps-worker/usecases/calc"
	m "gps-worker/usecases/mail"
	p "gps-worker/usecases/path"
	s "gps-worker/usecases/sort"
	v "gps-worker/usecases/validation"

	"github.com/go-redis/redis/v8"
)

func main() {
	entrypoint := &d.Position{Name: "Victim", Latitude: -23.16860649763682, Longitude: -46.86898723418065}
	positions := mock.Positions()

	cache := infra.RedisCache{Client: redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})}

	cache.Set(infra.CacheData{
		Key:   positions[0].Name,
		Value: fmt.Sprintf("%v %v", positions[0].Latitude, positions[0].Longitude),
	})

	if err := v.ValidateCoordinates(positions); err != nil {
		log.Fatalln(err)
	}

	paths := make(chan *d.Position)

	inRadius := c.GetInRadius(entrypoint, positions)
	c.GetDistances(entrypoint, inRadius)
	ordered := s.Sort(inRadius)
	go p.GetPositionPath(entrypoint, s.Filter(ordered), paths)
	m.SendHelperMail(paths)

	data, err := cache.Get(positions[0].Name)
	if err != nil {
		log.Fatalln(err)
	}

	lat, _ := strconv.ParseFloat(strings.Split(data.Value, " ")[0], 64)
	lng, _ := strconv.ParseFloat(strings.Split(data.Value, " ")[1], 64)
	cached := d.Position{Name: data.Key, Latitude: lat, Longitude: lng}

	fmt.Println(cached)
}
