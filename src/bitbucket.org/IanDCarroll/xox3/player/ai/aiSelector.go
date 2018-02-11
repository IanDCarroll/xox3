package ai

type aiSelector struct {
  ai Ai
}

func NewAiSelector(ai Ai) aiSelector {
  return aiSelector { ai }
}

func (s aiSelector) ReadFromShell() interface{} {
  return s.ai.GetMove()
}
