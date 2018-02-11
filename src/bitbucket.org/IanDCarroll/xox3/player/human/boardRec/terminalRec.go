package boardRec

import (
  "strings"
  "unicode/utf8"
)

type terminalRec struct {
  markers []string
}

func NewTerminalRec(markers []string) terminalRec { return terminalRec {markers} }

var padding string = " "
var bulkhead string = "|"
var decking string = "-"
var joint string = "+"
var nextRow string = "\n"

func (t terminalRec) RowSegment(name string, i, end, cellSize int) string {
  return t.segment(name, padding, bulkhead, i, end, cellSize) }

func (t terminalRec) DeckSegment(i, end, cellSize int) string {
  return t.segment(decking, decking, joint, i, end, cellSize)
}

func (t terminalRec) Marker(playerNumber int) string {
  return t.markers[t.zeroIndex(playerNumber)]
}

func (t terminalRec) segment(sigil, fill, div string, i, end, cellSize int) string {
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

func (t terminalRec) trueLen(sigil string) int {
  return utf8.RuneCountInString(sigil)
}

func (t terminalRec) itsNotTheLastOne(i, end int) bool {
  return i+1 != end
}

func (t terminalRec) zeroIndex(playerNumber int) int {
  return playerNumber - 1
}
