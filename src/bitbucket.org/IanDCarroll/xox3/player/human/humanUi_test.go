package human

import (
  "testing"
  "strconv"
  "bitbucket.org/IanDCarroll/xox3/board"
  "bitbucket.org/IanDCarroll/xox3/player/human/rec"
)

func TestSpaceIsUnavailableReturnsTrueIfTheSpaceIsTaken(t *testing.T) {
  //Given
  board := board.NewMNKX(3)
  subject := NewUI(displayStub {}, selectorMock {}, board, rec.NewTerminal())
  board.Mark(0, 1)
  //When
  actual := subject.spaceIsUnavailable(1)
  //Then
  expected := true
  assertEqual(t, actual, expected)
}

func TestSpaceIsUnavailableReturnsTrueIfTheSpaceIsTooHigh(t *testing.T) {
  //Given
  board := board.NewMNKX(3)
  subject := NewUI(displayStub {}, selectorMock {}, board, rec.NewTerminal())
  //When
  actual := subject.spaceIsUnavailable(10)
  //Then
  expected := true
  assertEqual(t, actual, expected)
}

func TestSpaceIsUnavailableReturnsTrueIfTheSpaceIsOffTooLow(t *testing.T) {
  //Given
  board := board.NewMNKX(3)
  subject := NewUI(displayStub {}, selectorMock {}, board, rec.NewTerminal())
  //When
  actual := subject.spaceIsUnavailable(0)
  //Then
  expected := true
  assertEqual(t, actual, expected)
}

func TestSpaceIsUnavailableReturnsFalseIfTheSpaceIsEmpty(t *testing.T) {
  //Given
  board := board.NewMNKX(3)
  subject := NewUI(displayStub {}, selectorMock {}, board, rec.NewTerminal())
  board.Mark(0, 1)
  //When
  actual := subject.spaceIsUnavailable(2)
  //Then
  expected := false
  assertEqual(t, actual, expected)
}

func TestInvalidReturnsFalseIfTheSpaceIsEmptyAndTheInputIsSane(t *testing.T) {
  //Given
  board := board.NewMNKX(3)
  subject := NewUI(displayStub {}, selectorMock {}, board, rec.NewTerminal())
  saneInput := "5"
  //When
  actual := subject.invalid(strconv.Atoi(saneInput))
  //Then
  expected := false
  assertEqual(t, actual, expected)
}

func TestInvalidReturnsTrueIfTheInputIsSaneButTheSpaceIsUnavailable(t *testing.T) {
  //Given
  board := board.NewMNKX(3)
  subject := NewUI(displayStub {}, selectorMock {}, board, rec.NewTerminal())
  board.Mark(4, 1)
  saneInput := "5"
  //When
  actual := subject.invalid(strconv.Atoi(saneInput))
  //Then
  expected := true
  assertEqual(t, actual, expected)
}

func TestInvalidReturnsTrueIfTheInputIsInsane(t *testing.T) {
  //Given
  board := board.NewMNKX(3)
  subject := NewUI(displayStub {}, selectorMock {}, board, rec.NewTerminal())
  saneInput := "Five"
  //When
  actual := subject.invalid(strconv.Atoi(saneInput))
  //Then
  expected := true
  assertEqual(t, actual, expected)
}

func assertEqual(t *testing.T, actual, expected interface{}) {
  if actual != expected {
    t.Errorf("\n%q\nshould have been\n%q", actual, expected)
  }
}

type displayStub struct {}
func (d displayStub) WriteToShell(board interface{}) {}

type selectorMock struct {}
func (s selectorMock) ReadFromShell() interface{} { return 2 }
