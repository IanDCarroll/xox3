package human

import (
  "testing"
  "bitbucket.org/IanDCarroll/xox3/rec"
  "bitbucket.org/IanDCarroll/xox3/player/human/boardRec"
  "bitbucket.org/IanDCarroll/xox3/board"
)

func TestBuildTheCurrentBoardReturnsAnInitialFormattedBoard(t *testing.T) {
  //Given
  subject := NewBoardDisplay(shellStub {}, boardRec.NewTerminalRec(), board.NewMNKX(3))
  //When
  actual := subject.buildTheCurrentBoard()
  //Then
  expected := " 1 | 2 | 3 \n" +
              "---+---+---\n" +
              " 4 | 5 | 6 \n" +
              "---+---+---\n" +
              " 7 | 8 | 9 \n"
  assertStringEqual(t, actual, expected)
}

func TestBuildTheCurrentBoardReturnsAFormatted4x4Board(t *testing.T) {
  //Given
  subject := NewBoardDisplay(shellStub {}, boardRec.NewTerminalRec(), board.NewMNKX(4))
  //When
  actual := subject.buildTheCurrentBoard()
  //Then
  expected := " 1  | 2  | 3  | 4  \n" +
              "----+----+----+----\n" +
              " 5  | 6  | 7  | 8  \n" +
              "----+----+----+----\n" +
              " 9  | 10 | 11 | 12 \n" +
              "----+----+----+----\n" +
              " 13 | 14 | 15 | 16 \n"
  assertStringEqual(t, actual, expected)
}

func TestSizeCellsReturnsTheRightCellSize(t *testing.T) {
  //Given
  board := board.NewMNKX(3)
  //When
  subject := NewBoardDisplay(shellStub {}, boardRec.NewTerminalRec(), board)
  //Then
  actual := subject.cellSize
  expected := 3
  assertIntEqual(t, actual, expected)
}

func TestSizeCellsReturnsTheRightCellSizeFor4x4Boards(t *testing.T) {
  //Given
  board := board.NewMNKX(4)
  //When
  subject := NewBoardDisplay(shellStub {}, boardRec.NewTerminalRec(), board)
  //Then
  actual := subject.cellSize
  expected := 4
  assertIntEqual(t, actual, expected)
}

func TestSizeCellsReturnsTheRightCellSizeFor10x10Boards(t *testing.T) {
  //Given
  board := board.NewMNKX(10)
  //When
  subject := NewBoardDisplay(shellStub {}, boardRec.NewTerminalRec(), board)
  //Then
  actual := subject.cellSize
  expected := 5
  assertIntEqual(t, actual, expected)
}

func TestAnnounceMessageReturnsAMessageFromThePlayersNumber(t *testing.T) {
  //Given
  marker := 1
  subject := NewMessageDisplay(shellStub {}, rec.NewEnglish())
  //When
  actual := subject.announceMessage(marker)
  //Then
  expected := "\nIt's player 1's turn...\n"
  assertStringEqual(t, actual, expected)
}


func TestAnnounceMessageReturnsAnyArbitraryPlayerNumer(t *testing.T) {
  //Given
  marker := 2
  subject := NewMessageDisplay(shellStub {}, rec.NewEnglish())
  //When
  actual := subject.announceMessage(marker)
  //Then
  expected := "\nIt's player 2's turn...\n"
  assertStringEqual(t, actual, expected)
}

func TestAnnounceMessagereturnsAnOverridingMessageString(t *testing.T) {
  //Given
  marker := "This is an overriding message from something other than player"
  subject := NewMessageDisplay(shellStub {}, rec.NewEnglish())
  //When
  actual := subject.announceMessage(marker)
  //Then
  assertStringEqual(t, actual, marker)
}

func assertStringEqual(t *testing.T, actual, expected string) {
  if actual != expected {
    t.Errorf("\n%q\nshould have been\n%q", actual, expected)
  }
}

func assertIntEqual(t *testing.T, actual, expected int) {
  if actual != expected {
    t.Errorf("\n%d\nshould have been\n%d", actual, expected)
  }
}

type shellStub struct { shellOutput string }
func (s shellStub) Read() interface{} { return "2" }
func (s shellStub) Write(message interface{}) { s.shellOutput = message.(string) }
