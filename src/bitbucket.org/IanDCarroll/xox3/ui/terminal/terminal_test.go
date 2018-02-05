package terminal

import (
  "testing"
  "strings"
)

func TestWelcomeReturnsAStringWithWelcome(t *testing.T) {
  //Given no input
  //When
  actual := Welcome()
  //Then
  expected := "Welcome"
  assertContains(t, actual, expected)
}

func assertContains(t *testing.T, actual string, expected string) {
  if !strings.Contains(actual, expected) {
    t.Errorf("\n%q\nshould contain\n%q", actual, expected)
  }
}
