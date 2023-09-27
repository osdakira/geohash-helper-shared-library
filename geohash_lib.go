package main

// go build -buildmode=c-shared -o geohash_lib_go.so ./geohash_lib.go

//#include <stdlib.h>
import "C"

import (
	"fmt"
	"log"
	"strings"
	"unsafe"
	"strconv"

	"github.com/mmcloughlin/geohash"

	"github.com/peterstace/simplefeatures/geom"
)

//export MakeGeohashWithPrecisionFromPolygonWkt
func MakeGeohashWithPrecisionFromPolygonWkt(wkt_polygon *C.char, chars C.int) *C.char {
	wktPolygon := C.GoString(wkt_polygon)
	geohashes := makeGeohashWithPrecisionFromPolygonWkt(wktPolygon, uint(chars))
	joinedString := strings.Join(geohashes, ",")
	return C.CString(joinedString) // 配列での返し方がわからないので、文字列にして返す
}

func makeGeohashWithPrecisionFromPolygonWkt(wktPolygon string, chars uint) []string {
	polygon, err := geom.UnmarshalWKT(wktPolygon)
	if err != nil {
		log.Fatal(err)
	}
	return makeGeohashWithPrecisionFromPolygon(polygon, "", chars)
}

func makeGeohashWithPrecisionFromPolygon(polygon geom.Geometry, baseHash string, chars uint) []string {
	geohashes := make(map[string]struct{})

	// Set が無いので、 uint64 を key にした構造体で実装
	visited := make(map[string]struct{})
	search := make(map[string]struct{})
	search[baseHash] = struct{}{} // 起点として hash を格納
	bits := int(chars) * 5

	for len(geohashes) == 0 || len(search) != 0 {
		// fmt.Println(search)
		// golang に .first も .pop も無いので、要素を一つ取り出すのに range して break する
		var hash string
		for k, _ := range search {
			hash = k
			break
		}
		delete(search, hash)
		visited[hash] = struct{}{}

		// 期待する長さになったら終了
		if len(hash) >= bits {
			hashInt, err := strconv.ParseUint(hash, 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			geohashStr := geohash.ConvertIntToString(hashInt, chars)
			geohashes[geohashStr] = struct{}{}
		} else {
			// 半分に分割して、 1 段階小さくする
			lHash := hash + "0"
			rHash := hash + "1"

			// １ 段階小さくしたものが、ポリゴンと接していたら、 search 対象
			if isInPolygon(visited, lHash, polygon) {
				search[lHash] = struct{}{}
			}
			if isInPolygon(visited, rHash, polygon) {
				search[rHash] = struct{}{}
			}
		}
	}
	return fetchKeys(geohashes)
}

func isInPolygon(visited map[string]struct{}, hash string, bounds geom.Geometry) bool {
	_, hasKey := visited[hash]
	if hasKey {
		return false
	}
	poly := decodeBinary(hash)
	return geom.Intersects(bounds, poly)
}

func decodeBinary(binaryString string) geom.Geometry {
	b := [2][2]float64{
		{-90.0, +90.0},
		{-180.0, +180.0},
	}
	for i, c := range binaryString {
		bit, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal(err)
		}
		k := 1 - (i % 2)
		b[k][bit ^ 1] = (b[k][0] + b[k][1]) / 2
    }
	input := fmt.Sprintf(
		"POLYGON((%f %f,%f %f,%f %f,%f %f, %f %f))",
		b[1][0], b[0][0],
		b[1][0], b[0][1],
		b[1][1], b[0][1],
		b[1][1], b[0][0],
		b[1][0], b[0][0],
	)
	p, err := geom.UnmarshalWKT(input)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

//export IsIntersect
func IsIntersect(c_a, c_b *C.char) int {
	a := C.GoString(c_a)
	b := C.GoString(c_b)
	if isIntersect(a, b) {
		return 1
	} else {
		return 0
	}
}

func isIntersect(a string, b string) bool {
	return strings.HasPrefix(a, b) || strings.HasPrefix(b, a)
}

//export IntersectGeohashes
func IntersectGeohashes(
	geohashesCA **C.char, sizeCA C.int,
	geohashesCB **C.char, sizeCB C.int) *C.char {

	geohashesA := convertArrayC2Go(geohashesCA, int(sizeCA))
	geohashesB := convertArrayC2Go(geohashesCB, int(sizeCB))

	geohashToSizeA := makeGeohashToSizeMap(geohashesA)
	geohashToSizeB := makeGeohashToSizeMap(geohashesB)

	geohashes := makeIntersectGeohashes(geohashToSizeA, geohashToSizeB)

	joinedString := strings.Join(geohashes, ",")
	return C.CString(joinedString) // 配列での返し方がわからないので、文字列にして返す
}

func makeIntersectGeohashes(geohashToSizeA map[string]int, geohashToSizeB map[string]int) []string {
	intersected := make(map[string]struct{})

	for b, sizeB := range geohashToSizeB {
		for a, sizeA := range geohashToSizeA {
			if sizeA > sizeB { // a の方が長い = a のエリアの方が小さい
				if strings.HasPrefix(a, b) { // b の範囲に a が含まれている
					intersected[a] = struct{}{}
				}
			} else if strings.HasPrefix(b, a) { // a の範囲に b が含まれている
				intersected[b] = struct{}{}
			}
		}
	}

	return fetchKeys(intersected)
}

func fetchKeys(hash map[string]struct{}) []string {
	keys := make([]string, len(hash))
	i := 0
	for k := range hash {
		keys[i] = k
		i++
	}
	return keys
}

func makeGeohashToSizeMap(geohashes []string) map[string]int {
	geohashToSize := make(map[string]int)
	for _, value := range geohashes {
		geohashToSize[value] = len(value)
	}
	return geohashToSize
}

func convertArrayC2Go(cstring **C.char, size int) []string {
	start := unsafe.Pointer(cstring)
	startPos := uintptr(start)
	pointerSize := unsafe.Sizeof(cstring)

	newStrings := make([]string, size)

	for i := 0; i < size; i++ {
		nextPos := startPos + uintptr(i)*pointerSize
		pointer := (**C.char)(unsafe.Pointer(nextPos))
		newStrings[i] = C.GoString(*pointer)
	}

	return newStrings
}

//export Free
func Free(p *C.char) {
	C.free(unsafe.Pointer(p))
}

func main() {}
