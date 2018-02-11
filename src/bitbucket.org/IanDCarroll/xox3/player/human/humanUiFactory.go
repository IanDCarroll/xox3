package human

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/rec"
  "bitbucket.org/IanDCarroll/xox3/board"
  "bitbucket.org/IanDCarroll/xox3/player/human/boardRec"
  "bitbucket.org/IanDCarroll/xox3/player/human/messageRec"
)

var terminalOut shell.ShellOut = shell.NewTerminal()
var terminalIn shell.ShellIn = shell.NewTerminal()
var messengerRec messageRec.Rec = rec.NewEnglish()

func BuildHumanUi(markers []string, board board.Board) ui.Ui {
  display := BuildInGameTerminalDisplay(markers, terminalOut, board)
  selector := buildInGameTerminalSelector(terminalIn)
  return NewUI(display, selector, board, messengerRec)
}

func BuildInGameTerminalDisplay(m []string, s shell.ShellOut, b board.Board) display.Display {
  messenger := BuildMessageDisplay(s)
  carpenter := BuildBoardDisplay(m, s, b)
  return NewInGameTerminalDisplay(messenger, carpenter)
}

func BuildMessageDisplay(shell shell.ShellOut) display.Display {
  return NewMessageDisplay(shell, messengerRec)
}

func BuildBoardDisplay(markers []string, shell shell.ShellOut, board board.Board) display.Display {
  rec := BuildBoardRec(markers)
  return NewBoardDisplay(shell, rec, board)
}

func BuildBoardRec(markers []string) boardRec.Rec {
  return boardRec.NewTerminalRec(markers)
}

func buildInGameTerminalSelector(shell shell.ShellIn) selector.Selector {
  return NewInGameTerminalSelector(shell)
}
