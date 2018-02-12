package gameRunner

import (
  "testing"
)

func TestAiEndgameReturnsADrawMessage(t * testing.T) {
  //Given
  aiPlayer := 2
  rec := buildGameRec()
  subject := NewEndGameDisplay(aiPlayer, rec, displayStub {})
  winner := 0
  //When
  actual := subject.aiEndgame(winner)
  //Then
  expected := rec.AiDraw()
  assertStringEqual(t, actual, expected)
}

func TestAiEndgameReturnsAnAiWinMessage(t * testing.T) {
  //Given
  aiPlayer := 2
  rec := buildGameRec()
  subject := NewEndGameDisplay(aiPlayer, rec, displayStub {})
  winner := 2
  //When
  actual := subject.aiEndgame(winner)
  //Then
  expected := rec.AiWin()
  assertStringEqual(t, actual, expected)
}

func TestAiEndgameReturnsAPlayerWinMessage(t * testing.T) {
  //Given
  aiPlayer := 2
  rec := buildGameRec()
  subject := NewEndGameDisplay(aiPlayer, rec, displayStub {})
  winner := 1
  //When
  actual := subject.aiEndgame(winner)
  //Then
  expected := rec.PlayerWin("1")
  assertStringEqual(t, actual, expected)
}

func assertStringEqual(t *testing.T, actual, expected string) {
  if actual != expected {
    t.Errorf("\n%q\nshould have been\n%q", actual, expected)
  }
}

type displayStub struct {}
func (d displayStub) WriteToShell(board interface{}) {}
