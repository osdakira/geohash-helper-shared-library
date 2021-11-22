package main

// go build -buildmode=c-shared -o geohash_lib_go.so ./geohash_lib.go

//#include <stdlib.h>
import "C"

import (
	"sort"
	"strings"
	"unsafe"
)

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

	geohashes := reduceOutside(intersected)
	return geohashes
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

func reduceOutside(geohashesMap map[string]struct{}) []string {
	// ジオハッシュの長い順(エリアの小さい順)にソート
	geohashes := fetchKeys(geohashesMap)
	sort.Slice(geohashes, func(i, j int) bool { return len(geohashes[i]) > len(geohashes[j]) })

	// 小さいエリアから順番に、自分を含んでいる大きなエリアを判定し、有ったら消していく
	for _, a := range geohashes {
		for b, _ := range geohashesMap {
			if isInSide(a, b) { // a は b の内部にあるので、 a が優先され b は削除できる
				delete(geohashesMap, b)
			}
		}
	}

	return fetchKeys(geohashesMap)
}

func isInSide(a, b string) bool {
	// Copy from https://golang.org/src/strings/strings.go .HasPrefix
	// a の方が長く、 b と一致するならば、 エリア a は エリア b の内部にある(接していない)
	return len(a) > len(b) && a[0:len(b)] == b
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
