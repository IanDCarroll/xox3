package rec

type Rec interface {
  Welcome() string
  WhichPlayer() string
  WhichMarker() string
  Markers() []string
  BadSelection() string
  StartGame() string
}
