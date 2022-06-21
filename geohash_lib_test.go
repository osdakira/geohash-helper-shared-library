package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMakeIntersectGeohashes(t *testing.T) {
	geohashesA := []string{"0", "01", "02", "012", "013", "0123", "0124", "012345", "12"}
	geohashesB := []string{"012"}
	expect := []string{"012", "0123", "012345", "0124"}

	geohashToSizeA := makeGeohashToSizeMap(geohashesA)
	geohashToSizeB := makeGeohashToSizeMap(geohashesB)
	actual := makeIntersectGeohashes(geohashToSizeA, geohashToSizeB)
	sort.Strings(actual)
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("%v", actual)
	}
}

func TestincreaseLengthToMax(t *testing.T) {
	geohash := []byte("xn77h")
	actual := increaseLengthToMax(geohash, 7)

	expect := make([][]byte, 32*32)
	i := 0
	for _, v1 := range base32 {
		for _, v2 := range base32 {
			expect[i] = append([]byte{}, geohash...)
			expect[i] = append(expect[i], v1, v2)
			i += 1
		}
	}

	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("%v", actual)
	}
}
