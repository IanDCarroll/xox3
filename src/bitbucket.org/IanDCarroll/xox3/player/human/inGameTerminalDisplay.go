package human

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/player/human/rec"
  "bitbucket.org/IanDCarroll/xox3/board"
  "strconv"
)

type terminalDisplay struct {
  shell shell.ShellOut
  rec rec.Rec
  board board.Board
  cellSize int
}

func NewInGameTerminalDisplay(shell shell.ShellOut, rec rec.Rec, board board.Board) terminalDisplay {
  return terminalDisplay {shell, rec, board, sizeCells(board)}
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

func (t terminalDisplay) WriteToShell(message interface{}) {
  t.shell.Write(t.announceMessage(message))
  t.shell.Write(t.buildTheCurrentBoard())
}

func (t terminalDisplay) announceMessage(message interface{}) string {
  var toTerminal string
  switch value := message.(type) {
  case int:
    toTerminal = t.announcePlayerMove(value)
  case string:
    toTerminal = t.announceMessageFromBeyond(value)
  }
  return toTerminal
}

func (t terminalDisplay) announcePlayerMove(playerNumber int) string {
  name := strconv.Itoa(playerNumber)
  return t.rec.PlayerMove(name)
}

func (t terminalDisplay) announceMessageFromBeyond(message string) string {
  return message
}

func (t terminalDisplay) buildTheCurrentBoard() string {
  var constructedBoard string
  for i := 0; i < t.board.Height(); i++ {
    constructedBoard += t.constructRow(i)
    if t.itsNotTheLastOne(i, t.board.Height()) {
      constructedBoard += t.constructDecking()
    }
  }
  return constructedBoard
}

func (t terminalDisplay) constructRow(iteration int) string {
  var constructedRow string
  start := iteration * t.board.Width()
  end := start + t.board.Width()
  for i := start; i < end; i++ {
    constructedRow += t.constructRowSegment(i, end)
  }
  return constructedRow
}

func (t terminalDisplay) itsNotTheLastOne(i, end int) bool {
  return i+1 != end
}

func (t terminalDisplay) constructDecking() string {
  var constructedDeck string
  for i := 0; i < t.board.Width(); i++ {
    constructedDeck += t.rec.DeckSegment(i, t.board.Width(), t.cellSize)
  }
  return constructedDeck
}

func (t terminalDisplay) constructRowSegment(i, end int) string {
  var segment string
  if t.board.View()[i] == 0 {
    segment += t.rec.RowSegment(strconv.Itoa(i + 1), i, end, t.cellSize)
  } else {
    segment += t.rec.RowSegment("#", i, end, t.cellSize)
  }
  return segment
}
