package ai

import  (
  "bitbucket.org/IanDCarroll/xox3/player/ai/aiRec"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
)

type aiSpyDisplay struct {
  rec aiRec.Rec
  spy display.Display
}

func NewAiSpyDisplay(rec aiRec.Rec, spy display.Display) aiSpyDisplay {
  return aiSpyDisplay {rec, spy}
}

func (d aiSpyDisplay) WriteToShell(message interface{}) {
  d.spy.WriteToShell(d.rec.AiMove())
}
