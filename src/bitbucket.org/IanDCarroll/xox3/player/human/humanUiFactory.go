package human

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/board"
  "bitbucket.org/IanDCarroll/xox3/player/human/rec"
)

var terminalOut shell.ShellOut = shell.NewTerminal()
var terminalIn shell.ShellIn = shell.NewTerminal()

func BuildHumanUi(rec rec.Rec, board board.Board) ui.Ui {
  display := BuildInGameTerminalDisplay(terminalOut, rec, board)
  selector := buildInGameTerminalSelector(terminalIn, rec)
  return NewUI(display, selector, board, rec)
}

func BuildInGameTerminalDisplay(shell shell.ShellOut, rec rec.Rec, board board.Board) display.Display {
  return NewInGameTerminalDisplay(shell, rec, board)
}

func buildInGameTerminalSelector(shell shell.ShellIn, rec rec.Rec) selector.Selector {
  return NewInGameTerminalSelector(shell)
}
