package rec

import (
  "strings"
  "unicode/utf8"
)

type terminal struct {}

func NewTerminal() terminal { return terminal {} }

var badSelection string = `
Woops! Sorry If it wasn't clear.
Please just type the number of the space you would like to choose.
Let's try again :)
`

var playerMovePart1 string = "\nIt's player "
var playerMovePart2 string = "'s turn...\n"

var padding string = " "
var bulkhead string = "|"
var decking string = "-"
var joint string = "+"
var nextRow string = "\n"

func (t terminal) BadSelection() string { return badSelection }

func (t terminal) PlayerMove(name string) string {
  return playerMovePart1 + name + playerMovePart2
}

func (t terminal) RowSegment(name string, i, end, cellSize int) string {
  return t.segment(name, padding, bulkhead, i, end, cellSize) }

func (t terminal) DeckSegment(i, end, cellSize int) string {
  return t.segment(decking, decking, joint, i, end, cellSize)
}

func (t terminal) segment(sigil, fill, div string, i, end, cellSize int) string {
  result := fill
  frontPadding := t.trueLen(fill)
  backPadding := cellSize - t.trueLen(sigil) - frontPadding
  result += sigil
  result += strings.Repeat(fill, backPadding)
  if t.itsNotTheLastOne(i, end) {
    result += div
  } else {
    result += nextRow
  }
  return result
}

func (t terminal) trueLen(sigil string) int {
  return utf8.RuneCountInString(sigil)
}

func (t terminal) itsNotTheLastOne(i, end int) bool {
  return i+1 != end
}
