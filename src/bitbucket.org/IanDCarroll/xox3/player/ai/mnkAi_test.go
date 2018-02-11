package ai

import (
  "testing"
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
)

func TestDumbChoiceReturnsZeroFromAnEmptyBoard(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  actual := subject.dumbButLegalMove()
  //Then
  expected := 0
  assertIntEqual(t, actual, expected)
}

func TestDumbChoiceReturnsTheFirstAvalilableSpaceFromAPlayedBoard(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  board.Mark(0, 1)
  //When
  actual := subject.dumbButLegalMove()
  //Then
  expected := 1
  assertIntEqual(t, actual, expected)
}

func assertIntEqual(t *testing.T, actual, expected int) {
  if actual != expected {
    t.Errorf("\n%d\nshould have been\n%d", actual, expected)
  }
}
