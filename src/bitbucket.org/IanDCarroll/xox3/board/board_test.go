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

func assertSliceEquals(t *testing.T, actual, expected []int) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("\n%d\nshould be\n%d", actual, expected)
  }
}
