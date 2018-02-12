package gameRunner

import  (
  "bitbucket.org/IanDCarroll/xox3/gameRunner/gameRec"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "strconv"
)

type endGameHijackingDisplay struct {
  aiPlayer int
  rec gameRec.Rec
  host display.Display
}

func NewEndGameDisplay(aiPlayer int, rec gameRec.Rec, host display.Display) endGameHijackingDisplay {
  return endGameHijackingDisplay {aiPlayer, rec, host}
}

func (d endGameHijackingDisplay) WriteToShell(message interface{}) {
  d.host.WriteToShell(d.aiEndgame(message.(int)))
}

func (d endGameHijackingDisplay) aiEndgame(winner int) string {
  if winner == 0 { return d.rec.AiDraw() }
  if winner == d.aiPlayer { return d.rec.AiWin() }
  return d.rec.PlayerWin(strconv.Itoa(winner))
}
