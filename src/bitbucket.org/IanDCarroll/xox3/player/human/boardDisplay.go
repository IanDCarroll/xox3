package human

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/player/human/boardRec"
  "bitbucket.org/IanDCarroll/xox3/board"
  "strconv"
)

type boardDisplay struct {
  shell shell.ShellOut
  rec boardRec.Rec
  board board.Board
  cellSize int
}

func NewBoardDisplay(shell shell.ShellOut, rec boardRec.Rec, board board.Board) boardDisplay {
  return boardDisplay {shell, rec, board, sizeCells(board)}
}

func sizeCells(board board.Board) int {
  power := 1
  factor := 0
  padding := 2
  for board.Size() >= power {
    factor ++
    power = power * 10
  }
  return factor + padding
}

func (b boardDisplay) WriteToShell(message interface{}) {
  b.shell.Write(b.buildTheCurrentBoard())
}

func (b boardDisplay) buildTheCurrentBoard() string {
  var constructedBoard string
  for i := 0; i < b.board.Height(); i++ {
    constructedBoard += b.constructRow(i)
    if b.itsNotTheLastOne(i, b.board.Height()) {
      constructedBoard += b.constructDecking()
    }
  }
  return constructedBoard
}

func (b boardDisplay) constructRow(iteration int) string {
  var constructedRow string
  start := iteration * b.board.Width()
  end := start + b.board.Width()
  for i := start; i < end; i++ {
    constructedRow += b.constructRowSegment(i, end)
  }
  return constructedRow
}

func (b boardDisplay) itsNotTheLastOne(i, end int) bool {
  return i+1 != end
}

func (b boardDisplay) constructDecking() string {
  var constructedDeck string
  for i := 0; i < b.board.Width(); i++ {
    constructedDeck += b.rec.DeckSegment(i, b.board.Width(), b.cellSize)
  }
  return constructedDeck
}

func (b boardDisplay) constructRowSegment(i, end int) string {
  var segment string
  mark := b.board.View()[i]
  if mark == 0 {
    segment += b.rec.RowSegment(strconv.Itoa(i + 1), i, end, b.cellSize)
  } else {
    segment += b.rec.RowSegment(b.rec.Marker(mark), i, end, b.cellSize)
  }
  return segment
}
