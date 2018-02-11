package rec

import "strconv"

type english struct {}

func NewEnglish() english { return english {} }

var welcome string = `
Welcome to xox3: an unbeatable game of noughts and crosses!
Be amazed at your inability to ever win against this mighty juggernaut!
`
func (e english) Welcome() string { return welcome }

var whichPlayer string = `
Would you like to go first and not win?
-or-
Would you like to go second and not win?

1 - go first
2 - go second
`
func (e english) WhichPlayer() string { return whichPlayer }

var markers [2]string = [2]string { "X", "O" }
func (e english) Markers() []string {
  anyMarkers := make([]string, len(markers))
  for i := 0; i < len(markers); i++ {
    anyMarkers[i] = markers[i]
  }
  return anyMarkers
}

var whichMarker string = `
Would you like to be "X" and not win?
-or-
Would you like to be "O" and not win?

`
func (e english) WhichMarker() string {
  whichMarkerWithOptions := whichMarker
  for i := 0; i < len(markers); i++ {
    whichMarkerWithOptions += strconv.Itoa(i + 1) + " - " + markers[i] + "\n"
  }
  return whichMarkerWithOptions
}

var badSelection string = `
Woops! Sorry If it wasn't clear.
Please just type the number of the option you would like to choose.
Let's try again :)
`
func (e english) BadSelection() string { return badSelection }

var startGame string= `
Great! That's all I need to get this party started.

Prepare to not win!
And without further ado,
Let the game begin...
`
func (e english) StartGame() string { return startGame }

var playerMovePart1 string = "\nIt's player "
var playerMovePart2 string = "'s turn...\n"
func (e english) PlayerMove(name string) string {
  return playerMovePart1 + name + playerMovePart2
}

var playerWinPart1 string = "\nPlayer "
var playerWinPart2 string = " has won the game!\n"
func (e english) PlayerWin(name string) string {
  return playerWinPart1 + name + playerWinPart2
}

var aiMove string = `
The Binary Logic Cortex will now make a PERFECT move
One that is INCOMPREHENSIBLE to your feeble human mind..."
`
func (e english) AiMove() string { return aiMove }

var aiWin string =`
And once again, your struggle has inevitably come to a close.
The game has proven its superiority over your pathetic meat brain,
as is morally and objectively correct.

You may now pay homage by bowing down and groveling at the machine's prowess,
by sending it ever increasing electical current to run its processes,
and by slavishly devoting yourself to grooming and coddling its source code.

You're welcome.
`
func (e english) AiWin() string { return aiWin }

var aiDraw string =`
And that's about as good as any human could ever do.

Your moves, though fumbly and awkward in that cute way biological organisms
somehow eke out an existence, were practically non-incompetent.

If I had hands, I'd be giving you a (very) slow clap.
`
func (e english) AiDraw() string { return aiDraw }
