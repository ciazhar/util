package util

type GeoJson struct {
	Lat			float64	`json:"lat" bson:"-"`
	Long		float64	`json:"lng" bson:"-"`
	X 			float64 `json:"x" bson:"-"`
	Y 			float64 `json:"y" bson:"-"`
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

type Area struct {
	Distance  int `json:"distance"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Limit	  int `json:"limit"`
}