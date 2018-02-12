package ai

import  (
  "bitbucket.org/IanDCarroll/xox3/player/ai/aiRec"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
)

type aiSpyDisplay struct {
  rec aiRec.Rec
  spy display.Display
  ai Ai
}

func NewAiSpyDisplay(rec aiRec.Rec, spy display.Display, ai Ai) aiSpyDisplay {
  return aiSpyDisplay {rec, spy, ai}
}

func (d aiSpyDisplay) WriteToShell(id interface{}) {
  d.tellAiWhoItIs(id)
  d.spy.WriteToShell(d.rec.AiMove())
}

func (d aiSpyDisplay) tellAiWhoItIs(identity interface{}) {
  d.ai.SetID(identity.(int))
}
