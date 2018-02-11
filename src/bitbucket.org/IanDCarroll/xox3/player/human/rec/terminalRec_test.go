package rec

import (
  "testing"
  "strings"
)

func TestBadSelectionIsPolite(t *testing.T) {
  //Given
  subject := NewTerminal()
  //When
  actual := subject.BadSelection()
  //Then
  expected := "Please"
  assertContains(t, actual, expected)
}

func assertContains(t *testing.T, actual string, expected string) {
  if !strings.Contains(actual, expected) {
    t.Errorf("\n%q\nshould contain\n%q", actual, expected)
  }
}
