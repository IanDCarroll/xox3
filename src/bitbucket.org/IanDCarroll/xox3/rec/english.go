package rec

type english struct {}

func NewEnglish() english { return english {} }

var welcome string = `
Welcome to xox3: an unbeatable game of noughts and crosses!
Be amazed at your inability to ever win against this mighty juggernaut!

( Actual game coming soon... )
`

func (e english) Welcome() string {
  return welcome
}
