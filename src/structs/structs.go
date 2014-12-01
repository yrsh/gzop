package structs

import (
	"encoding/json"
)

type Geojson struct {
	Type     string          `json:"type"`
	Crs      json.RawMessage `json:"crs"`
	Features []Feature       `json:"features"`
}

type Feature struct {
	Type       string          `json:"type"`
	Properties json.RawMessage `json:"properties"`
	Geometry   json.RawMessage `json:"geometry"`
	Bbox       []float64       `json:"bbox"` //lng,lat
}

type Geometry struct {
	Type            string
	Coordinates     json.RawMessage
	Point           []float64
	LineString      [][]float64
	Polygon         [][][]float64
	MultiPoint      [][]float64
	MultiLineString [][][]float64
	MultiPolygon    [][][][]float64
}

// type Point []float64

// type LineString []Point

// type Polygon []Polyline

// type MultiPolygon []Polygon

//`json:"field_a"`
