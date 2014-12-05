package jsonParser

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	simplifier "github.com/yrsh/simplify-go"
	"structs"
)

func ProcessJSON(geojson structs.Geojson, zoom int) []byte {
	simplified := geojson
	tol := PixelSize(zoom, 256)
	for i := range simplified.Features {
		processFeature(&simplified.Features[i], tol)
	}
	return marshalGeom(simplified)
}

func processFeature(feat *structs.Feature, tol float64) {
	var geom structs.Geometry
	var exGeom structs.ExportGeom
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
		smpl := simplifier.Simplify(lineString, tol, true)
		exGeom.Coordinates = smpl
		exGeom.Type = geom.Type
		(*feat).Geometry = marshalGeom(exGeom)
		(*feat).Bbox = bboxf(findLineBounds(smpl))
		//
	case "Polygon":
		var polygon [][][]float64
		err = json.Unmarshal(geom.Coordinates, &polygon)
		smpl := simplifyPolygon(polygon, tol)
		exGeom.Coordinates = smpl
		exGeom.Type = geom.Type
		(*feat).Geometry = marshalGeom(exGeom)
		(*feat).Bbox = bboxf(findPolygonBounds(smpl))
		//
	case "MultiPoint":
		var multiPoint [][]float64
		err = json.Unmarshal(geom.Coordinates, &multiPoint)
		//
	case "MultilineString":
		var multiLineString [][][]float64
		err = json.Unmarshal(geom.Coordinates, &multiLineString)
		//
	case "MultiPolygon":
		var multiPolygon [][][][]float64
		err = json.Unmarshal(geom.Coordinates, &multiPolygon)
		smpl := simplifyMultiPolygon(multiPolygon, tol)
		exGeom.Coordinates = smpl
		exGeom.Type = geom.Type
		(*feat).Geometry = marshalGeom(exGeom)
		(*feat).Bbox = bboxf(findMultiPolygonBounds(smpl))
		//
	default:
		err = errors.New("Unknown type of geometry")
	}
	if err != nil {
		log.Print(err)
	}

}

//-----------------------------------------

func simplifyPolygon(pl [][][]float64, tol float64) [][][]float64 {
	r := make([][][]float64, len(pl))
	for i := range pl {
		r[i] = simplifier.Simplify(pl[i], tol, true)
	}
	return r
}

func simplifyMultiPolygon(mpl [][][][]float64, tol float64) [][][][]float64 {
	r := make([][][][]float64, len(mpl))
	for i := range mpl {
		r[i] = simplifyPolygon(mpl[i], tol)
	}
	return r
}

//-----------------------------------------
func findLineBounds(ls [][]float64) [][]float64 {
	minLng := ls[0][0]
	maxLng := ls[0][0]
	minLat := ls[0][1]
	maxLat := ls[0][1]
	for i := range ls {
		if ls[i][0] < minLng {
			minLng = ls[i][0]
		}
		if ls[i][0] > maxLng {
			maxLng = ls[i][0]
		}
		if ls[i][1] < minLat {
			minLat = ls[i][1]
		}
		if ls[i][1] > maxLat {
			maxLat = ls[i][1]
		}
	}
	return [][]float64{{minLng, minLat}, {maxLng, maxLat}}
}

func findPolygonBounds(pl [][][]float64) [][]float64 {
	var bounds [][]float64
	for i := range pl {
		b := findLineBounds(pl[i])
		bounds = append(bounds, b[0], b[1])
	}
	return findLineBounds(bounds)
}

func findMultiPolygonBounds(mpl [][][][]float64) [][]float64 {
	var bounds [][]float64
	for i := range mpl {
		b := findPolygonBounds(mpl[i])
		bounds = append(bounds, b[0], b[1])
	}
	return findLineBounds(bounds)
}

func bboxf(f [][]float64) []float64 {
	return []float64{f[0][0], f[0][1], f[1][0], f[1][1]}
}

//---------------------------------------------------

func marshalGeom(i interface{}) []byte {
	b, err := json.Marshal(i)
	if err == nil {
		return b
	} else {
		log.Print(err)
		return []byte{}
	}
}

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
