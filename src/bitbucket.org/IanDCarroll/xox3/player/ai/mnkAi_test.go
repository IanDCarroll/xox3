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

func TestAiIdentityDefaultsToPlayingFirst(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  actualID := subject.identity
  actualOP := subject.opponent
  //Then
  expectedID := 1
  expectedOP := 2
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

func TestAiIdentityCanBeSetToPlayingSecond(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  subject.SetID(2)
  //Then
  actualID := subject.identity
  actualOP := subject.opponent
  expectedID := 2
  expectedOP := 1
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

func TestAiIdentityIsUnaffectedBySubesequentSetIDCalls(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  subject.SetID(2)
  subject.SetID(2)
  //Then
  actualID := subject.identity
  actualOP := subject.opponent
  expectedID := 2
  expectedOP := 1
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

func TestAiIdentityIsChangedIfSubesequentSetIDCallsAreDifferent(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  subject.SetID(2)
  subject.SetID(1)
  //Then
  actualID := subject.identity
  actualOP := subject.opponent
  expectedID := 1
  expectedOP := 2
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

func TestAiIdentityIsUnaffectedBySetIDCallsWithDefaultValues(t * testing.T) {
  //Given
  rules := rules.NewMNKX(3)
  board := board.NewMNKX(3)
  subject := NewMNKAi(rules, board)
  //When
  subject.SetID(1)
  //Then
  actualID := subject.identity
  actualOP := subject.opponent
  expectedID := 1
  expectedOP := 2
  assertIntEqual(t, actualID, expectedID)
  assertIntEqual(t, actualOP, expectedOP)
}

func assertIntEqual(t *testing.T, actual, expected int) {
  if actual != expected {
    t.Errorf("\n%d\nshould have been\n%d", actual, expected)
  }
}
