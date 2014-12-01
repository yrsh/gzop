package jsonParser

import (
	"encoding/json"
	"errors"
	"log"
	"simplifier"
	"structs"
)

func ProcessJSON(geojson structs.Geojson, zoom int) {
	simplified := geojson
	tol := PixelSize(zoom, 256)
	err := errors.New("error parsing feature")
	for i := range simplified.Features {
		//simplified.Features[i] = processFeature(simplified.Features[i], tol)
		processFeature(&simplified.Features[i], tol)
	}
}

func processFeature(feat *structs.Feature, tol float64) {
	var geom structs.Geometry
	err := json.Unmarshal((*feat).Geometry, &geom)
	if err != nil {
		log.Print(err)
	}
	//
	switch geom.Type {
	case "Point":
		var point []float64
		err = json.Unmarshal(geom.Coordinates, &point)
		//
	case "LineString":
		var lineString [][]float64
		err = json.Unmarshal(geom.Coordinates, &lineString)
		geom.LineString = simplifier.Simplify(lineString, tol, true)
		(*feat).Geometry = marshalGeom(geom.LineString)
		//
	case "Polygon":
		var polygon [][][]float64
		err = json.Unmarshal(geom.Coordinates, &polygon)
		geom.Polygon = simplifyPolygon(polygon, tol, true)
		(*feat).Geometry = marshalGeom(geom.Polygon)
		//
	case "MultiPoint":
		var multiPoint [][]float64
		err = json.Unmarshal(geom.Coordinates, &multiPoint)
		//
	case "MultilineString":
		var multilineString [][][]float64
		err = json.Unmarshal(geom.Coordinates, &multiLineString)
		//
	case "MultiPolygon":
		var multiPolygon [][][][]float64
		err = json.Unmarshal(geom.Coordinates, &multiPolygon)
		geom.MultiPolygon = simplifyPolygon(multiPolygon, tol, true)
		(*feat).Geometry = marshalGeom(MultiPolygon)
		//
	default:
		err = errors.New("Unknown type of geometry")
	}

}

// func simplifyLineString(ls [][]float64, tol float64) [][]float64 {
// 	return simplifier.Simplify(ls, tol, true)
// }

func simplifyPolygon(pl [][][]float64, tol float64) [][][]float64 {
	r := make([][][]float64, len(pl))
	for i := range pl {
		r[i] = simplifier.Simplify(pl[i], tol, true)
	}
	return r
}

func simplifyMultiPolygon(mpl [][][][]float64, tol float64) [][][]float64 {
	r := make([][][][]float64, len(mpl))
	for i := range mpl {
		r[i] = simplifyPolygon(mpl[i], tol)
	}
	return r
}

func marshalGeom(i interface{}) []byte {
	b, err := json.Marshal(i)
	if err == nil {
		return b
	} else {
		log.Print(err)
		return i
	}
}

//---------------------------------------------------
func PixelSize(zoom, tileSize int) float64 {
	_, latT := Conv(180, -85, float64(zoom), float64(tileSize))
	YmaxPx := latT * 256
	YPx := 170.0 / float64(YmaxPx)
	return YPx
}

func Conv(lng float64, lat float64, zoom, tileSize float64) (int, int) {
	_lat := (1 - math.Log(math.Tan(lat*math.Pi/180)+
		1/math.Cos(lat*math.Pi/180))/math.Pi) /
		2 * math.Pow(2, zoom)
	lattile := int(math.Floor(_lat))
	_lng := (lng + 180) / 360 * math.Pow(2, zoom)
	lngtile := int(math.Floor(_lng))
	return lngtile, lattile
}
