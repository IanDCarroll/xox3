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
  display := BuildInGameTerminalDisplay(markers, board)
  selector := buildInGameTerminalSelector()
  return NewHumanUI(display, selector, board, messengerRec)
}

func BuildInGameTerminalDisplay(markers []string, board board.Board) display.Display {
  messenger := buildMessageDisplay()
  carpenter := buildBoardDisplay(markers, board)
  return NewInGameTerminalDisplay(messenger, carpenter)
}

func buildMessageDisplay() display.Display {
  return NewMessageDisplay(terminalOut, messengerRec)
}

func buildBoardDisplay(markers []string, board board.Board) display.Display {
  rec := buildBoardRec(markers)
  return NewBoardDisplay(terminalOut, rec, board)
}

func buildBoardRec(markers []string) boardRec.Rec {
  return boardRec.NewTerminalRec(markers)
}

func buildInGameTerminalSelector() selector.Selector {
  return NewInGameTerminalSelector(terminalIn)
}
