package player

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
  "bitbucket.org/IanDCarroll/xox3/player/human"
  "bitbucket.org/IanDCarroll/xox3/player/ai"
)

var terminalOut shell.ShellOut = shell.NewTerminal()

func BuildHumanPlayer(playerNumber int, markers []string, board board.Board) player {
  ui := human.BuildHumanUi(markers, board)
  rules := buildMNK3Rules()
  return New(ui, rules, board, playerNumber)
}

func BuildAiPlayer(playerNumber int, markers []string, board board.Board) player {
  spy := human.BuildInGameTerminalDisplay(markers, terminalOut, board)
  ui := ai.BuildAiUi(board, spy)
  rules := buildMNK3Rules()
  return New(ui, rules, board, playerNumber)
}

func buildPlayerUI(display display.Display, selector selector.Selector) ui.Ui {
  return ui.New(display, selector)
}

func buildMNK3Rules() rules.Rules {
  return rules.NewMNKX(3)
}
