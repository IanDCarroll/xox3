package menu

import (
  "testing"
)

func TestGetGameParamsReturnsTwoInts(t *testing.T) {
  //Given
  shell := shellStub {}
  subject := New(buildMenuUI(shell, shell), buildEnglishRec())
  //When
  actualOrder, actualMarker := subject.GetGameParams()
  //Then
  expected := 2
  assertEqual(t, "Order Value", actualOrder, expected)
  assertEqual(t, "Marker Value", actualMarker, expected)
}

func assertEqual(t *testing.T, name string, actual int, expected int) {
  if actual != expected {
    t.Errorf("\n%q was %q\nshould have been\n%q", name, actual, expected)
  }
}

type shellStub struct { shellOutput string }
func (s shellStub) Read() interface{} { return "2" }
func (s shellStub) Write(message interface{}) { s.shellOutput = message.(string) }
