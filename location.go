package util

type GeoJson struct {
	Lat			float64	`json:"lat,omitempty" bson:"-"`
	Long		float64	`json:"long,omitempty" bson:"-"`
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

type Area struct {
	Distance  int `json:"distance"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}