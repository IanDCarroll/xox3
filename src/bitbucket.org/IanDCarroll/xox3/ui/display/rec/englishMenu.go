package rec

import "strconv"

type englishMenu struct {}

func NewEnglishMenu() englishMenu { return englishMenu {} }

var welcome string = `
Welcome to xox3: an unbeatable game of noughts and crosses!
Be amazed at your inability to ever win against this mighty juggernaut!
`

var whichPlayer string = `
Would you like to go first and not win?
-or-
Would you like to go second and not win?

1 - go first
2 - go second
`

var whichMarker string = `
Would you like to be "X" and not win?
-or-
Would you like to be "O" and not win?

`

var markers [2]string = [2]string { "X", "O" }

var badSelection string = `
Woops! Sorry If it wasn't clear.
Please just type the number of the option you would like to choose.
Let's try again :)
`

var startGame string= `
Great! That's all I need to get this party started.

Prepare to not win!
And without further ado,
Let the game begin...
`

func (e englishMenu) Welcome() string { return welcome }

func (e englishMenu) WhichPlayer() string { return whichPlayer }

func (e englishMenu) WhichMarker() string {
  whichMarkerWithOptions := whichMarker
  for i := 0; i < len(markers); i++ {
    whichMarkerWithOptions += strconv.Itoa(i + 1) + " - " + markers[i] + "\n"
  }
  return whichMarkerWithOptions
}

func (e englishMenu) Markers() []string {
  anyMarkers := make([]string, len(markers))
  for i := 0; i < len(markers); i++ {
    anyMarkers[i] = markers[i]
  }
  return anyMarkers
}

func (e englishMenu) BadSelection() string { return badSelection }

func (e englishMenu) StartGame() string { return startGame }
