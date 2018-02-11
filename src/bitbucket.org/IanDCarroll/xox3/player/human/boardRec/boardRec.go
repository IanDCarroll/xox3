package boardRec

type Rec interface {
  RowSegment(string, int, int, int) string
  DeckSegment(int, int, int) string
}
