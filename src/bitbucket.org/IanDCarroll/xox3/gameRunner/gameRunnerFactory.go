package gameRunner

import (
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/rec"
  "bitbucket.org/IanDCarroll/xox3/gameRunner/gameRec"
  "bitbucket.org/IanDCarroll/xox3/player"
  "bitbucket.org/IanDCarroll/xox3/player/human"
  "bitbucket.org/IanDCarroll/xox3/board"
)

func BuildAiGameRunner(humanGoes, marker int, markers []string) GameRunner {
  mappedMarkers := mapMarkers(humanGoes, marker, markers)
  board := board.NewMNKX(3)
  if humanGoes == 1 { return buildHumanFirst(humanGoes, mappedMarkers, board) }
  return buildAiFirst(humanGoes, mappedMarkers, board)
}

func mapMarkers(humanGoes, markerChoice int, markers []string) []string {
  if humanGoes != markerChoice { return []string{markers[1], markers[0]} }
  return markers
}

func buildHumanFirst(humanGoes int, markers []string, board board.Board) GameRunner {
  aiGoes := 2
  human := player.BuildHumanPlayer(humanGoes, markers, board)
  ai := player.BuildAiPlayer(aiGoes, markers, board)
  display := buildEndGameDisplay(aiGoes, markers, board)
  return NewGameRunner(human, ai, display)
}

func buildAiFirst(humanGoes int, markers []string, board board.Board) GameRunner {
  aiGoes := 1
  human := player.BuildHumanPlayer(humanGoes, markers, board)
  ai := player.BuildAiPlayer(aiGoes, markers, board)
  display := buildEndGameDisplay(aiGoes, markers, board)
  return NewGameRunner(ai, human, display)
}

func buildEndGameDisplay(aiPlayer int, markers []string, board board.Board) display.Display {
  rec := buildGameRec()
  host := human.BuildInGameTerminalDisplay(markers, board)
  return NewEndGameDisplay(aiPlayer, rec, host)
}

func buildGameRec() gameRec.Rec {
  return rec.NewEnglish()
}
