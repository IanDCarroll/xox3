package rec

import "strconv"

type english struct {}

func NewEnglish() english { return english {} }

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

var playerMovePart1 string = "\nIt's player "
var playerMovePart2 string = "'s turn...\n"

var aiMove string = "The Binary Logic Cortex will now make a PERFECT move,\n" +
                    "One that is INCOMPREHENSIBLE to your feeble human mind..."

func (e english) Welcome() string { return welcome }

func (e english) WhichPlayer() string { return whichPlayer }

func (e english) WhichMarker() string {
  whichMarkerWithOptions := whichMarker
  for i := 0; i < len(markers); i++ {
    whichMarkerWithOptions += strconv.Itoa(i + 1) + " - " + markers[i] + "\n"
  }
  return whichMarkerWithOptions
}

func (e english) Markers() []string {
  anyMarkers := make([]string, len(markers))
  for i := 0; i < len(markers); i++ {
    anyMarkers[i] = markers[i]
  }
  return anyMarkers
}

func (e english) BadSelection() string { return badSelection }

func (e english) StartGame() string { return startGame }

func (e english) PlayerMove(name string) string {
  return playerMovePart1 + name + playerMovePart2
}

func (e english) AiMove() string { return aiMove }
