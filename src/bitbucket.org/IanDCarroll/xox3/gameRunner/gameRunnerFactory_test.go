package gameRunner

import (
  "testing"
  "reflect"
)

func TestMapMarkersReturnsTheRightSliceIfHumanGoesFirstAndChoosesMarkerOne(t * testing.T) {
  //Given
  order, marker := 1, 1
  markers := []string{"X", "O"}
  //When
  actual := mapMarkers(order, marker, markers)
  //Then
  assertSliceEquals(t, actual, markers)
}

func TestMapMarkersReturnsTheRightSliceIfHumanGoesFirstAndChoosesMarkerTwo(t * testing.T) {
  //Given
  order, marker := 1, 2
  markers := []string{"X", "O"}
  //When
  actual := mapMarkers(order, marker, markers)
  //Then
  expected := []string{"O", "X"}
  assertSliceEquals(t, actual, expected)
}

func TestMapMarkersReturnsTheRightSliceIfHumanGoesSecondAndChoosesMarkerOne(t * testing.T) {
  //Given
  order, marker := 2, 1
  markers := []string{"X", "O"}
  //When
  actual := mapMarkers(order, marker, markers)
  //Then
  expected := []string{"O", "X"}
  assertSliceEquals(t, actual, expected)
}

func TestMapMarkersReturnsTheRightSliceIfHumanGoesSecondAndChoosesMarkerTwo(t * testing.T) {
  //Given
  order, marker := 2, 2
  markers := []string{"X", "O"}
  //When
  actual := mapMarkers(order, marker, markers)
  //Then
  assertSliceEquals(t, actual, markers)
}


func assertSliceEquals(t *testing.T, actual, expected []string) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("\n%p\nshould be\n%p", actual, expected)
  }
}
