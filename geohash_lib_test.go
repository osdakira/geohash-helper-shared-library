package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMakeIntersectGeohashes(t *testing.T) {
	geohashesA := []string{"0", "01", "02", "012", "013", "0123", "0124", "012345", "12"}
	geohashesB := []string{"012"}
	expect := []string{"012345", "0124"}

	geohashToSizeA := makeGeohashToSizeMap(geohashesA)
	geohashToSizeB := makeGeohashToSizeMap(geohashesB)
	actual := makeIntersectGeohashes(geohashToSizeA, geohashToSizeB)
	sort.Strings(actual)
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("%v", actual)
	}
}
