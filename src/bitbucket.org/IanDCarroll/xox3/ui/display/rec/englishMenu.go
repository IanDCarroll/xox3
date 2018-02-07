package rec

type englishMenu struct {}

func NewEnglishMenu() englishMenu { return englishMenu {} }

var welcome string = `
Welcome to xox3: an unbeatable game of noughts and crosses!
Be amazed at your inability to ever win against this mighty juggernaut!

( Actual game coming soon... )
`

func (e englishMenu) Welcome() string {
  return welcome
}
