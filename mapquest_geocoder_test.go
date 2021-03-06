package geo

import (
	"fmt"
	"testing"
)

// TODO Test extracting LatLng from Mapquest Response
func TestMapQuestExtractLatLngFromRequest(t *testing.T) {
	g := &MapQuestGeocoder{}

	data, err := GetMockResponse("test/data/mapquest_geocode_success.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	point, err := g.extractLatLngFromResponse(data)
	if err != nil {
		t.Error("%v\n", err)
	}

	if point.lat != 37.62181845 || point.lng != -122.383992092462 {
		t.Error(fmt.Sprintf("Expected: [37.62181845, -122.383992092462], Got: [%f, %f]", point.lat, point.lng))
	}
}

// TODO Test extracting LatLng from Mapquest Response when no results are returned
func TestMapQuestExtractLatLngFromRequestZeroResults(t *testing.T) {
	g := &MapQuestGeocoder{}

	data, err := GetMockResponse("test/data/mapquest_geocode_zero_results.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	_, err = g.extractLatLngFromResponse(data)
	if err != mapquestZeroResultsError {
		t.Error(fmt.Sprintf("Expected error: %v, Got: %v"), mapquestZeroResultsError, err)
	}
}
