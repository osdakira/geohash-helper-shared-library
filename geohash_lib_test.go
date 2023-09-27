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

func TestConvertPolygon2GeohashWithPrecision(t *testing.T) {
	polygonWkt := "POLYGON ((132.9709406620001 34.11831060600008, 132.9704386200001 34.11865367400003, 132.9697404750001 34.11975379000006, 132.96891048500004 34.11988086000002, 132.96853966600008 34.11865384600003, 132.96873584000002 34.11737278000004, 132.969587982 34.11632664000007, 132.97019897200005 34.116362571000025, 132.97015482900008 34.11731862700003, 132.9707228300001 34.11726456400004, 132.97113777200002 34.11758954100003, 132.9709406620001 34.11831060600008))"
	expect := []string{"wynd1f9q",	"wynd1f9r",	"wynd1f9w",	"wynd1f9x",	"wynd1f9y",	"wynd1f9z",	"wynd1fbf",	"wynd1fbg",	"wynd1fbu",	"wynd1fbv",	"wynd1fby",	"wynd1fbz",	"wynd1fc0",	"wynd1fc1",	"wynd1fc2",	"wynd1fc3",	"wynd1fc4",	"wynd1fc5",	"wynd1fc6",	"wynd1fc7",	"wynd1fc8",	"wynd1fc9",	"wynd1fcb",	"wynd1fcc",	"wynd1fcd",	"wynd1fce",	"wynd1fcf",	"wynd1fcg",	"wynd1fch",	"wynd1fcj",	"wynd1fck",	"wynd1fcm",	"wynd1fcn",	"wynd1fcp",	"wynd1fcq",	"wynd1fcr",	"wynd1fcs",	"wynd1fct",	"wynd1fcu",	"wynd1fcv",	"wynd1fcw",	"wynd1fcx",	"wynd1fcy",	"wynd1fcz",	"wynd1ff5",	"wynd1ff7",	"wynd1ffh",	"wynd1ffj",	"wynd1ffk",	"wynd1ffm",	"wynd1ffn",	"wynd1ffp",	"wynd1ffq",	"wynd1ffr",	"wynd1ffs",	"wynd1fft",	"wynd1ffw",	"wynd1ffx",	"wynd1g0b",	"wynd1g0c",	"wynd1g0d",	"wynd1g0e",	"wynd1g0f",	"wynd1g0g",	"wynd1g0u",	"wynd1g0v",	"wynd1g0y",	"wynd1g0z",	"wynd1g10",	"wynd1g11",	"wynd1g12",	"wynd1g13",	"wynd1g14",	"wynd1g15",	"wynd1g16",	"wynd1g17",	"wynd1g18",	"wynd1g19",	"wynd1g1b",	"wynd1g1c",	"wynd1g1d",	"wynd1g1e",	"wynd1g1f",	"wynd1g1g",	"wynd1g1h",	"wynd1g1j",	"wynd1g1k",	"wynd1g1m",	"wynd1g1n",	"wynd1g1p",	"wynd1g1q",	"wynd1g1r",	"wynd1g1s",	"wynd1g1t",	"wynd1g1u",	"wynd1g1v",	"wynd1g1w",	"wynd1g1x",	"wynd1g1y",	"wynd1g1z",	"wynd1g2b",	"wynd1g2c",	"wynd1g2f",	"wynd1g30",	"wynd1g31",	"wynd1g32",	"wynd1g33",	"wynd1g34",	"wynd1g36",	"wynd1g38",	"wynd1g39",	"wynd1g3b",	"wynd1g3d",	"wynd1g40",	"wynd1g41",	"wynd1g42",	"wynd1g43",	"wynd1g44",	"wynd1g45",	"wynd1g46",	"wynd1g48",	"wynd1g49",	"wynd1g4h",	"wynd1g4j"}

	actual := makeGeohashWithPrecisionFromPolygonWkt(polygonWkt, 8)
	sort.Strings(actual)
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("%v", actual)
	}
}
