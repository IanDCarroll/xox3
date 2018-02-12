package gameRec

type Rec interface {
  AiDraw() string
  PlayerWin(string) string
  AiWin() string
}
