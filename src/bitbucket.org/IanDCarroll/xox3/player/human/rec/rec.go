package rec

type Rec interface {
  BadSelection() string
  PlayerMove(string) string
  RowSegment(string, int, int, int) string
  DeckSegment(int, int, int) string
}
