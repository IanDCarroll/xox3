package player


import (
  "testing"
  "reflect"
  "bitbucket.org/IanDCarroll/xox3/board"
)

func buildTestPlayer(board board.Board) player {
  display := displayStub {}
  ui := buildPlayerUI(display, selectorMock {})
  rules := buildMNK3Rules()
  marker := 1
  return New(ui, rules, board, marker)
}

func TestPlayerReturnsTheGameState(t *testing.T) {
  //Given
  board := buildMNK3Board()
  subject := buildTestPlayer(board)
  //When
  actualGameOver, actualMarker := subject.TakeATurn()
  //Then
  expectedGameOver := false
  expectedMarker := 0
  assertEqual(t, "GameOver Value", actualGameOver, expectedGameOver)
  assertEqual(t, "Marker Value", actualMarker, expectedMarker)
}

func TestPlayerMarksTheBoard(t *testing.T) {
  //Given
  board := buildMNK3Board()
  subject := buildTestPlayer(board)
  //When
  actualGameOver, actualMarker := subject.TakeATurn()
  //Then
  expectedGameOver := false
  expectedMarker := 0
  actualBoard := board.View()
  expectedBoard := []int{0,0,1, 0,0,0, 0,0,0}
  assertEqual(t, "GameOver Value", actualGameOver, expectedGameOver)
  assertEqual(t, "Marker Value", actualMarker, expectedMarker)
  assertSliceEquals(t, actualBoard, expectedBoard)
}

func TestPlayerReturnsAWonGameState(t *testing.T) {
  //Given
  board := buildMNK3Board()
  subject := buildTestPlayer(board)
  board.Mark(0, 1)
  board.Mark(1, 1)
  //When
  actualGameOver, actualMarker := subject.TakeATurn()
  //Then
  expectedGameOver := true
  expectedMarker := 1
  assertEqual(t, "GameOver Value", actualGameOver, expectedGameOver)
  assertEqual(t, "Marker Value", actualMarker, expectedMarker)
}

func assertEqual(t *testing.T, name string, actual, expected interface{}) {
  if actual != expected {
    t.Errorf("\n%q was %q\nshould have been\n%q", name, actual, expected)
  }
}

func assertSliceEquals(t *testing.T, actual, expected []int) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("\n%d\nshould be\n%d", actual, expected)
  }
}

type displayStub struct {}
func (d displayStub) WriteToShell(board interface{}) {}

type selectorMock struct {}
func (s selectorMock) ReadFromShell() interface{} { return 2 }
