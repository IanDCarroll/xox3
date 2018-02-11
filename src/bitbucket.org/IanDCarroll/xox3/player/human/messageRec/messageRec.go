package messageRec

type Rec interface {
  BadSelection() string
  PlayerMove(string) string
}
