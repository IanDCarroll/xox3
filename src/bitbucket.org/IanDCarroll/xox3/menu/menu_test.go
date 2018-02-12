package menu

import (
  "testing"
  "reflect"
)

func TestGetGameParamsReturnsTwoIntsAndASlice(t *testing.T) {
  //Given
  shell := shellStub {}
  rec := buildEnglishRec()
  subject := New(buildMenuUI(shell, shell), rec)
  //When
  actualOrder, actualMarker, actualSlice := subject.GetGameParams()
  //Then
  expected := 2
  expectedSlice := rec.Markers()
  assertEqual(t, "Order Value", actualOrder, expected)
  assertEqual(t, "Marker Value", actualMarker, expected)
  assertSliceEquals(t, actualSlice, expectedSlice)
}

func assertEqual(t *testing.T, name string, actual int, expected int) {
  if actual != expected {
    t.Errorf("\n%q was %q\nshould have been\n%q", name, actual, expected)
  }
}

func assertSliceEquals(t *testing.T, actual, expected []string) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("\n%p\nshould be\n%p", actual, expected)
  }
}

type shellStub struct { shellOutput string }
func (s shellStub) Read() interface{} { return "2" }
func (s shellStub) Write(message interface{}) { s.shellOutput = message.(string) }
