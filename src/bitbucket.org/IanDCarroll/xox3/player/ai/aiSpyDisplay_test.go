package ai

import (
  "testing"
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
)

func TestAiSpyDisplayTellsAiWhoItIs(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  ai := NewMNKAi(rules, board)
  subject := buildAiSpyDisplay(displayStub {}, ai)
  //When
  subject.WriteToShell(2)
  //Then
  actualID := ai.identity
  actualOP := ai.opponent
  expectedID := 2
  expectedOP := 1
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

type displayStub struct {}
func (d displayStub) WriteToShell(board interface{}) {}
