package board

import (
  "testing"
  "reflect"
)

func TestBoardIsInitializedToAllZeros(t *testing.T) {
  //Given
  subject := NewMNKX(3)
  //When
  actual := subject.View()
  //Then
  expected := []int{0,0,0, 0,0,0, 0,0,0}
  assertSliceEquals(t, actual, expected)
}

func TestBoardCanBeMarked(t *testing.T) {
  //Given
  subject := NewMNKX(3)
  //When
  subject.Mark(4, 1)
  //Then
  actual := subject.View()
  expected := []int{0,0,0, 0,1,0, 0,0,0}
  assertSliceEquals(t, actual, expected)
}

func TestBoardRetainsItsState(t *testing.T) {
  //Given
  subject := NewMNKX(3)
  //When
  subject.Mark(4, 1)
  subject.Mark(1, 2)
  //Then
  actual := subject.View()
  expected := []int{0,2,0, 0,1,0, 0,0,0}
  assertSliceEquals(t, actual, expected)
}

func TestSizeReturnsSizeOfTheBoard(t *testing.T) {
  //Given
  subject := NewMNKX(3)
  //When
  actual := subject.Size()
  //Then
  expected := 9
  assertEqual(t, actual, expected)
}

func TestAvailableSpacesReturnsTheIndeciesOfAll0ValueSpaces(t *testing.T) {
  //Given
  subject := NewMNKX(3)
  //When
  actual := subject.AvailableSpaces()
  //Then
  expected := []int{0,1,2,3,4,5,6,7,8}
  assertSliceEquals(t, actual, expected)
}

func TestAvailableSpacesReturnsTheIndeciesOfAPlayedBoard(t *testing.T) {
  //Given
  subject := NewMNKX(3)
  subject.Mark(4, 1)
  //When
  actual := subject.AvailableSpaces()
  //Then
  expected := []int{0,1,2,3,5,6,7,8}
  assertSliceEquals(t, actual, expected)
}

func assertEqual(t *testing.T, actual, expected interface{}) {
  if actual != expected {
    t.Errorf("\n%q\nshould have been\n%q", actual, expected)
  }
}

func assertSliceEquals(t *testing.T, actual, expected []int) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("\n%d\nshould be\n%d", actual, expected)
  }
}
