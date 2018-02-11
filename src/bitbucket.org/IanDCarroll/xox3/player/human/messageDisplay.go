package human

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/player/human/messageRec"
  "strconv"
)

type messageDisplay struct {
  shell shell.ShellOut
  rec messageRec.Rec
}

func NewMessageDisplay(shell shell.ShellOut, rec messageRec.Rec) messageDisplay {
  return messageDisplay {shell, rec}
}

func (m messageDisplay) WriteToShell(message interface{}) {
  m.shell.Write(m.announceMessage(message))
}

func (m messageDisplay) announceMessage(message interface{}) string {
  var toTerminal string
  switch value := message.(type) {
  case int:
    toTerminal = m.announcePlayerMove(value)
  case string:
    toTerminal = m.announceMessageFromBeyond(value)
  }
  return toTerminal
}

func (m messageDisplay) announcePlayerMove(playerNumber int) string {
  name := strconv.Itoa(playerNumber)
  return m.rec.PlayerMove(name)
}

func (m messageDisplay) announceMessageFromBeyond(message string) string {
  return message
}
