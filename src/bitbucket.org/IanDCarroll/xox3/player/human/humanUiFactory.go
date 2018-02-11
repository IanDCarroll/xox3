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

func BuildHumanUi(rec boardRec.Rec, board board.Board) ui.Ui {
  display := BuildInGameTerminalDisplay(terminalOut, rec, board)
  selector := buildInGameTerminalSelector(terminalIn)
  return NewUI(display, selector, board, messengerRec)
}

func BuildInGameTerminalDisplay(s shell.ShellOut, r boardRec.Rec, b board.Board) display.Display {
  messenger := BuildMessageDisplay(s)
  carpenter := BuildBoardDisplay(s, r, b)
  return NewInGameTerminalDisplay(messenger, carpenter)
}

func BuildMessageDisplay(shell shell.ShellOut) display.Display {
  return NewMessageDisplay(shell, messengerRec)
}

func BuildBoardDisplay(shell shell.ShellOut, rec boardRec.Rec, board board.Board) display.Display {
  return NewBoardDisplay(shell, rec, board)
}

func buildInGameTerminalSelector(shell shell.ShellIn) selector.Selector {
  return NewInGameTerminalSelector(shell)
}
