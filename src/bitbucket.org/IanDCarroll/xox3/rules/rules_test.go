package rules

import (
  "testing"
  "reflect"
)

func TestRulesReturnFalse0ForAnEmptyBoard(t *testing.T) {
  //Given
  board := []int{0,0,0, 0,0,0, 0,0,0}
  subject := NewMNKX(3)
  //When
  actualBool, actualPlayer := subject.Gameover(board)
  //Then
  expectedBool := false
  expectedPlayer := 0
  assertJudgementEquals(t, actualBool, expectedBool, actualPlayer, expectedPlayer)
}

func TestRulesReturnTrue0ForADrawnGame(t *testing.T) {
  //Given
  board := []int{1,1,2,
                 2,1,1,
                 1,2,2}
  subject := NewMNKX(3)
  //When
  actualBool, actualPlayer := subject.Gameover(board)
  //Then
  expectedBool := true
  expectedPlayer := 0
  assertJudgementEquals(t, actualBool, expectedBool, actualPlayer, expectedPlayer)
}

func TestRulesReturnTrue1ForAFullP1WonGame(t *testing.T) {
  //Given
  board := []int{2,1,2,
                 1,1,2,
                 2,1,1}
  subject := NewMNKX(3)
  //When
  actualBool, actualPlayer := subject.Gameover(board)
  //Then
  expectedBool := true
  expectedPlayer := 1
  assertJudgementEquals(t, actualBool, expectedBool, actualPlayer, expectedPlayer)
}

func TestRulesReturnTrue1ForASparseP1WonGame(t *testing.T) {
  //Given
  board := []int{0,0,0,
                 1,1,1,
                 0,2,2}
  subject := NewMNKX(3)
  //When
  actualBool, actualPlayer := subject.Gameover(board)
  //Then
  expectedBool := true
  expectedPlayer := 1
  assertJudgementEquals(t, actualBool, expectedBool, actualPlayer, expectedPlayer)
}

func TestConstructRulesReturnsGoodRulesForAnMNK3Game(t *testing.T) {
  //Given
  //When
  actual := constructMNKXRules(3)
  //Then
  expected := [][]int{
    {0, 1, 2},
    {3, 4, 5},
    {6, 7, 8},

    {0, 3, 6},
    {1, 4, 7},
    {2, 5, 8},

    {0, 4, 8},
    {2, 4, 6},
  }
  assertSliceEquals(t, actual, expected)
}

func assertJudgementEquals(t *testing.T, acB, exB bool, acP, exP int) {
  if acB != exB || acP != exP {
    t.Errorf("\n%t %d\nshould be\n%t %d", acB, acP, exB, exP)
  }
}

func assertSliceEquals(t *testing.T, actual, expected interface{}) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("\n%d\nshould be\n%d", actual, expected)
  }
}
