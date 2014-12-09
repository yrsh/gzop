package structs

import (
	"encoding/json"
)

type Geojson struct {
	Type     string                 `json:"type"`
	Crs      map[string]interface{} `json:"crs"`
	Features []Feature              `json:"features"`
}

type Feature struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   json.RawMessage        `json:"geometry"`
	Bbox       []float64              `json:"bbox"` //lng,lat
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

type ExportGeom struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}
