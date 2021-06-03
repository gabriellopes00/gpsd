package domain

type Position struct {
	Name      string  `json:"name" valid:"-"`
	Latitude  float64 `json:"latitude" valid:"latitude"`
	Longitude float64 `json:"longitude" valid:"longitude"`
	Distance  uint16  `json:"distance,omitempty" valid:"-"`
}
