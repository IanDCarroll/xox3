package terminal

import "fmt"

func Display(message string) {
  fmt.Println(message)
}

func Welcome() string {
  return `
  Welcome to xox3: an unbeatable game of noughts and crosses!
  Be amazed at your inability to ever win against this mighty juggernaut!

  ( Actual game coming soon... )
  `
}
